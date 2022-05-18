package main

import (
	"image"
	"io/ioutil"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/solarlune/goaseprite"
	"github.com/solarlune/resolv"
)

type Player struct {
	Level         *Level
	Aseprite      goaseprite.File
	Sprite        *ebiten.Image
	Object        *resolv.Object
	SpriteOffSetX float64
	SpriteOffSetY float64
	OnGround      bool
	DeltaX        float64
	Flipped       bool
}

func NewPlayer(level *Level) *Player {
	player := &Player{
		Level: level,
	}

	//Loads player json data
	s, _ := ioutil.ReadFile("assets/gfx/player.json")
	player.Aseprite = *goaseprite.Read(s)

	player.Aseprite.Play("idle")

	//Gets the slice that is used to better position player and collision rect
	slice := player.Aseprite.Slices[0]
	player.SpriteOffSetX = float64(-slice.Keys[0].X)
	player.SpriteOffSetY = float64(-slice.Keys[0].Y)

	//loads the player Sprite from embedded
	player.Sprite = level.Game.LoadResource("assets/gfx/player.png").AsImage()

	//Resolv
	player.Object = resolv.NewObject(0, 0, 16, 16)
	player.Level.Space.Add(player.Object)

	return player
}

func (p *Player) Update() {
	p.OnGround = false
	p.Aseprite.Update(float32(1.0 / 60))

	p.DeltaX = 0.0
	dy := 2.0
	moveSpeed := 2.0

	//Checks if the player contacts the walls or floor
	if check := p.Object.Check(0, dy, "solid"); check != nil {
		if dy >= 0 {
			p.OnGround = true
		}
		dy = check.ContactWithCell(check.Cells[0]).Y()
	}

	if p.OnGround {
		//Checks for input
		if ebiten.IsKeyPressed(ebiten.KeyA) {
			p.DeltaX = -moveSpeed
		}
		if ebiten.IsKeyPressed(ebiten.KeyD) {
			p.DeltaX = moveSpeed
		}
	}

	//Resolv
	p.Object.Y += dy
	p.Object.X += p.DeltaX

}

func (p *Player) Draw(img *ebiten.Image) {
	//Animation
	anim := "idle"
	if !p.OnGround {
		anim = "fall"
	} else if p.DeltaX != 0 {
		anim = "run"
	} else {
		anim = "idle"
	}

	if p.DeltaX < 0 {
		p.Flipped = true
	} else if p.DeltaX > 0 {
		p.Flipped = false
	}

	p.Aseprite.Play(anim)

	sub := p.Sprite.SubImage(image.Rect(p.Aseprite.CurrentFrameCoords()))
	options := &ebiten.DrawImageOptions{}

	//Flips image
	if p.Flipped {
		options.GeoM.Translate(-float64(sub.Bounds().Dx()/2), -float64(sub.Bounds().Dy()/2))
		options.GeoM.Scale(-1, 1)
		options.GeoM.Translate(float64(sub.Bounds().Dx()/2), float64(sub.Bounds().Dy()/2))
	}

	//Resolv Object
	options.GeoM.Translate(p.Object.X+p.SpriteOffSetX, p.Object.Y+p.SpriteOffSetY)

	img.DrawImage(sub.(*ebiten.Image), options)
}
