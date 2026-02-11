package config

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
				c.SpriteInformation.AddFood(1)
			},
		},
		{
			ID:          "water",
			Menu:        "main",
			Label:       "water",
			Description: "waters the creature",
			Shortcut:    'w', // changed from duplicate 'f'
			Run: func() {
				c.SpriteInformation.AddWater(1)
			},
		},
		{
			ID:          "pet",
			Menu:        "main",
			Label:       "pet",
			Description: "play with your character",
			Shortcut:    'p',
			Run: func() {
				c.SpriteInformation.Happiness++
			},
		},
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

		// Move submenu

		{
			ID:          "ocean",
			Menu:        "move",
			Label:       "ocean",
			Description: "go to the ocean",
			Shortcut:    'o',
			Run: func() {
				c.SpriteInformation.Location = "ocean"
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
				c.SpriteInformation.Location = "forest"
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
				c.SpriteInformation.Location = "barn"
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

		c := cmd
		menuView.AddItem(
			c.Label,
			c.Description,
			c.Shortcut,
			func() {
				c.Run()
			},
		)
	}
}
