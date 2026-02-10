package app

import (
	"context"
	"fmt"
	"image"
	"strings"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const (
	Maxfood     = 100
	Maxwater    = 100
	DefaultFont = tcell.ColorAntiqueWhite
)

type SpriteMetaData struct {
	Health    int
	Location  string
	Happiness int
	Food      int
	State     string
	Message   Message
	Water     int
}

type Message struct {
	LastSet      time.Time
	Flash        bool
	FlashCurrent bool
	Message      string
	Colour       tcell.Color
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
	// Header := a.View.Header.Header
	// Header.SetDirection(tview.FlexColumn)
	// Header.AddItem(a.View.HeaderFeedback, 0, 1, true)
	// Header.AddItem(a.View.HeaderCharacter, 0, 1, true)
	// Header.AddItem(a.View.HeaderThought, 0, 1, true)
	// Header.AddItem(a.View.HeaderWeather, 0, 1, true)
	content := tview.NewFlex().
		AddItem(a.View.Menu, 0, 1, true).
		AddItem(a.View.CharacterimageView, 0, 1, true).
		AddItem(a.View.Bar, 0, 1, true)

	a.View.Bar.AddItem(a.View.SpriteLocation, 3, 1, true)
	a.View.Bar.AddItem(a.View.SpriteHealth, 3, 1, true)
	a.View.Bar.AddItem(a.View.SpriteHappiness, 3, 1, true)
	a.View.Bar.AddItem(a.View.SpriteHunger, 3, 1, true)
	a.View.Bar.AddItem(a.View.SpriteWater, 3, 1, true)

	a.Approot = tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(a.View.Header.Header, 3, 1, false).
		AddItem(content, 0, 1, true)

}

func (a *App) UpdateBar() {
	header := a.View.Header.SubHeader
	header.HeaderThought.SetText(fmt.Sprintf("Thought: %s", "Oni is hungry"))
	header.HeaderWeather.SetText(fmt.Sprintf("Weather: %s", "Sunny"))

	a.View.SpriteLocation.SetText(a.FormatStrings("Location: %s%s%s", a.Spriteinfo.Location, DefaultFont))

	a.View.SpriteHealth.SetText(
		fmt.Sprintf("Health: %d", a.Spriteinfo.Health),
	)
	a.View.SpriteHappiness.SetText(
		fmt.Sprintf("Happiness: %d", a.Spriteinfo.Happiness),
	)
	a.UpdateWater()
	a.UpdateFood()
	a.UpdateMessage()

	a.View.Header.SubHeader.HeaderCharacter.SetText(
		fmt.Sprintf("Character: %s", a.Spriteinfo.State),
	)
}

func (a *App) GetCurrentImages() (*[]image.Image, time.Duration) {
	// 1. Map state → *Location
	states := map[string]*Location{
		"Happy":  &a.CharacterMood.Happy,
		"Sad":    &a.CharacterMood.Sad,
		"Sick":   &a.CharacterMood.Sick,
		"Angry":  &a.CharacterMood.Angry,
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
		a.Spriteinfo.UpdateHealth()
		a.Deadscreen()
	}
}

func (s *SpriteMetaData) UpdateHealth() {
	if s.Food > 0 && s.Health < 100 {
		if s.Health > 90 {
			health := 100 - s.Health
			s.Health += health
		} else {
			s.Health += 10
		}
	}

	if ((float64(s.Food) / float64(10)) * 100) <= 20 {
		s.Health -= 1
	}
	if s.Health <= 0 {
		s.State = "Dead"
	}
	if s.Food > 0 {
		s.Food -= 1
	}
	if s.Water > 0 {
		s.Water -= 1
	}
}

func (a *App) Deadscreen() {
	if a.Spriteinfo.State == "Dead" {
		a.Approot = tview.NewFlex().SetDirection(tview.FlexColumn)
		deadscreen := a.TextView(true, "")
		deadscreen.SetTextAlign(tview.AlignCenter)
		deadscreen.SetBorder(true)
		deadscreen.SetText("Oni is Dead")
		a.Approot.AddItem(deadscreen, 0, 1, true)
		if err := a.App.SetRoot(a.Approot, true).Run(); err != nil {
			fmt.Println("err: ", err)
			return
		}
	}
}

func (m *Message) ExpireIfNeeded() {
	if m.Message != "" && time.Since(m.LastSet) > 3*time.Second {
		m.Message = ""
	}
}

func (a *App) SetNewMessage(text string, flash bool, colour tcell.Color) {
	a.Spriteinfo.Message.Message = text
	a.Spriteinfo.Message.Flash = flash
	a.Spriteinfo.Message.Colour = colour
	a.Spriteinfo.Message.LastSet = time.Now()
}
func (a *App) UpdateMessage() {
	SubHeader := a.View.Header.SubHeader

	Preamble := "Message: %s%s%s"
	m := a.Spriteinfo.Message.Message

	a.Spriteinfo.Message.ExpireIfNeeded()

	var color tcell.Color

	if a.Spriteinfo.Message.Colour != tcell.ColorAntiqueWhite {
		color = a.Spriteinfo.Message.Colour
	}

	if a.Spriteinfo.Message.Flash == true {
		if a.Spriteinfo.Message.FlashCurrent == true {
			color = a.Spriteinfo.Message.Colour
			a.Spriteinfo.Message.FlashCurrent = false
		} else {
			color = tcell.ColorAntiqueWhite
			a.Spriteinfo.Message.FlashCurrent = true
		}
	}

	SubHeader.HeaderFeedback.SetText(
		a.FormatStrings(Preamble, m, color),
	)
}

func (a *App) FormatStrings(Preamble, Message string, Messagecolor tcell.Color) string {
	if Preamble == "" {
		return Message
	}

	return fmt.Sprintf(
		Preamble,
		"["+Messagecolor.String()+"]",
		Message,
		"[-]",
	)
}

func (a *App) UpdateFood() {
	Preamble := "Food: %s%s%s"

	bars := a.BarValue(a.Spriteinfo.Food)

	msg := a.FormatStrings(Preamble, bars, tcell.ColorGreen)

	a.View.SpriteHunger.SetText(msg)

}

func (a *App) BarValue(total int) string {
	filledBars := total / 10
	barval := "|"
	if filledBars == 10 {
		barval = "█"
	}
	bars := strings.Repeat(barval, filledBars)

	return bars
}

func (a *App) UpdateWater() {
	Preamble := "Water: %s%s%s"
	bars := a.BarValue(a.Spriteinfo.Water)

	msg := a.FormatStrings(Preamble, bars, tcell.ColorGreen)

	a.View.SpriteWater.SetText(msg)

}
