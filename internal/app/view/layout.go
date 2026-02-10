package view

import (
	"github.com/Darkpowercross/pocket-oni/internal/app/view/tools"
)

type View struct {
	CharacterimageView SpriteView
	Menu               MenuView
	Header             HeaderView
	Bar                SpriteInformationView
}

type LayoutProperties struct {
	Title   string
	Border  bool
	Install func(t *tools.Tools, p LayoutProperties)
}

// func (a *App) ApplyLayout() {
// 	a.View.Header.SetHeaderView(a)
// 	for _, l := range Layoutorder {
// 		layout, ok := Layout[l]
// 		if !ok {
// 			continue
// 		}
// 		layout.Install(a, layout)
// 	}
// }
