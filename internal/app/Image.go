package app

import (
	"image"
	"os"

	"github.com/Darkpowercross/pocket-oni/internal/handlers"
	"github.com/rivo/tview"
)

// returns a tview image variable
func (a *App) Image(setborder bool, title string) *tview.Image {
	imageView := tview.NewImage()
	imageView.SetBorder(setborder)
	imageView.SetTitle(title)

	return imageView
}

func (a *App) SetCharacter(setborder bool, title string) {
	a.View.CharacterimageView = a.Image(setborder, title)
}

type ImageMood struct {
	Happy  Location
	Sad    Location
	Sick   Location
	Angry  Location
	Eating Location
}

type Location struct {
	Forest Mood
	Ocean  Mood
	Barn   Mood
}

type Mood struct {
	Images []image.Image
	Delay  []int
}

func (a *App) GetMoodImages(chardir string, background string) {
	scaledFrames, delay := handlers.ScaleFrames(chardir)
	scaledBackgroundImage := handlers.ScalebackgroundImage(background)

	for i, img := range scaledFrames {
		scaledFrames[i] = handlers.GenerateGifBackground(img, scaledBackgroundImage)
	}

	mstruct := Mood{
		Images: scaledFrames,
		Delay:  delay,
	}

	states := map[string]*Location{
		"oni_happy.gif": &a.CharacterMood.Happy,
		"oni_angry.gif": &a.CharacterMood.Angry,
		"oni_sad.gif":   &a.CharacterMood.Sad,
		"oni_sick.gif":  &a.CharacterMood.Sick,
		"oni_eating.gif": &a.CharacterMood.Eating,
	}

	state, ok := states[chardir]
	if !ok || state == nil {
		return
	}

	locations := map[string]*Mood{
		"forest.png": &state.Forest,
		"ocean.png":  &state.Ocean,
		"barn.png":   &state.Barn,
	}

	loc, ok := locations[background]
	if !ok || loc == nil {
		return
	}
	*loc = mstruct
}

func (a *App) GetMoodGifs() {
	MoodGifs, err := os.ReadDir("internal/characters/oni/65x65")
	if err != nil {
		return
	}

	BackgroundImages, err := os.ReadDir("internal/characters/backgrounds")

	for _, g := range MoodGifs {
		for _, b := range BackgroundImages {
			a.GetMoodImages(g.Name(), b.Name())
		}
	}

}
