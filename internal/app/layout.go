package app

import (
	"github.com/rivo/tview"
)

type LayoutTypes int

type LayoutFields int

type HeaderFields int

const (
	Menu LayoutFields = iota
	Sprite
	SpriteDetails
	SpriteLocation
	SpriteHealth
	SpriteHappiness
	SpriteHunger
	SpriteWater
)

const (
	HeaderWeather HeaderFields = iota
	HeaderCharThought
	HeaderFeedback
	HeaderCharacter
	Header
)

type view struct {
	CharacterimageView *tview.Image
	Menu               *tview.List
	Header             HeaderView
	// HeaderFeedback     *tview.TextView
	// HeaderCharacter    *tview.TextView
	Bar             *tview.Flex
	SpriteLocation  *tview.TextView
	SpriteHappiness *tview.TextView
	SpriteHunger    *tview.TextView
	SpriteWater     *tview.TextView
	SpriteHealth    *tview.TextView
	// HeaderThought      *tview.TextView
	// HeaderWeather      *tview.TextView
}

type HeaderView struct {
	Header    *tview.Flex
	SubHeader SubHeaders
}
type SubHeaders struct {
	HeaderFeedback  *tview.TextView
	HeaderCharacter *tview.TextView
	HeaderThought   *tview.TextView
	HeaderWeather   *tview.TextView
}

type LayoutProperties struct {
	Layouttype LayoutTypes
	Title      string
	Border     bool
	Install    func(a *App, p LayoutProperties)
}

var Layoutorder = []LayoutFields{Menu, Sprite, SpriteDetails, SpriteLocation, SpriteHealth, SpriteHappiness, SpriteHunger, SpriteWater}

var Layout = map[LayoutFields]LayoutProperties{
	Menu: {
		Border: true,
		Title:  "MenuAnimated",
		Install: func(a *App, p LayoutProperties) {
			a.View.Menu = a.List(p.Border, p.Title)
		},
	},
	Sprite: {
		Border: true,
		Title:  "Oni",
		Install: func(a *App, p LayoutProperties) {
			a.View.CharacterimageView = a.Image(p.Border, p.Title)
		},
	},
	SpriteDetails: {
		Border: true,
		Title:  "User details",
		Install: func(a *App, p LayoutProperties) {
			a.View.Bar = tview.NewFlex().SetDirection(tview.FlexRow)
		},
	},
	SpriteLocation: {
		Border: true,
		Title:  "Location",
		Install: func(a *App, p LayoutProperties) {
			a.View.SpriteLocation = a.TextView(p.Border, p.Title)
			a.View.SpriteLocation.SetDynamicColors(true)
		},
	},
	SpriteHealth: {
		Border: true,
		Title:  "Health",
		Install: func(a *App, p LayoutProperties) {
			a.View.SpriteHealth = a.TextView(p.Border, p.Title)
		},
	},
	SpriteHappiness: {
		Border: true,
		Title:  "Happiness",
		Install: func(a *App, p LayoutProperties) {
			a.View.SpriteHappiness = a.TextView(p.Border, p.Title)
		},
	},
	SpriteHunger: {
		Border: true,
		Title:  "Hunger",
		Install: func(a *App, p LayoutProperties) {
			a.View.SpriteHunger = a.TextView(p.Border, p.Title)
			a.View.SpriteHunger.SetDynamicColors(true)
		},
	},
	SpriteWater: {
		Border: true,
		Title:  "Water",
		Install: func(a *App, p LayoutProperties) {
			a.View.SpriteWater = a.TextView(p.Border, p.Title)
			a.View.SpriteWater.SetDynamicColors(true)
		},
	},
}

func (a *App) ApplyLayout() {
	a.View.Header.SetHeaderView(a)
	for _, l := range Layoutorder {
		layout, ok := Layout[l]
		if !ok {
			continue
		}
		layout.Install(a, layout)
	}
}

func (h *HeaderView) SetHeaderView(a *App) {
	Headerorder := []HeaderFields{Header, HeaderFeedback, HeaderWeather, HeaderCharacter, HeaderCharThought}

	var HeaderLayout = map[HeaderFields]LayoutProperties{
		Header: {
			Border: true,
			Title:  "",
			Install: func(a *App, p LayoutProperties) {
				h.Header = a.FlexView(p.Border, p.Title)
				h.Header.SetDirection(tview.FlexColumn)
			},
		},
		HeaderFeedback: {
			Border: false,
			Title:  "",
			Install: func(a *App, p LayoutProperties) {
				h.SubHeader.HeaderFeedback = a.TextView(p.Border, p.Title)
				h.SubHeader.HeaderFeedback.SetDynamicColors(true)
				h.Header.AddItem(h.SubHeader.HeaderFeedback, 0, 1, true)
			},
		},
		HeaderCharacter: {
			Border: false,
			Title:  "",
			Install: func(a *App, p LayoutProperties) {
				h.SubHeader.HeaderCharacter = a.TextView(p.Border, p.Title)
				h.SubHeader.HeaderCharacter.SetDynamicColors(true)
				h.Header.AddItem(h.SubHeader.HeaderCharacter, 0, 1, true)
			},
		},
		HeaderWeather: {
			Border: false,
			Title:  "",
			Install: func(a *App, p LayoutProperties) {
				h.SubHeader.HeaderWeather = a.TextView(p.Border, p.Title)
				h.SubHeader.HeaderWeather.SetDynamicColors(true)
				h.Header.AddItem(h.SubHeader.HeaderWeather, 0, 1, true)

			},
		},
		HeaderCharThought: {
			Border: false,
			Title:  "",
			Install: func(a *App, p LayoutProperties) {
				h.SubHeader.HeaderThought = a.TextView(p.Border, p.Title)
				h.SubHeader.HeaderThought.SetDynamicColors(true)
				h.Header.AddItem(h.SubHeader.HeaderThought, 0, 1, true)
			},
		},
	}

	for _, l := range Headerorder {
		layout, ok := HeaderLayout[l]
		if !ok {
			continue
		}
		layout.Install(a, layout)
	}

}
