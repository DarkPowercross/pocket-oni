package main

import (
	"fmt"

	appnew "github.com/Darkpowercross/pocket-oni/internal/config"
	"github.com/Darkpowercross/pocket-oni/internal/config/image"
)

func main() {
	app := &appnew.App{}

	c := &appnew.Config{
		App:       app,
		Character: &image.CharacterImages{},
	}

	c.ApplyConfig()
	if err := c.App.Run(); err != nil {
		fmt.Println("err: ", err)
		return
	}

}
