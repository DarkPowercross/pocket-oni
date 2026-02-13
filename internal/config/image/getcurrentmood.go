package image

import (
	"image"
	"time"

	"github.com/Darkpowercross/pocket-oni/internal/config/appenums"
)

func (c *CharacterImages) GetCurrentImages(SpriteState appenums.Mood, SpriteLocation appenums.Locations) (*[]image.Image, time.Duration) {
	// 1. Map state → *Location
	states := map[appenums.Mood]*Location{
		appenums.Happy:  &c.Happy,
		appenums.Sad:    &c.Sad,
		appenums.Sick:   &c.Sick,
		appenums.Angry:  &c.Angry,
		appenums.Eating: &c.Eating,
	}

	state, ok := states[SpriteState]
	if !ok || state == nil {
		state = &c.Happy
	}

	locations := map[appenums.Locations]*Mood{
		appenums.Forest: &state.Forest,
		appenums.Ocean:  &state.Ocean,
		appenums.Barn:   &state.Barn,
	}

	mood, ok := locations[SpriteLocation]
	if !ok || mood == nil {
		mood = &state.Forest
	}

	var delay time.Duration
	if len(mood.Delay) > 0 {
		delay = time.Duration(mood.Delay[0])
	}

	return &mood.Images, delay
}
