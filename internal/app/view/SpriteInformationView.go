package view

import (
	"github.com/Darkpowercross/pocket-oni/internal/app/view/tools"
	"github.com/rivo/tview"
)

type SpriteInformationViewTypes int

const (
	SpriteLocation SpriteInformationViewTypes = iota
	SpriteHealth
	SpriteHappiness
	SpriteHunger
	SpriteWater
	SpriteDetails
)

type SpriteInformationView struct {
	InformationView     *tview.Flex
	SubInformationViews SubInformationViews
}

type SubInformationViews struct {
	SpriteLocation  *tview.TextView
	SpriteHappiness *tview.TextView
	SpriteHunger    *tview.TextView
	SpriteWater     *tview.TextView
	SpriteHealth    *tview.TextView
}

func (s *SpriteInformationView) SetSpriteInformationView(t *tools.Tools) {
	InformationOrder := []SpriteInformationViewTypes{
		SpriteDetails,
		SpriteLocation,
		SpriteHappiness,
		SpriteHunger,
		SpriteWater,
	}

	var InformationLayout = map[SpriteInformationViewTypes]LayoutProperties{
		SpriteDetails: {
			Border: true,
			Title:  "User details",
			Install: func(t *tools.Tools, p LayoutProperties) {
				s.InformationView = tview.NewFlex().SetDirection(tview.FlexRow)
			},
		},
		SpriteLocation: {
			Border: true,
			Title:  "Location",
			Install: func(t *tools.Tools, p LayoutProperties) {
				s.SubInformationViews.SpriteLocation = t.TextView(p.Border, p.Title)
				s.SubInformationViews.SpriteLocation.SetDynamicColors(true)
				s.InformationView.AddItem(s.SubInformationViews.SpriteLocation, 0, 1, true)

			},
		},
		SpriteHealth: {
			Border: true,
			Title:  "Health",
			Install: func(t *tools.Tools, p LayoutProperties) {
				s.SubInformationViews.SpriteHealth = t.TextView(p.Border, p.Title)
				s.SubInformationViews.SpriteLocation.SetDynamicColors(true)
				s.InformationView.AddItem(s.SubInformationViews.SpriteHealth, 0, 1, true)
			},
		},
		SpriteHappiness: {
			Border: true,
			Title:  "Happiness",
			Install: func(t *tools.Tools, p LayoutProperties) {
				s.SubInformationViews.SpriteHappiness = t.TextView(p.Border, p.Title)
				s.SubInformationViews.SpriteLocation.SetDynamicColors(true)
				s.InformationView.AddItem(s.SubInformationViews.SpriteHappiness, 0, 1, true)
			},
		},
		SpriteHunger: {
			Border: true,
			Title:  "Hunger",
			Install: func(t *tools.Tools, p LayoutProperties) {
				s.SubInformationViews.SpriteHunger = t.TextView(p.Border, p.Title)
				s.SubInformationViews.SpriteHunger.SetDynamicColors(true)
				s.InformationView.AddItem(s.SubInformationViews.SpriteHunger, 0, 1, true)
			},
		},
		SpriteWater: {
			Border: true,
			Title:  "Water",
			Install: func(t *tools.Tools, p LayoutProperties) {
				s.SubInformationViews.SpriteWater = t.TextView(p.Border, p.Title)
				s.SubInformationViews.SpriteHunger.SetDynamicColors(true)
				s.InformationView.AddItem(s.SubInformationViews.SpriteWater, 0, 1, true)
			},
		},
	}

	for _, l := range InformationOrder {
		layout, ok := InformationLayout[l]
		if !ok {
			continue
		}
		layout.Install(t, layout)
	}

}
