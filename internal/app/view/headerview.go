package view

import (
	"github.com/Darkpowercross/pocket-oni/internal/app/view/tools"
	"github.com/rivo/tview"
)

type HeaderFields int

const (
	HeaderWeather HeaderFields = iota
	HeaderCharThought
	HeaderFeedback
	HeaderCharacter
	Header
)

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

func (h *HeaderView) SetHeaderView(t *tools.Tools) {
	Headerorder := []HeaderFields{Header, HeaderFeedback, HeaderWeather, HeaderCharacter, HeaderCharThought}

	var HeaderLayout = map[HeaderFields]LayoutProperties{
		Header: {
			Border: true,
			Title:  "",
			Install: func(t *tools.Tools, p LayoutProperties) {
				h.Header = t.FlexView(p.Border, p.Title)
				h.Header.SetDirection(tview.FlexColumn)
			},
		},
		HeaderFeedback: {
			Border: false,
			Title:  "",
			Install: func(t *tools.Tools, p LayoutProperties) {
				h.SubHeader.HeaderFeedback = t.TextView(p.Border, p.Title)
				h.SubHeader.HeaderFeedback.SetDynamicColors(true)
				h.Header.AddItem(h.SubHeader.HeaderFeedback, 0, 1, true)
			},
		},
		HeaderCharacter: {
			Border: false,
			Title:  "",
			Install: func(t *tools.Tools, p LayoutProperties) {
				h.SubHeader.HeaderCharacter = t.TextView(p.Border, p.Title)
				h.SubHeader.HeaderCharacter.SetDynamicColors(true)
				h.Header.AddItem(h.SubHeader.HeaderCharacter, 0, 1, true)
			},
		},
		HeaderWeather: {
			Border: false,
			Title:  "",
			Install: func(t *tools.Tools, p LayoutProperties) {
				h.SubHeader.HeaderWeather = t.TextView(p.Border, p.Title)
				h.SubHeader.HeaderWeather.SetDynamicColors(true)
				h.Header.AddItem(h.SubHeader.HeaderWeather, 0, 1, true)

			},
		},
		HeaderCharThought: {
			Border: false,
			Title:  "",
			Install: func(t *tools.Tools, p LayoutProperties) {
				h.SubHeader.HeaderThought = t.TextView(p.Border, p.Title)
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
		layout.Install(t, layout)
	}

}
