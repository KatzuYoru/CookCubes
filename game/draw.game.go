package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

func (g *Game) Draw(screen *ebiten.Image) {
	currentScreen, _ := g.NavigationStack.Peek()

	if currentScreen == "game" {
		// draw somethings
	} else {
		width, height := screen.Bounds().Dx(), screen.Bounds().Dy()
		overlay := ebiten.NewImage(width, height)
		overlay.Fill(color.RGBA{20, 20, 30, 255})
		screen.DrawImage(overlay, &ebiten.DrawImageOptions{})

		if currentScreen == "Menu" {
			for _, button := range g.MenuButtons {
				button.Draw(screen)
			}
		} else if currentScreen == "Settings" {
			// text.Draw(screen, "Settings", g.TitleFont, width/5, height/4, constants.Colors.Rust)
			for _, button := range g.settingsButtons {
				button.Draw(screen)
			}
		}
	}
}
