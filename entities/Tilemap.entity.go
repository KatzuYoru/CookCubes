package entities

import (
	"encoding/json"
	"os"
)

type Layers struct {
	Data   []int `json:"data"`
	Height int   `json:"height"`
	Width  int   `json:"width"`
}

type TileMap struct {
	Height     int      `json:"height"`
	Width      int      `json:"width"`
	TileWidth  int      `json:"tilewidth"`
	TileHeight int      `json:"tileheight"`
	Layers     []Layers `json:"layers"`
}

func LoadTileMap(fileName string) (*TileMap, error) {
	file, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	var tileMap TileMap

	if err = json.Unmarshal(file, &tileMap); err != nil {
		return nil, err
	}

	return &tileMap, nil
}
