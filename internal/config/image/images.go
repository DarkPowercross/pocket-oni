package image

import (
	"image"
	"os"
)

type CharacterImages struct {
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

func (c *CharacterImages) GetCharGifs() {
	MoodGifs, err := os.ReadDir("internal/characters/oni/65x65")
	if err != nil {
		return
	}

	BackgroundImages, err := os.ReadDir("internal/characters/backgrounds")

	for _, g := range MoodGifs {
		for _, b := range BackgroundImages {
			c.SetCharImages(g.Name(), b.Name())
		}
	}

}
