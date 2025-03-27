package entities

import "github.com/hajimehoshi/ebiten/v2"

type Sprite struct {
	Img      *ebiten.Image
	X        float64
	Y        float64
	Strength float64
	InitX    float64
	InitY    float64
	Width    float64
	Height   float64
}

func GetSprite(tile *ebiten.Image, initX float64, initY float64, width float64, height float64) *Sprite {
	return &Sprite{
		Img:    tile,
		X:      initX,
		Y:      initY,
		InitX:  initX,
		InitY:  initY,
		Width:  width,
		Height: height,
	}
}
