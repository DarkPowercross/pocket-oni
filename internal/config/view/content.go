package view

import (
	"github.com/Darkpowercross/pocket-oni/internal/config/view/tools"
	"github.com/rivo/tview"
)

type Content struct {
	root *tview.Flex
	Sub  Subcontent
}

type Subcontent struct {
	Menu              MenuView
	Sprite            SpriteView
	SpriteInformation SpriteInformationView
}

func (c *Content) SetContent() {
	root := tools.FlexView(false, "")
	root.SetDirection(tview.FlexColumn)

	c.Sub.Menu.SetMenuView()
	c.Sub.Sprite.SetSpriteView()
	c.Sub.SpriteInformation.SetSpriteInformationView()

	root.AddItem(c.Sub.Menu.Get(), 0, 1, false).
		AddItem(c.Sub.Sprite.Get(), 0, 1, false).
		AddItem(c.Sub.SpriteInformation.Get(), 0, 1, false)

	c.root = root
}

func (c *Content) Get() *tview.Flex {
	return c.root
}
