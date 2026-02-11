package information

import (
	"github.com/Darkpowercross/pocket-oni/internal/config/information/infotools"
	"github.com/Darkpowercross/pocket-oni/internal/config/references"
	"github.com/gdamore/tcell/v2"
)

func (s *InformationMetaData) GetWeather() string {
	Preamble := infotools.GeneratePreamble("Weather:", references.HeaderIndent)
	msg := s.Weather

	color := tcell.ColorOrangeRed

	return infotools.FormatStrings(Preamble, msg, color)
}
