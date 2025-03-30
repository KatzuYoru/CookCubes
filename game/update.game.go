package game

func (g *Game) Update() error {
	screen, _ := g.NavigationStack.Peek()

	if screen == "Menu" {
		for _, button := range g.MenuButtons {
			button.Update()
		}
	} else if screen == "Settings" {
		for _, button := range g.settingsButtons {
			button.Update()
		}
	} else if screen == "game" {
		// Update game logic
		// Player movement, etc.
	}

	return nil
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 800, 600
}
