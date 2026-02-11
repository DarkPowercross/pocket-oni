package config

import (
	"context"

	"github.com/Darkpowercross/pocket-oni/internal/config/image"
	"github.com/Darkpowercross/pocket-oni/internal/config/information"
)

type Config struct {
	App               *App
	Character         *image.CharacterImages
	SpriteInformation information.SpriteMetaData
	Commands          []Command
}

func (c *Config) ApplyConfig() {

	c.App.Approot.Build()
	c.App.ApplyApp()
	c.Character.GetCharGifs()
	c.App.InputHandler()

	c.SpriteInformation = information.SpriteMetaData{
		State:    "Happy",
		Location: "Forest",
		Health:   100,
		Food:     100,
		Water:    100,
	}

	c.InitCommands()
	c.SetMenu("main")

	c.App.App.SetFocus(
		c.App.Approot.View.Content.Sub.Menu.Get(),
	)

	go c.AnimateTUI(context.Background())
}
