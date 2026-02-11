package tools

import "github.com/rivo/tview"

func FlexView(setborder bool, title string) *tview.Flex {
	Flexview := tview.NewFlex()
	Flexview.SetBorder(setborder)
	Flexview.SetTitle(title)

	return Flexview
}
