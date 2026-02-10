package app

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type SpriteInfo struct {
	Health   int
	Location string
}

type View struct {
	Menu *tview.List
}

type Command struct {
	ID          string
	Menu        string
	Label       string
	Description string
	Shortcut    rune
	Run         func(*App)
}

func (a *App) SetCommands() {
	a.Commands = []Command{
		{
			ID:          "feed",
			Menu:        "main",
			Label:       "feed",
			Description: "feeds the creature",
			Shortcut:    'f',
			Run: func(a *App) {
				if a.Spriteinfo.Food < Maxfood {
					a.Spriteinfo.Food += 1
				} else {
					a.SetNewMessage("Oni is full", true, tcell.ColorIndianRed)
				}
			},
		},
		{
			ID:          "water",
			Menu:        "main",
			Label:       "water",
			Description: "waters the creature",
			Shortcut:    'f',
			Run: func(a *App) {
				if a.Spriteinfo.Water < Maxwater {
					a.Spriteinfo.Water += 1
				} else {
					a.SetNewMessage("Oni cant be watered", true, tcell.ColorDarkRed)
				}
			},
		},
		{
			ID:          "pet",
			Menu:        "main",
			Label:       "pet",
			Description: "play with your character",
			Shortcut:    'p',
			Run: func(a *App) {
				a.Spriteinfo.Happiness += 1
			},
		},
		{
			ID:          "move",
			Menu:        "main",
			Label:       "move",
			Description: "move your character",
			Shortcut:    'm',
			Run: func(a *App) {
				a.SetMenu("move")
			},
		},
		{
			ID:          "ocean",
			Menu:        "move",
			Label:       "ocean",
			Description: "go to the ocean",
			Shortcut:    'o',
			Run: func(a *App) {
				a.Spriteinfo.Location = "ocean"
				a.SetMenu("main")
			},
		},
		{
			ID:          "forest",
			Menu:        "move",
			Label:       "forest",
			Description: "go to the ocean",
			Shortcut:    'f',
			Run: func(a *App) {
				a.Spriteinfo.Location = "forest"
				a.SetMenu("main")
			},
		},

		{
			ID:          "barn",
			Menu:        "move",
			Label:       "barn",
			Description: "go to the barn",
			Shortcut:    'f',
			Run: func(a *App) {
				a.Spriteinfo.Location = "barn"
				a.SetMenu("main")
			},
		},
		{
			ID:          "return",
			Menu:        "move",
			Label:       "return",
			Description: "go to the barn",
			Shortcut:    'r',
			Run: func(a *App) {
				a.SetMenu("main")
			},
		},
	}
}

func (a *App) SetMenu(menu string) {
	a.View.Menu.Clear()

	for _, cmd := range a.Commands {
		if cmd.Menu != menu {
			continue
		}

		c := cmd
		a.View.Menu.AddItem(
			c.Label,
			c.Description,
			c.Shortcut,
			func() { c.Run(a) },
		)
	}
}
