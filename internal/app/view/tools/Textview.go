package tools

import (
	"github.com/rivo/tview"
)

func (t *Tools) TextView(setborder bool, title string) *tview.TextView {
	Textview := tview.NewTextView()
	Textview.SetBorder(setborder)
	Textview.SetTitle(title)

	return Textview
}
