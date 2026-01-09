package main

import (
	"context"
	"fmt"

	"github.com/Darkpowercross/tamagotchi/internal/app"
)

func main() {
	var a app.App
	a.Appstart()

	Health := 45
	Location := "forest"

	a.Spriteinfo = &app.SpriteMetaData{
		Health:   Health,
		Location: Location,
	}

	// scaledFrames, delay := handlers.ScaleFrames("oni_happy.gif")

	// Moodnames := []string{"happy", "sad", "sick", "angry"}
	// for _, mood := range Moodnames {
	// 	a.GetMoodImages(mood)
	// }

	a.GetMoodGifs()

	ctx, cancel := context.WithCancel(context.Background())
	a.AnimCancel = cancel

	go a.AnimateGIF(ctx)

	if err := a.App.SetRoot(a.Approot, true).Run(); err != nil {
		fmt.Println("err: ", err)
		return
	}
}
