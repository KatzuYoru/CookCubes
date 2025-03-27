package game

import (
	"image"

	"github.com/KotzuYaru/cubes/constants"
	"github.com/hajimehoshi/ebiten/v2"
)

func (g *Game) Draw(screen *ebiten.Image) {
	opts := ebiten.DrawImageOptions{}

	for _, layer := range g.Tilemap.Layers {
		for index, id := range layer.Data {

			x := index % layer.Width
			y := index / layer.Width

			x *= g.Tilemap.TileWidth
			y *= g.Tilemap.TileWidth

			srcX := (id - 1) % constants.TILE_COUNT
			srcY := (id - 1) / constants.TILE_COUNT

			srcX *= g.Tilemap.TileWidth
			srcY *= g.Tilemap.TileWidth
			opts.GeoM.Translate(float64(x), float64(y))

			screen.DrawImage(
				g.Tileset.SubImage(image.Rect(srcX, srcY, srcX+g.Tilemap.TileWidth, srcY+g.Tilemap.TileWidth)).(*ebiten.Image),
				&opts,
			)

			opts.GeoM.Reset()
		}
	}
	opts.GeoM.Translate(g.Player.X, g.Player.Y)
	screen.DrawImage(g.Player.Img, &opts)
}
