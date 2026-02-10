package view

import (
	"github.com/Darkpowercross/pocket-oni/internal/app/view/tools"
	"github.com/rivo/tview"
)

type SpriteTypes int

const (
	Sprite SpriteTypes = iota
)

type SpriteView struct {
	Sprite *tview.Image
}

func (s SpriteView) SetSpriteView(t *tools.Tools) {
	SpriteOrder := []SpriteTypes{
		Sprite,
	}

	var SpriteLayout = map[SpriteTypes]LayoutProperties{

		Sprite: {
			Border: true,
			Title:  "Oni",
			Install: func(t *tools.Tools, p LayoutProperties) {
				s.Sprite = t.Image(p.Border, p.Title)
			},
		},
	}

	for _, l := range SpriteOrder {
		layout, ok := SpriteLayout[l]
		if !ok {
			continue
		}
		layout.Install(t, layout)
	}
}
