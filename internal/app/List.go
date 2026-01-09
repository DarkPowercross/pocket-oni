package app

import (
	"github.com/rivo/tview"
)

func (a *App) List(setborder bool, title string) *tview.List {
	list := tview.NewList()
	list.SetBorder(setborder)
	list.SetTitle(title)

	return list
}

