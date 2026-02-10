package view

import (
	"github.com/Darkpowercross/pocket-oni/internal/app_new/view/tools"
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

func (s SpriteInformationView) Get() *tview.Flex{
	return s.InformationView
}

type SubInformationViews struct {
	SpriteLocation  *tview.TextView
	SpriteHappiness *tview.TextView
	SpriteHunger    *tview.TextView
	SpriteWater     *tview.TextView
	SpriteHealth    *tview.TextView
}

func (s *SpriteInformationView) SetSpriteInformationView() {
	InformationOrder := []SpriteInformationViewTypes{
		SpriteDetails,
		SpriteHealth,
		SpriteLocation,
		SpriteHappiness,
		SpriteHunger,
		SpriteWater,
	}

	var InformationLayout = map[SpriteInformationViewTypes]LayoutProperties{
		SpriteDetails: {
			Border: true,
			Title:  "User details",
			Install: func(p LayoutProperties) {
				s.InformationView = tview.NewFlex().SetDirection(tview.FlexRow)
				s.InformationView.SetBorder(p.Border)
			},
		},
		SpriteLocation: {
			Border: true,
			Install: func(p LayoutProperties) {
				s.SubInformationViews.SpriteLocation = tools.TextView(p.Border, p.Title)
				s.SubInformationViews.SpriteLocation.SetDynamicColors(true)
				s.InformationView.AddItem(s.SubInformationViews.SpriteLocation, 3, 1, true)

			},
		},
		SpriteHealth: {
			Border: true,
			Install: func(p LayoutProperties) {
				s.SubInformationViews.SpriteHealth = tools.TextView(p.Border, p.Title)
				s.SubInformationViews.SpriteHealth.SetDynamicColors(true)
				s.InformationView.AddItem(s.SubInformationViews.SpriteHealth, 3, 1, true)
			},
		},
		SpriteHappiness: {
			Border: true,
			Install: func(p LayoutProperties) {
				s.SubInformationViews.SpriteHappiness = tools.TextView(p.Border, p.Title)
				s.SubInformationViews.SpriteLocation.SetDynamicColors(true)
				s.InformationView.AddItem(s.SubInformationViews.SpriteHappiness, 3, 1, true)
			},
		},
		SpriteHunger: {
			Border: true,
			Install: func(p LayoutProperties) {
				s.SubInformationViews.SpriteHunger = tools.TextView(p.Border, p.Title)
				s.SubInformationViews.SpriteHunger.SetDynamicColors(true)
				s.InformationView.AddItem(s.SubInformationViews.SpriteHunger, 3, 1, true)
			},
		},
		SpriteWater: {
			Border: true,
			Install: func(p LayoutProperties) {
				s.SubInformationViews.SpriteWater = tools.TextView(p.Border, p.Title)
				s.SubInformationViews.SpriteHunger.SetDynamicColors(true)
				s.InformationView.AddItem(s.SubInformationViews.SpriteWater, 3, 1, true)
			},
		},
	}

	for _, l := range InformationOrder {
		layout, ok := InformationLayout[l]
		if !ok {
			continue
		}
		layout.Install(layout)
	}

}
