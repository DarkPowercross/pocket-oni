package view

import (
	"github.com/Darkpowercross/pocket-oni/internal/app_new/view/tools"
	"github.com/rivo/tview"
)

type SpriteTypes int

const (
	Sprite SpriteTypes = iota
)

type SpriteView struct {
	Sprite *tview.Image
}

func (s SpriteView) Get() *tview.Image {
	return s.Sprite
}

func (s *SpriteView) SetSpriteView() {
	SpriteOrder := []SpriteTypes{
		Sprite,
	}

	var SpriteLayout = map[SpriteTypes]LayoutProperties{

		Sprite: {
			Border: true,
			Title:  "Oni",
			Install: func(p LayoutProperties) {
				s.Sprite = tools.Image(p.Border, p.Title)
			},
		},
	}

	for _, l := range SpriteOrder {
		layout, ok := SpriteLayout[l]
		if !ok {
			continue
		}
		layout.Install(layout)
	}
}
