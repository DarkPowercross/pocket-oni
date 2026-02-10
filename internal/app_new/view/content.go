package view

import (
	"github.com/Darkpowercross/pocket-oni/internal/app_new/view/tools"
	"github.com/rivo/tview"
)

type Content struct {
	root *tview.Flex
	sub  subcontent
}

type subcontent struct {
	Menu              MenuView
	Sprite            SpriteView
	SpriteInformation SpriteInformationView
}

func (c *Content) SetContent() {
	root := tools.FlexView(false, "")
	root.SetDirection(tview.FlexColumn)

	c.sub.Menu.SetMenuView()
	c.sub.Sprite.SetSpriteView()
	c.sub.SpriteInformation.SetSpriteInformationView()

	root.AddItem(c.sub.Menu.Get(), 0, 1, false).
		AddItem(c.sub.Sprite.Get(), 0, 1, false).
		AddItem(c.sub.SpriteInformation.Get(), 0, 1, false)

	c.root = root
}

func (c *Content) Get() *tview.Flex {
	return c.root
}
