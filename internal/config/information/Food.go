package information

import (
	"fmt"

	"github.com/Darkpowercross/pocket-oni/internal/config/appenums"
	"github.com/Darkpowercross/pocket-oni/internal/config/information/infotools"
	"github.com/Darkpowercross/pocket-oni/internal/config/references"
	"github.com/gdamore/tcell/v2"
)

func (s *InformationMetaData) AddFoodMessage(delta float64) {
	msg := s.AddFood(delta)
	if msg != "" {
		s.Message.SetNewMessage(
			msg,
			true,
			references.Error,
		)
	}
}

func (s *InformationMetaData) AddFood(delta float64) string {
	newValue := s.Food + delta

	if newValue > references.Maxfood && s.Water >= references.Maxfood {
		return "Cannot Feed"
	} else if newValue > references.Maxwater {
		s.Food = references.Maxfood
		return ""
	}

	if newValue < 0 {
		s.Food = 0
		return ""
	}

	s.Food = newValue
	return ""
}

func (s *InformationMetaData) GetFood() string {
	Preamble := infotools.GeneratePreamble("Food:", references.InformationIndent)
	Percent := fmt.Sprintf(" %.0f%s", s.Food, "%")
	msg := infotools.IntToBar(s.Food) + Percent

	color := tcell.ColorGreen

	return infotools.FormatStrings(Preamble, msg, color)
}

func (s *InformationMetaData) UpdateFood() {
	locationbuff := s.FoodLocationBuff()
	WeathBuff := s.Weather.WeatherBuff()

	s.AddFood(locationbuff + WeathBuff[s.Location])
}

func (s *InformationMetaData) FoodLocationBuff() float64 {
	switch s.Location {
	case appenums.Barn:
		return -1
	case appenums.Ocean:
		return -2
	case appenums.Forest:
		return -3
	default:
		return -1
	}
}
