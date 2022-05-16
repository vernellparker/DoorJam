package main

import (
	"embed"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/solarlune/ldtkgo"
)

//go:embed assets/*
var assets embed.FS

type Game struct {
	Level                     *Level
	LdtkProject               *ldtkgo.Project
	ScreenWidth, ScreenHeight int
}

func NewGame() *Game {
	var err error

	ebiten.SetWindowTitle("Door Jam")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	game := &Game{
		ScreenWidth:  384,
		ScreenHeight: 216,
	}

	game.LdtkProject, err = ldtkgo.Open("assets/levels.ldtk")
	if err != nil {
		panic(err)
	}

	game.Level = NewLevel(game)

	return game
}

func (game *Game) Update() error {
	game.Level.Update()
	return nil
}

func (game *Game) Draw(image *ebiten.Image) {
	image.Fill(color.RGBA{40, 40, 30, 255})
	game.Level.Draw(image)
}

func (game *Game) Layout(w, h int) (int, int) {
	return game.ScreenWidth, game.ScreenHeight
}
func (game *Game)Delta() float64 {
	return 1.0/ebiten.CurrentTPS()
}

func main() {
	game := NewGame()
	err := ebiten.RunGame(game)
	if err != nil {
		return
	}
}
