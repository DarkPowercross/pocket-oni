package view

import (
	"github.com/Darkpowercross/pocket-oni/internal/config/references"
	"github.com/Darkpowercross/pocket-oni/internal/config/view/tools"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type MenuTypes int

const (
	Menu MenuTypes = iota
)

type MenuView struct {
	Menu *tview.List
}

func (m MenuView) Get() *tview.List {
	return m.Menu
}

func (s *MenuView) SetMenuView() {
	Menuorder := []MenuTypes{
		Menu,
	}

	var MenuLayout = map[MenuTypes]LayoutProperties{
		Menu: {
			Border: true,
			Title:  "Menu",
			Install: func(p LayoutProperties) {
				s.Menu = tools.List(p.Border, p.Title)
				menuView := s.Menu
				menuView.SetMainTextColor(tcell.ColorAntiqueWhite)
				menuView.SetSecondaryTextColor(tcell.ColorBlue)
				menuView.SetBackgroundColor(references.BorderBackgrounds)
				menuView.SetSelectedBackgroundColor(tcell.ColorBlue)
				menuView.SetSelectedTextColor(references.BorderBackgrounds)
				menuView.SetBorderColor(references.BorderBackgrounds)
			},
		},
	}

	for _, l := range Menuorder {
		layout, ok := MenuLayout[l]
		if !ok {
			continue
		}
		layout.Install(layout)
	}
}
func (m *MenuView) Clear() {
	m.Menu.Clear()
}

func (m *MenuView) AddItem(
	label string,
	description string,
	shortcut rune,
	selected func(),
) {
	m.Menu.AddItem(label, description, shortcut, selected)
}
