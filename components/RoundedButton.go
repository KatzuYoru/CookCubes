package components

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func DrawRoundedRect(screen *ebiten.Image, x, y, width, height, radius float32, clr color.Color) {
	// Ensure radius is not too large
	if radius > width/2 {
		radius = width / 2
	}
	if radius > height/2 {
		radius = height / 2
	}

	// Create a new image for the button
	buttonImg := ebiten.NewImage(int(width), int(height))

	// Draw the main rectangle (center part)
	vector.DrawFilledRect(buttonImg, radius, 0, width-2*radius, height, clr, true)

	// Draw the left and right side rectangles
	vector.DrawFilledRect(buttonImg, 0, radius, radius, height-2*radius, clr, true)
	vector.DrawFilledRect(buttonImg, width-radius, radius, radius, height-2*radius, clr, true)

	// Draw the four corners as circles
	r, g, b, a := clr.RGBA()
	rgba := color.RGBA{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), uint8(a >> 8)}

	// Top left corner
	vector.DrawFilledCircle(buttonImg, radius, radius, radius, rgba, true)

	// Top right corner
	vector.DrawFilledCircle(buttonImg, width-radius, radius, radius, rgba, true)

	// Bottom left corner
	vector.DrawFilledCircle(buttonImg, radius, height-radius, radius, rgba, true)

	// Bottom right corner
	vector.DrawFilledCircle(buttonImg, width-radius, height-radius, radius, rgba, true)

	// Draw the button image to the screen
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage(buttonImg, op)
}
