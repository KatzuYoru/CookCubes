package game

import (
	"fmt"
	"os"

	"github.com/KotzuYaru/cubes/components"
	"github.com/KotzuYaru/cubes/constants"
	"github.com/KotzuYaru/cubes/entities"
	"github.com/KotzuYaru/cubes/utils"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

type Game struct {
	Tilemap         *entities.TileMap
	Tileset         *ebiten.Image
	Player          *entities.Player
	TitleFont       font.Face
	MenuButtons     []*components.Button
	settingsButtons []*components.Button
	SettingsButton  *components.Button
	NavigationStack utils.Stack
}

func NewGame() (*Game, error) {
	game := &Game{}

	titleFont, err := loadFont(24, 72)
	if err != nil {
		return nil, fmt.Errorf("failed to load font: %w", err)
	}
	game.TitleFont = titleFont

	tileMap, err := entities.LoadTileMap("assets/maps/spawn.json")
	if err != nil {
		return nil, fmt.Errorf("unable to load the tilemap: %w", err)
	}

	tileset, err := loadImage("assets/images/TilesetFloor.png")
	if err != nil {
		return nil, fmt.Errorf("failed to load tileset image: %w", err)
	}

	player, err := entities.GetNewEntity[entities.Player]("player", 100, 100, "assets/images/ninja.png")
	if err != nil {
		return nil, fmt.Errorf("unable to load the player: %w", err)
	}

	game.Tilemap = tileMap
	game.Tileset = tileset
	game.Player = player
	game.NavigationStack.Push("Menu")
	screenWidth, screenHeight := ebiten.WindowSize()
	buttonX := (screenWidth - constants.Buttons.ButtonWidth) / 2
	game.SettingsButton = createButton(100, 100, "Play", titleFont, func() {
		fmt.Println("Setting icon clicked")
	})

	game.MenuButtons = createMenuButtons(buttonX, screenHeight, titleFont, game)
	game.settingsButtons = createSettingsButtons(buttonX, screenHeight, titleFont, game)

	return game, nil
}

func loadFont(size float64, dpi float64) (font.Face, error) {
	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		return nil, err
	}
	return opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    size,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
}

func loadImage(path string) (*ebiten.Image, error) {
	img, _, err := ebitenutil.NewImageFromFile(path)
	return img, err
}

func createButton(x, y int, label string, font font.Face, action func()) *components.Button {
	return components.NewButton(
		x, y,
		constants.Buttons.ButtonWidth, constants.Buttons.ButtonHeight,
		label, font, action,
	)
}

func createMenuButtons(x, y int, font font.Face, game *Game) []*components.Button {
	return []*components.Button{
		createButton(x, y/2-60, "Play", font, func() { game.NavigationStack.Push("Play") }),
		createButton(x, y/2, "Settings", font, func() {
			game.NavigationStack.Push("Settings")
		}),
		createButton(x, y/2+60, "Exit", font, func() {
			os.Exit(0)
		}),
	}
}

func createSettingsButtons(x, y int, font font.Face, game *Game) []*components.Button {
	return []*components.Button{
		createButton(x, y/2-90, "Volume", font, func() { fmt.Println("Volume button clicked") }),
		createButton(x, y/2-30, "Hardness", font, func() { fmt.Println("Hardness button clicked") }),
		createButton(x, y/2+30, "Auto Save", font, func() { fmt.Println("Auto Save button clicked") }),
		createButton(x, y/2+90, "Back", font, func() {
			game.NavigationStack.Pop()
		}),
	}
}
