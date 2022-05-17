package main

import (
	"image"
	"io/ioutil"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/solarlune/goaseprite"
)

type Player struct {
	Level    *Level
	Aseprite goaseprite.File
	Sprite   *ebiten.Image
}

func NewPlayer(level *Level) *Player {
	player := &Player{
		Level: level,
	}

	//Loads player json data
	s, _ := ioutil.ReadFile("assets/gfx/player.json")
	player.Aseprite = *goaseprite.Read(s)

	player.Aseprite.Play("idle")

	player.Sprite = level.Game.LoadResource("assets/gfx/player.png").AsImage()
	return player
}

func (p *Player) Update() {
	p.Aseprite.Update(float32(1.0 / 60))
}

func (p *Player) Draw(img *ebiten.Image) {

	sub := p.Sprite.SubImage(image.Rect(p.Aseprite.CurrentFrameCoords()))
	img.DrawImage(sub.(*ebiten.Image), &ebiten.DrawImageOptions{})
}
