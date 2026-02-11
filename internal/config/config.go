package config

import (
	"context"

	"github.com/Darkpowercross/pocket-oni/internal/config/image"
	"github.com/Darkpowercross/pocket-oni/internal/config/information"
	"github.com/Darkpowercross/pocket-oni/internal/config/references"
)

type Config struct {
	App               *App
	Character         *image.CharacterImages
	SpriteInformation information.InformationMetaData
	Commands          []Command
}

func (c *Config) ApplyConfig() {

	c.App.Approot.Build()
	c.App.ApplyApp()
	c.Character.GetCharGifs()
	c.App.InputHandler()

	c.SpriteInformation = information.InformationMetaData{
		State:     "Happy",
		Location:  "Forest",
		Health:    references.MaxHealth,
		Food:      references.Maxfood,
		Water:     references.Maxwaste,
		Waste:     references.MinWaste,
		Happiness: references.MaxHappiness,
		Comfort:   references.MaxComfort,
	}

	c.InitCommands()
	c.SetMenu("main")

	c.App.App.SetFocus(
		c.App.Approot.View.Content.Sub.Menu.Get(),
	)

	go c.AnimateTUI(context.Background())
	go c.SpriteInformation.InformationLogic()
}
