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
	Health    int
	Location  string
	Happiness int
	Food      int
	State     string
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
	Header             *tview.Flex
	Bar                *tview.Flex
	SpriteLocation     *tview.TextView
	SpriteHappiness    *tview.TextView
	SpriteHunger       *tview.TextView

	SpriteHealth *tview.TextView
}

func (a *App) Appstart() {
	a.App = tview.NewApplication()

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
	if len(*frames) == 0 {
		return
	}

	ticker := time.NewTicker(time.Duration(delay) * time.Millisecond * 5)
	defer ticker.Stop()

	i := 0

	for {
		select {
		case <-ctx.Done():
			return

		case <-ticker.C:
			frames, _ := a.GetCurrentImages()
			if len(*frames) == 0 {
				return
			}

			frame := (*frames)[i]
			i = (i + 1) % len(*frames)

			a.App.QueueUpdateDraw(func() {
				a.View.CharacterimageView.SetImage(frame)
				a.UpdateBar()
			})
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
	a.View.Bar.AddItem(a.View.SpriteHappiness, 3, 1, true)
	a.View.Bar.AddItem(a.View.SpriteHunger, 3, 1, true)

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
	a.View.SpriteHappiness.SetText(
		fmt.Sprintf("Happiness: %d", a.Spriteinfo.Happiness),
	)
	a.View.SpriteHunger.SetText(
		fmt.Sprintf("Food: %d", a.Spriteinfo.Food),
	)
}

func (a *App) GetCurrentImages() (*[]image.Image, time.Duration) {
	// 1. Map state → *Location
	states := map[string]*Location{
		"Happy": &a.CharacterMood.Happy,
		"Sad":   &a.CharacterMood.Sad,
		"Sick":  &a.CharacterMood.Sick,
		"Angry": &a.CharacterMood.Angry,
		"Eating": &a.CharacterMood.Eating,
	}

	// 2. Resolve state (fallback to Happy)
	state, ok := states[a.Spriteinfo.State]
	if !ok || state == nil {
		state = &a.CharacterMood.Happy
	}

	// 3. Map location → *Mood
	locations := map[string]*Mood{
		"forest": &state.Forest,
		"ocean":  &state.Ocean,
		"barn":   &state.Barn,
	}

	// 4. Resolve location (fallback to forest)
	mood, ok := locations[a.Spriteinfo.Location]
	if !ok || mood == nil {
		mood = &state.Forest
	}

	// 5. Safe delay handling
	var delay time.Duration
	if len(mood.Delay) > 0 {
		delay = time.Duration(mood.Delay[0])
	}

	return &mood.Images, delay
}

func (a *App) UpdateStats() {
	for {
		time.Sleep(time.Second * 10)
		a.Spriteinfo.Health -= 1
		a.Spriteinfo.State = "Sick"
	}
}
