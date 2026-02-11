package image

import (
	"image"
	"time"
)

func (c *CharacterImages) GetCurrentImages(SpriteState, SpriteLocation string) (*[]image.Image, time.Duration) {
	// 1. Map state → *Location
	states := map[string]*Location{
		"Happy":  &c.Happy,
		"Sad":    &c.Sad,
		"Sick":   &c.Sick,
		"Angry":  &c.Angry,
		"Eating": &c.Eating,
	}

	
	state, ok := states[SpriteState]
	if !ok || state == nil {
		state = &c.Happy
	}

	locations := map[string]*Mood{
		"forest": &state.Forest,
		"ocean":  &state.Ocean,
		"barn":   &state.Barn,
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
