package information

import "github.com/Darkpowercross/pocket-oni/internal/config/appenums"

type InformationMetaData struct {
	Health    float64
	Food      float64
	Water     float64
	Comfort   float64
	Happiness float64

	Waste float64

	Location appenums.Locations
	Weather  Weather
	State    appenums.Mood

	Message Message
	Thought Message
}
