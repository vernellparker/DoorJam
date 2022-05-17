package main

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type Resources struct {
	Data interface{}
}

func (r Resources) AsImage() *ebiten.Image {
	return r.Data.(*ebiten.Image)
}

func (game *Game) LoadResource(resourcePath string) *Resources {
	res, exists := game.Resources[resourcePath]
	if !exists {
		res = &Resources{}
		fileData, err := assets.Open(resourcePath)
		if err != nil {
			panic(err)
		}

		img, _, err := image.Decode(fileData)
		if err != nil {
			panic(err)
		}
		res.Data = ebiten.NewImageFromImage(img)
		game.Resources[resourcePath] = res
	}
	return res
}
