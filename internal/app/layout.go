package app

import "github.com/rivo/tview"

type LayoutTypes int

type LayoutFields int

const (
	Menu LayoutFields = iota
	Sprite
	Title
	SpriteDetails
	SpriteLocation
	SpriteHealth
	SpriteHappiness
	SpriteHunger
)


type LayoutProperties struct {
	Layouttype LayoutTypes
	Title      string
	Border     bool
	Install    func(a *App, p LayoutProperties)
}

var Layoutorder = []LayoutFields{Title, Menu, Sprite, SpriteDetails, SpriteLocation, SpriteHealth, SpriteHappiness, SpriteHunger}

var Layout = map[LayoutFields]LayoutProperties{
	Menu: {
		Border: true,
		Title:  "Animated Sprite",
		Install: func(a *App, p LayoutProperties) {
			a.View.Menu = a.List(p.Border, p.Title)
		},
	},
	Sprite: {
		Border: true,
		Title:  "Menu",
		Install: func(a *App, p LayoutProperties) {
			a.View.CharacterimageView = a.Image(p.Border, p.Title)
		},
	},
	Title: {
		Border: true,
		Title:  "This is a title",
		Install: func(a *App, p LayoutProperties) {
			a.View.Header = a.FlexView(p.Border, p.Title).
			AddItem(a.TextView(true, p.Title), 3, 1, true)
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
		},
	},
}

func (a *App) ApplyLayout() {
	for _, l := range Layoutorder {
		layout, ok := Layout[l]
		if !ok {
			continue
		}
		layout.Install(a, layout)
	}
}
