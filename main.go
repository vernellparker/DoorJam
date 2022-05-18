package main

import (
	"embed"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/solarlune/ldtkgo"
)

//go:embed assets/*
var assets embed.FS

type Game struct {
	Level                     *Level
	LdtkProject               *ldtkgo.Project
	ScreenWidth, ScreenHeight int
	Resources                 map[string]*Resources
	Debug                     bool
}

func NewGame() *Game {
	var err error

	ebiten.SetWindowTitle("Door Jam")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	game := &Game{
		ScreenWidth:  384,
		ScreenHeight: 216,
		Resources:    map[string]*Resources{},
	}

	game.LdtkProject, err = ldtkgo.Open("assets/levels.ldtk")
	if err != nil {
		panic(err)
	}

	game.Level = NewLevel(game)

	return game
}

func (game *Game) Update() error {

	//Reload the level
	if ebiten.IsKeyPressed(ebiten.KeyR) {
		game.Level = NewLevel(game)
	}
	//Puts the game in Debug mode
	if inpututil.IsKeyJustPressed(ebiten.KeyF1) {
		game.Debug = !game.Debug
	}

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

func main() {
	game := NewGame()
	err := ebiten.RunGame(game)
	if err != nil {
		return
	}
}
