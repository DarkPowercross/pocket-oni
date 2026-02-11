package config

import "github.com/gdamore/tcell/v2"

func (a *App) InputHandler() {
	a.App.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEsc || event.Rune() == 'q' {
			a.App.Stop()
			return nil
		}
		return event
	})
}
