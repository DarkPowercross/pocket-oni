package view

import (
	"github.com/Darkpowercross/pocket-oni/internal/config/references"
	"github.com/Darkpowercross/pocket-oni/internal/config/view/tools"
	"github.com/rivo/tview"
)

type AppRootView struct {
	Root *tview.Flex
	View View
}

func (a *AppRootView) Build() {
	a.View.Header.SetHeaderView()
	a.View.Content.SetContent()

	root := tools.FlexView(true, "")
	root.SetDirection(tview.FlexRow)
	root.SetBackgroundColor(references.BorderBackgrounds)

	root.
		AddItem(a.View.Header.root, 3, 1, false).
		AddItem(a.View.Content.Get(), 0, 1, false)

	a.Root = root
}
