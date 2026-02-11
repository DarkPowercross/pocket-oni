package view

import (
	"github.com/Darkpowercross/pocket-oni/internal/config/view/tools"
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
