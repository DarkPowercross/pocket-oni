package information

import (
	"fmt"

	"github.com/Darkpowercross/pocket-oni/internal/config/appenums"
	"github.com/Darkpowercross/pocket-oni/internal/config/information/infotools"
	"github.com/Darkpowercross/pocket-oni/internal/config/references"
	"github.com/gdamore/tcell/v2"
)

func (s *InformationMetaData) AddWaterMessage(delta float64) {
	msg := s.AddWater(delta)
	if msg != "" {
		s.Message.SetNewMessage(
			msg,
			true,
			references.Error,
		)
	}
}
func (s *InformationMetaData) AddWater(delta float64) string {
	newValue := s.Water + delta

	if newValue > references.Maxwater && s.Water >= references.Maxwater {
		return "Cannot water"
	} else if newValue > references.Maxwater {
		s.Water = references.Maxwater
		return ""
	}

	if newValue < 0 {
		s.Water = 0
		return ""
	}

	s.Water = newValue
	return ""
}

func (s *InformationMetaData) GetWater() string {
	Preamble := infotools.GeneratePreamble("Water:", references.InformationIndent)
	Percent := fmt.Sprintf(" %.0f%s", s.Water, "%")
	msg := infotools.IntToBar(s.Water) + Percent

	color := tcell.ColorBlue

	return infotools.FormatStrings(Preamble, msg, color)
}

func (s *InformationMetaData) UpdateWater() {
	waterlocationbuff := s.waterlocationbuff()
	weathBuff := s.Weather.WeatherBuff()[s.Location]
	s.AddWater(waterlocationbuff + weathBuff)

}

func (s *InformationMetaData) waterlocationbuff() float64 {
	switch s.Location {
	case appenums.Barn:
		return 0
	case appenums.Forest:
		return -1
	case appenums.Ocean:
		return -2
	}
	return 0
}
