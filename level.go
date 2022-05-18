package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/solarlune/ldtkgo"
	renderer "github.com/solarlune/ldtkgo/ebitenrenderer"
	"github.com/solarlune/resolv"
)

type Level struct {
	Game        *Game
	Renderer    *renderer.EbitenRenderer
	Player      *Player
	GameObjects []GameObject
	Space       *resolv.Space
}

func NewLevel(game *Game) *Level {

	level := &Level{
		Game:        game,
		GameObjects: []GameObject{},
	}

	level.Renderer = renderer.NewEbitenRenderer(renderer.NewDiskLoader("assets"))

	level.Load(game.LdtkProject.Levels[0])

	return level
}

func (l *Level) Load(ldtkLevel *ldtkgo.Level) {
	//Resolv
	l.Space = resolv.NewSpace(ldtkLevel.Width, ldtkLevel.Height, 16, 16)

	//adds to GameObjects
	l.Add(NewPlayer(l))

	l.Renderer.Render(ldtkLevel)

	//This is the code that loads the initgrid for collision detection
	for _, intGridValue := range ldtkLevel.LayerByIdentifier("IntGrid").IntGrid {
		pos := intGridValue.Position
		obj := resolv.NewObject(float64(pos[0]), float64(pos[1]), 16, 16)
		obj.AddTags("solid")
		l.Space.Add(obj)
	}
}

func (l *Level) Update() {
	for _, g := range l.GameObjects {
		g.Update()
	}
}

func (l *Level) Draw(img *ebiten.Image) {
	//renders layers from ldtk
	for _, layer := range l.Renderer.RenderedLayers {
		img.DrawImage(layer.Image, &ebiten.DrawImageOptions{})
	}

	//Renders all game objects to the screen
	for _, g := range l.GameObjects {
		g.Draw(img)
	}

	//Game Debug View
	if l.Game.Debug {
		for _, obj := range l.Space.Objects() {
			c := color.RGBA{255, 255, 255, 64}
			ebitenutil.DrawRect(img, obj.X, obj.Y, obj.W, obj.H, c)
		}
	}

}

//Add things to level
func (l *Level) Add(gameObject GameObject) {
	l.GameObjects = append(l.GameObjects, gameObject)
}

//Remove things to level
func (l *Level) Remove(gameObject GameObject) {
	for i, g := range l.GameObjects {
		if g == gameObject {
			l.GameObjects = append(l.GameObjects[:i], l.GameObjects[i+1:]...)
			break
		}
	}
}
