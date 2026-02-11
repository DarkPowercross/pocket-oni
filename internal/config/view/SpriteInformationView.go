package view

import (
	"github.com/Darkpowercross/pocket-oni/internal/config/references"
	"github.com/Darkpowercross/pocket-oni/internal/config/view/tools"
	"github.com/gdamore/tcell/v2"
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
	SpriteComfort
)

type SpriteInformationView struct {
	InformationView     *tview.Flex
	SubInformationViews SubInformationViews
}

func (s SpriteInformationView) Get() *tview.Flex {
	return s.InformationView
}

type SubInformationViews struct {
	SpriteLocation  *tview.TextView
	SpriteHappiness *tview.TextView
	SpriteHunger    *tview.TextView
	SpriteWater     *tview.TextView
	SpriteHealth    *tview.TextView
	SpriteComfort   *tview.TextView
}

func (s *SpriteInformationView) SetSpriteInformationView() {
	InformationOrder := []SpriteInformationViewTypes{
		SpriteDetails,
		SpriteLocation,
		SpriteHealth,
		SpriteHappiness,
		SpriteHunger,
		SpriteWater,
		SpriteComfort,
	}

	var InformationLayout = map[SpriteInformationViewTypes]LayoutProperties{
		SpriteDetails: {
			Border: false,
			Title:  "User details",
			Install: func(p LayoutProperties) {
				s.InformationView = tview.NewFlex().SetDirection(tview.FlexRow)
				s.InformationView.SetBorder(p.Border)
				s.InformationView.SetBackgroundColor(tcell.ColorAntiqueWhite)
			},
		},
		SpriteLocation: {
			Border: true,
			Install: func(p LayoutProperties) {

				s.SubInformationViews.SpriteLocation = tools.TextView(p.Border, p.Title)
				s.addStatView(&s.SubInformationViews.SpriteLocation)
				s.InformationView.AddItem(s.SubInformationViews.SpriteLocation, 0, 1, true)
			},
		},
		SpriteHealth: {
			Border: true,
			Install: func(p LayoutProperties) {
				s.SubInformationViews.SpriteHealth = tools.TextView(p.Border, p.Title)
				s.addStatView(&s.SubInformationViews.SpriteHealth)
				s.InformationView.AddItem(s.SubInformationViews.SpriteHealth, 0, 1, true)
			},
		},
		SpriteHappiness: {
			Border: true,
			Install: func(p LayoutProperties) {
				s.SubInformationViews.SpriteHappiness = tools.TextView(p.Border, p.Title)
				s.addStatView(&s.SubInformationViews.SpriteHappiness)
				s.InformationView.AddItem(s.SubInformationViews.SpriteHappiness, 0, 1, true)
			},
		},
		SpriteHunger: {
			Border: true,
			Install: func(p LayoutProperties) {
				s.SubInformationViews.SpriteHunger = tools.TextView(p.Border, p.Title)
				s.addStatView(&s.SubInformationViews.SpriteHunger)
				s.InformationView.AddItem(s.SubInformationViews.SpriteHunger, 0, 1, true)
			},
		},
		SpriteWater: {
			Border: true,
			Install: func(p LayoutProperties) {
				s.SubInformationViews.SpriteWater = tools.TextView(p.Border, p.Title)
				s.addStatView(&s.SubInformationViews.SpriteWater)
				s.InformationView.AddItem(s.SubInformationViews.SpriteWater, 0, 1, true)
			},
		},
		SpriteComfort: {
			Border: true,
			Install: func(p LayoutProperties) {

				s.SubInformationViews.SpriteComfort = tools.TextView(p.Border, p.Title)
				s.addStatView(&s.SubInformationViews.SpriteComfort)
				s.InformationView.AddItem(s.SubInformationViews.SpriteComfort, 0, 1, true)
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

func (s *SpriteInformationView) addStatView(tv **tview.TextView) {
	*tv = tools.TextView(true, "")
	(*tv).SetDynamicColors(true)
	(*tv).SetBackgroundColor(references.BorderBackgrounds)
	(*tv).SetBorderColor(references.BorderBackgrounds)
	(*tv).SetBorderPadding(
		references.SetBorderPaddingTop,
		references.SetBorderPaddingBottom,
		references.SetBorderPaddingLeft,
		references.SetBorderPaddingRight,
	)
}
