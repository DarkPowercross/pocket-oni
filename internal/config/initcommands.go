package config

import "github.com/Darkpowercross/pocket-oni/internal/config/appenums"

type Command struct {
	ID          string
	Menu        string
	Label       string
	Description string
	Shortcut    rune
	Run         func()
}

func (c *Config) InitCommands() {
	c.Commands = []Command{
		{
			ID:          "feed",
			Menu:        "main",
			Label:       "feed",
			Description: "feeds the creature",
			Shortcut:    'f',
			Run: func() {
				c.SpriteInformation.AddFoodMessage(1)
			},
		},
		{
			ID:          "water",
			Menu:        "main",
			Label:       "water",
			Description: "waters the creature",
			Shortcut:    'w',
			Run: func() {
				c.SpriteInformation.AddWaterMessage(1)
			},
		},
		// {
		// 	ID:          "pet",
		// 	Menu:        "main",
		// 	Label:       "pet",
		// 	Description: "play with your character",
		// 	Shortcut:    'p',
		// 	Run: func() {
		// 		c.SpriteInformation.SetHappiness(1)
		// 	},
		// },
		{
			ID:          "move",
			Menu:        "main",
			Label:       "move",
			Description: "move your character",
			Shortcut:    'm',
			Run: func() {
				c.SetMenu("move")
			},
		},

		{
			ID:          "ocean",
			Menu:        "move",
			Label:       "ocean",
			Description: "go to the ocean",
			Shortcut:    'o',
			Run: func() {
				c.SpriteInformation.Location = appenums.Ocean
				c.SetMenu("main")
			},
		},
		{
			ID:          "forest",
			Menu:        "move",
			Label:       "forest",
			Description: "go to the forest",
			Shortcut:    'f',
			Run: func() {
				c.SpriteInformation.Location = appenums.Forest
				c.SetMenu("main")
			},
		},
		{
			ID:          "barn",
			Menu:        "move",
			Label:       "barn",
			Description: "go to the barn",
			Shortcut:    'b',
			Run: func() {
				c.SpriteInformation.Location = appenums.Barn
				c.SetMenu("main")
			},
		},
		{
			ID:          "return",
			Menu:        "move",
			Label:       "return",
			Description: "return to main menu",
			Shortcut:    'r',
			Run: func() {
				c.SetMenu("main")
			},
		},
	}
}

func (c *Config) SetMenu(menu string) {
	menuView := c.App.Approot.View.Content.Sub.Menu
	menuView.Clear()

	for _, cmd := range c.Commands {
		if cmd.Menu != menu {
			continue
		}

		cmdCopy := cmd
		menuView.AddItem(
			cmdCopy.Label,
			cmdCopy.Description,
			cmdCopy.Shortcut,
			func() {
				cmdCopy.Run()
			},
		)
	}

	if menu == "main" && c.SpriteInformation.Location == appenums.Barn {
		menuView.AddItem(
			"clean",
			"clean the barn area",
			'c',
			func() {
				c.SpriteInformation.CleanWaste()
			},
		)
	}
}
