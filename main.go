package main

import (
	"fmt"

	"github.com/Darkpowercross/pocket-oni/internal/config"
	"github.com/Darkpowercross/pocket-oni/internal/config/image"
)

func main() {
	app := &config.App{}

	

	c := &config.Config{
		App:       app,
		Character: &image.CharacterImages{},
	}

	c.ApplyConfig()
	if err := c.App.Run(); err != nil {
		fmt.Println("err: ", err)
		return
	}
}
