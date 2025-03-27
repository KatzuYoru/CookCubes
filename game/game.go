package game

import (
	"fmt"
	"log"

	"github.com/KotzuYaru/cubes/entities"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	Tilemap *entities.TileMap
	Tileset *ebiten.Image
	Player  *entities.Player
}

func NewGame() (*Game, error) {
	tileMap, err := entities.LoadTileMap("assets/maps/spawn.json")
	if err != nil {
		fmt.Println(err)
		log.Fatal("Unable to load the tilemap")
	}
	tileset, _, err := ebitenutil.NewImageFromFile("assets/images/TilesetFloor.png")
	if err != nil {
		return nil, fmt.Errorf("failed to load tileset image: %w", err)
	}

	player, err := entities.GetNewEntity[entities.Player]("player", 100, 100, "assets/images/ninja.png")
	if err != nil {
		log.Fatal("Unable to load the player")
	}

	return &Game{
		Tilemap: tileMap,
		Tileset: tileset,
		Player:  player,
	}, nil
}
