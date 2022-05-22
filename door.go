package main

import (
	"image"
	"io/ioutil"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/solarlune/goaseprite"
	"github.com/solarlune/resolv"
)

type Door struct {
	Level    *Level
	Aseprite goaseprite.File
	Sprite   *ebiten.Image
	Object   *resolv.Object
}

func NewDoor(level *Level) *Door {
	door := &Door{
		Level: level,
	}

	//Loads player json data
	s, _ := ioutil.ReadFile("assets/gfx/Door.json")
	door.Aseprite = *goaseprite.Read(s)

	door.Aseprite.Play("closed")

	//loads the player Sprite from embedded
	door.Sprite = level.Game.LoadResource("assets/gfx/Door.png").AsImage()

	//Resolv
	door.Object = resolv.NewObject(0, 0, 24, 24)
	door.Level.Space.Add(door.Object)

	return door
}

func (d *Door) Update() {
	d.Aseprite.Update(float32(1.0 / 60))
}

func (p *Door) Draw(img *ebiten.Image) {

	sub := p.Sprite.SubImage(image.Rect(p.Aseprite.CurrentFrameCoords()))
	options := &ebiten.DrawImageOptions{}

	//Resolv Object
	options.GeoM.Translate(p.Object.X, p.Object.Y)

	img.DrawImage(sub.(*ebiten.Image), options)
}

func (d *Door) Depth() int {

	return 100
}
