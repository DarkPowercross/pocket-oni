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

	// 2. Resolve state (fallback to Happy)
	state, ok := states[SpriteState]
	if !ok || state == nil {
		state = &c.Happy
	}

	// 3. Map location → *Mood
	locations := map[string]*Mood{
		"forest": &state.Forest,
		"ocean":  &state.Ocean,
		"barn":   &state.Barn,
	}

	// 4. Resolve location (fallback to forest)
	mood, ok := locations[SpriteLocation]
	if !ok || mood == nil {
		mood = &state.Forest
	}

	// 5. Safe delay handling
	var delay time.Duration
	if len(mood.Delay) > 0 {
		delay = time.Duration(mood.Delay[0])
	}

	return &mood.Images, delay
}
