package information

import (
	"github.com/Darkpowercross/pocket-oni/internal/config/information/infotools"
	"github.com/gdamore/tcell/v2"
)

func (s *SpriteMetaData) SetHealth(n int) {
	s.Health = n
}

func (s *SpriteMetaData) GetHealth() string {
	Preamble := "Health: %s%d%s"
	msg := s.Health

	color := tcell.ColorOrangeRed

	return infotools.FormatStrings(Preamble, msg, color)
}
