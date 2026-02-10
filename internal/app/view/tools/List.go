package tools

import (
	"github.com/rivo/tview"
)

func (t *Tools) List(setborder bool, title string) *tview.List {
	list := tview.NewList()
	list.SetBorder(setborder)
	list.SetTitle(title)

	return list
}

