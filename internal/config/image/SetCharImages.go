package image

import (
	"github.com/Darkpowercross/pocket-oni/internal/config/image/Imagetools"
	handlers "github.com/Darkpowercross/pocket-oni/internal/config/image/Imagetools"
)

func (c *CharacterImages) SetCharImages(chardir string, background string) {
	scaledFrames, delay := Imagetools.ScaleGifFrames(chardir)
	scaledBackgroundImage := handlers.ScalebackgroundImage(background)

	for i, img := range scaledFrames {
		scaledFrames[i] = handlers.GenerateGifBackground(img, scaledBackgroundImage)
	}

	mstruct := Mood{
		Images: scaledFrames,
		Delay:  delay,
	}

	states := map[string]*Location{
		"oni_happy.gif":  &c.Happy,
		"oni_angry.gif":  &c.Angry,
		"oni_sad.gif":    &c.Sad,
		"oni_sick.gif":   &c.Sick,
		"oni_eating.gif": &c.Eating,
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
