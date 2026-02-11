package information

import (
	"github.com/Darkpowercross/pocket-oni/internal/config/information/infotools"
	"github.com/gdamore/tcell/v2"
)

func (s *SpriteMetaData) SetHappiness(n int) {
	s.Happiness = n
}

func (s *SpriteMetaData) GetHappiness() string {
	Preamble := "Happiness: %s%d%s"
	msg := s.Happiness

	color := tcell.ColorBlue

	return infotools.FormatStrings(Preamble, msg, color)
}
