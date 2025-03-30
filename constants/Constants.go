package constants

import "image/color"

const (
	TILESIZE   = 16
	TILE_COUNT = 22
)

var Colors = struct {
	Rust color.RGBA
}{
	Rust: color.RGBA{183, 65, 14, 255},
}

var Buttons = struct {
	ButtonWidth  int
	ButtonHeight int
}{
	ButtonWidth:  200,
	ButtonHeight: 50,
}
