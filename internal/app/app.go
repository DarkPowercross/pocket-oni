package app

import (
	"context"
	"fmt"
	"image"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type SpriteMetaData struct {
	Health   int
	Location string
}

type App struct {
	App           *tview.Application
	View          view
	Approot       *tview.Flex
	Spriteinfo    *SpriteMetaData
	AnimCancel    context.CancelFunc
	CharacterMood ImageMood
	Commands      []Command
}

type view struct {
	CharacterimageView *tview.Image
	Menu               *tview.List
	Header             *tview.TextView
	Bar                *tview.Flex
	SpriteLocation     *tview.TextView

	SpriteHealth       *tview.TextView
}

// create an tview application varqff
func (a *App) Appstart() {
	a.App = tview.NewApplication()
	// a.SetCharacter(true, "Animated Sprite")

	a.ApplyLayout()
	a.SetCommands()
	a.SetMenu("main")
	a.InputHandler()
	a.ApplyRoot()

}

func (a *App) InputHandler() {
	a.App.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEsc || event.Rune() == 'q' {
			a.App.Stop()
			return nil
		}
		return event
	})
}

func (a *App) AnimateGIF(ctx context.Context) {
	frames, delay := a.GetCurrentImages()
	if len(frames) == 0 {
		return
	}

	ticker := time.NewTicker(delay)
	defer ticker.Stop()

	i := 0

	for {
		select {
		case <-ctx.Done():
			return

		case <-ticker.C:
			i = (i + 1) % len(frames)

			a.App.QueueUpdateDraw(func() {
				a.View.CharacterimageView.SetImage(frames[i])
				a.UpdateBar()
			})

			time.Sleep(20 * time.Millisecond)
		}
	}
}

func (a *App) ApplyRoot() {

	content := tview.NewFlex().
		AddItem(a.View.Menu, 0, 1, true).
		AddItem(a.View.CharacterimageView, 0, 1, true).
		AddItem(a.View.Bar, 0, 1, true)

	a.View.Bar.AddItem(a.View.SpriteLocation, 3, 1, true)
	a.View.Bar.AddItem(a.View.SpriteHealth, 3, 1, true)

	a.Approot = tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(a.View.Header, 3, 1, false).
		AddItem(content, 0, 1, true)

}

func (a *App) UpdateBar() {
	a.View.SpriteLocation.SetText(
		fmt.Sprintf("Location: %s", a.Spriteinfo.Location),
	)

	a.View.SpriteHealth.SetText(
		fmt.Sprintf("Health: %d", a.Spriteinfo.Health),
	)
}

func (a *App) GetCurrentImages() ([]image.Image, time.Duration) {
	switch a.Spriteinfo.Location {
	case "forest":
		return a.CharacterMood.Happy.Forest.Images, time.Duration(a.CharacterMood.Happy.Forest.Delay[0])
	case "Ocean":
		return a.CharacterMood.Happy.Ocean.Images, time.Duration(a.CharacterMood.Happy.Ocean.Delay[0])
	case "Barn":
		return a.CharacterMood.Happy.Barn.Images, time.Duration(a.CharacterMood.Happy.Barn.Delay[0])

	default:
		return a.CharacterMood.Happy.Forest.Images, time.Duration(a.CharacterMood.Happy.Forest.Delay[0])
	}
}
