package information

import (
	"github.com/Darkpowercross/pocket-oni/internal/config/appenums"
	"github.com/Darkpowercross/pocket-oni/internal/config/information/infotools"
	"github.com/Darkpowercross/pocket-oni/internal/config/references"
	"github.com/gdamore/tcell/v2"
)

func (s *InformationMetaData) SetLocation(location appenums.Locations) {
	s.Location = location
}

func (s *InformationMetaData) GetLocation() string {
	Preamble := infotools.GeneratePreamble("Location:", references.InformationIndent)
	msg := s.Location

	color := tcell.ColorForestGreen

	return infotools.FormatStrings(Preamble, msg, color)
}
