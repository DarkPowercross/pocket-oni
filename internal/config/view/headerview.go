package view

import (
	"github.com/Darkpowercross/pocket-oni/internal/config/references"
	"github.com/Darkpowercross/pocket-oni/internal/config/view/tools"
	"github.com/rivo/tview"
)

type HeaderView struct {
	root *tview.Flex

	HeaderFeedback  *tview.TextView
	HeaderCharacter *tview.TextView
	HeaderThought   *tview.TextView
	HeaderWeather   *tview.TextView
}

func (h *HeaderView) SetHeaderView() {
	root := tools.FlexView(true, "")
	root.SetDirection(tview.FlexColumn)
	root.SetBackgroundColor(references.BorderBackgrounds)
	root.SetBorderColor(references.BorderBackgrounds)
	

	h.HeaderFeedback = buildHeaderText()
	h.HeaderWeather = buildHeaderText()
	h.HeaderCharacter = buildHeaderText()
	h.HeaderThought = buildHeaderText()

	root.
		AddItem(h.HeaderFeedback, 0, 1, true).
		AddItem(h.HeaderWeather, 0, 1, true).
		AddItem(h.HeaderCharacter, 0, 1, true).
		AddItem(h.HeaderThought, 0, 1, true)

	h.root = root
}

func (h *HeaderView) Get() *tview.Flex {
	return h.root
}

func buildHeaderText() *tview.TextView {
	tv := tools.TextView(false, "")
	tv.SetDynamicColors(true)
	tv.SetBackgroundColor(references.BorderBackgrounds)
	return tv
}
