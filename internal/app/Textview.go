package app

import (
	"github.com/rivo/tview"
)

func (a *App) TextView(setborder bool, title string) *tview.TextView {
	Textview := tview.NewTextView()
	Textview.SetBorder(setborder)
	Textview.SetTitle(title)

	return Textview
}
