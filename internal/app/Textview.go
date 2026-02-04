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

func (a *App) FlexView(setborder bool, title string) *tview.Flex {
	Flexview := tview.NewFlex()
	Flexview.SetBorder(setborder)
	Flexview.SetTitle(title)
	return Flexview
}
