package view

import (
	"github.com/Darkpowercross/pocket-oni/internal/app/view/tools"
	"github.com/rivo/tview"
)

type MenuTypes int

const (
	Menu MenuTypes = iota
)

type MenuView struct {
	Menu *tview.List
}

func (s *MenuView) SetMenuView(t *tools.Tools) {
	Menuorder := []MenuTypes{
		Menu,
	}

	var MenuLayout = map[MenuTypes]LayoutProperties{
		Menu: {
			Border: true,
			Title:  "MenuAnimated",
			Install: func(t *tools.Tools, p LayoutProperties) {
				s.Menu = t.List(p.Border, p.Title)
			},
		},
	}

	for _, l := range Menuorder {
		layout, ok := MenuLayout[l]
		if !ok {
			continue
		}
		layout.Install(t, layout)
	}
}
