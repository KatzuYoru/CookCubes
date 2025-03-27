package entities

import (
	"fmt"
	"image"

	"github.com/KotzuYaru/cubes/constants"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func GetNewEntity[T any](entityType string, initX, initY float64, spritePath string) (*T, error) {
	img, _, err := ebitenutil.NewImageFromFile(spritePath)
	if err != nil {
		return nil, fmt.Errorf("failed to load sprite: %w", err)
	}
	tile := img.SubImage(image.Rect(0, 0, constants.TILESIZE, constants.TILESIZE)).(*ebiten.Image)

	// TODO: to be replaced
	width, height := tile.Size()

	var entity any

	switch entityType {
	case "player":
		entity = &Player{
			Sprite: GetSprite(tile, initX, initY, float64(width), float64(height)),
		}
	case "enemy":
		entity = &Enemy{
			Sprite: GetSprite(tile, initX, initY, float64(width), float64(height)),
		}
	default:
		return nil, fmt.Errorf("unknown entity type: %s", entityType)
	}
	typedEntity, ok := entity.(*T)
	if !ok {
		return nil, fmt.Errorf("failed to cast entity to expected type")
	}

	return typedEntity, nil
}

func (a *Sprite) CollidesWith(b *Sprite) bool {
	return a.X < b.X+b.Width &&
		a.X+a.Width > b.X &&
		a.Y < b.Y+b.Height &&
		a.Y+a.Height > b.Y
}
