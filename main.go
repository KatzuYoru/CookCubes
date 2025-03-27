package main

import (
	"log"

	"github.com/KotzuYaru/cubes/game"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	game, err := game.NewGame()

	if err != nil {
		log.Fatal("Server error, unable to start the game.")
	}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
