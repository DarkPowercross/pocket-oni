package config

import (
	"github.com/Darkpowercross/pocket-oni/internal/config/view"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const (
	Maxfood     = 100
	Maxwater    = 100
	DefaultFont = tcell.ColorAntiqueWhite
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
