package config

import (
	"github.com/Darkpowercross/pocket-oni/internal/config/view"
	"github.com/rivo/tview"
)

type App struct {
	App     *tview.Application
	Approot view.AppRootView
}

func (a *App) Run() error {
	return a.App.Run()
}

func (a *App) ApplyApp() {
	a.App = tview.NewApplication()
	a.App.SetRoot(a.Approot.Root, true)

}
