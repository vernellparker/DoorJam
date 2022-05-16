package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/solarlune/goaseprite"
)

type Player struct{
	Level *Level
	Aseprite goaseprite.File
	Sprite *ebiten.Image

}

func NewPlayer(level *Level) *Player {
	player := &Player{
		Level: level,
	}
	asset := OpenAssets("assets/gfx/player.json")
	player.Aseprite = *goaseprite.Read(asset)
	player.Aseprite.Play("idle")
	player.Sprite = level.Game.LoadResource("assets/gfx/player.png").AsImage()
	return player
}

func (p *Player) Update() {
	p.Aseprite.Update(float32(p.Level.Game.Delta()))
}

func (p *Player) Draw(img *ebiten.Image) {
	
}