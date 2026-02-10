package appnew

import (
	"github.com/Darkpowercross/pocket-oni/internal/app_new/view"
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

type Config struct {
	App App
}

func (a *App) Run() error{
	return a.App.Run()
}

func (a *App) ApplyApp() {
	a.App = tview.NewApplication()
	a.App.SetRoot(a.Approot.Root, true)
}

func (c *Config) ApplyConfig() {
	c.App.Approot.Build()
	c.App.ApplyApp()
}
