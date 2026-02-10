// package main

// import (
// 	"context"
// 	"fmt"

// 	"github.com/Darkpowercross/pocket-oni/internal/app"
// 	"github.com/gdamore/tcell/v2"
// )

// type Config struct {
// 	Health   int `yaml:"health"`
// 	Location int `yaml:"location"`
// }

// func main() {
// 	var a app.App
// 	a.Appstart()

// 	a.Spriteinfo = &app.SpriteMetaData{
// 		Health:    100,
// 		Location:  "forest",
// 		Happiness: 100,
// 		State:     "Happy",
// 		Water:     app.Maxwater,
// 		Food:      app.Maxfood,
// 	}
// 	a.SetNewMessage("Welcome", false, tcell.ColorBlue)

// 	a.GetMoodGifs()

// 	ctx, cancel := context.WithCancel(context.Background())
// 	a.AnimCancel = cancel

// 	go a.AnimateGIF(ctx)

// 	go a.UpdateStats()

// 	if err := a.App.SetRoot(a.Approot, true).Run(); err != nil {
// 		fmt.Println("err: ", err)
// 		return
// 	}
// }

package main

import (
	"fmt"

	appnew "github.com/Darkpowercross/pocket-oni/internal/app_new"
)

func main() {
	var c appnew.Config

	c.ApplyConfig()
	if err := c.App.Run(); err != nil {
		fmt.Println("err: ", err)
		return
	}

}
