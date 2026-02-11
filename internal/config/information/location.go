package information

import (
	"github.com/Darkpowercross/pocket-oni/internal/config/information/infotools"
	"github.com/gdamore/tcell/v2"
)

func (s *SpriteMetaData) SetLocation(location string) {
	s.Location = location
}

func (s *SpriteMetaData) GetLocation() string {
	Preamble := "Location: %s%s%s"
	msg := s.Location

	color := tcell.ColorForestGreen

	return infotools.FormatStrings(Preamble, msg, color)
}
