package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/Darkpowercross/pocket-oni/internal/app"
)

type Config struct {
	Health   int `yaml:"health"`
	Location int `yaml:"location"`
}

func main() {
	var a app.App
	a.Appstart()

	home, err := os.UserHomeDir()
	if err != nil {
		return
	}

	configDir := filepath.Join(home, ".pocket_oni/")
	configFile := filepath.Join(configDir, "config.conf")

	if err := os.MkdirAll(configDir, 0755); err != nil {
		fmt.Println("failed to create config directory:", err)
		return
	}

	_, err = os.ReadFile(configFile)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			err = os.WriteFile(configFile, []byte("Health:45\n"), 0644)
			if err != nil {
				fmt.Println("failed to create config file:", err)
			}
			return
		}

		fmt.Println("failed to read config:", err)
		return
	}

	Health := 100
	Location := "forest"
	SpriteHappiness := 100
	SpriteState := "Eating"

	a.Spriteinfo = &app.SpriteMetaData{
		Health:    Health,
		Location:  Location,
		Happiness: SpriteHappiness,
		State: SpriteState,
	}

	a.GetMoodGifs()

	ctx, cancel := context.WithCancel(context.Background())
	a.AnimCancel = cancel

	go a.AnimateGIF(ctx)

	go a.UpdateStats()

	if err := a.App.SetRoot(a.Approot, true).Run(); err != nil {
		fmt.Println("err: ", err)
		return
	}
}
