package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	renderer "github.com/solarlune/ldtkgo/ebitenrenderer"
)

type Level struct {
	Game     *Game
	Renderer *renderer.EbitenRenderer
}

func NewLevel(game *Game) *Level {
	level := &Level{
		Game: game,
	}

	level.Renderer = renderer.NewEbitenRenderer(renderer.NewDiskLoader("assets"))

	level.Renderer.Render(level.Game.LdtkProject.Levels[0])
	return level
}

func (l *Level) Update() {

}

func (l *Level) Draw(img *ebiten.Image) {
	for _, layer := range l.Renderer.RenderedLayers {
		img.DrawImage(layer.Image, &ebiten.DrawImageOptions{})
	}
}
