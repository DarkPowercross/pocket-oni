package tools

import "github.com/rivo/tview"

func Image(setborder bool, title string) *tview.Image {
	imageView := tview.NewImage()
	imageView.SetBorder(setborder)
	imageView.SetTitle(title)

	return imageView
}
