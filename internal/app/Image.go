package app

import (
	"image"
	"os"

	"github.com/Darkpowercross/tamagotchi/internal/handlers"
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
	Happy Location
	Sad   Location
	Sick  Location
	Angry Location
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
	ScalebackgroundImage := handlers.ScalebackgroundImage(background)

	for i, image := range scaledFrames {
		scaledFrames[i] = handlers.GenerateGifBackground(image, ScalebackgroundImage)
	}

	mstruct := Mood{
		Images: scaledFrames,
		Delay:  delay,
	}

	switch background {
	case "forest.png":
		{
			switch chardir {
			case "oni_happy.gif":
				a.CharacterMood.Happy.Forest = mstruct
			case "oni_angry.gif":
				a.CharacterMood.Angry.Forest = mstruct
			case "oni_sad.gif":
				a.CharacterMood.Sad.Forest = mstruct
			case "oni_sick.gif":
				a.CharacterMood.Sick.Forest = mstruct
			}
		}
	case "ocean.png":
		{
			switch chardir {
			case "oni_happy.gif":
				a.CharacterMood.Happy.Ocean = mstruct
			case "oni_angry.gif":
				a.CharacterMood.Angry.Ocean = mstruct
			case "oni_sad.gif":
				a.CharacterMood.Sad.Ocean = mstruct
			case "oni_sick.gif":
				a.CharacterMood.Sick.Ocean = mstruct
			}
		}
	}
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

// func (a *App) SetMoodLocation(location, background string, mood Mood){
// 	switch background{
// 	case "forest.png":
// 		switch location{
// 		case "happy":
// 			a.CharacterMood.Happy.Forest = mood
// 		case "angry":
// 			a.CharacterMood.Angry.Forest = mood
// 		case "sad":
// 			a.CharacterMood.Sad.Forest = mood
// 		case "sick":

// 	}
// 		a.CharacterMood.Happy.Barn = mood
// 	}
// }
