package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	renderer "github.com/solarlune/ldtkgo/ebitenrenderer"
)

type Level struct {
	Game        *Game
	Renderer    *renderer.EbitenRenderer
	Player *Player
	GameObjects []GameObject
}

func NewLevel(game *Game) *Level {
	level := &Level{
		Game:        game,
		GameObjects: []GameObject{},
	}

	level.Renderer = renderer.NewEbitenRenderer(renderer.NewDiskLoader("assets"))

	level.Renderer.Render(level.Game.LdtkProject.Levels[0])

	level.Player = NewPlayer(level)
	return level
}

func (l *Level) Update() {
	for _, g := range l.GameObjects {
		g.Update()
	}
}

func (l *Level) Draw(img *ebiten.Image) {
	for _, layer := range l.Renderer.RenderedLayers {
		img.DrawImage(layer.Image, &ebiten.DrawImageOptions{})
	}

	for _, g := range l.GameObjects {
		g.Draw(img)
	}
}

func (l *Level) Add(gameObject GameObject) {
	l.GameObjects = append(l.GameObjects, gameObject)
}

func (l *Level) Remove(gameObject GameObject) {
	for i, g := range l.GameObjects {
		if g == gameObject {
			l.GameObjects = append(l.GameObjects[:i], l.GameObjects[i+1:]...)
			break
		}
	}
}
