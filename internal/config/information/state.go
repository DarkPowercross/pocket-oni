package information

import (
	"github.com/Darkpowercross/pocket-oni/internal/config/information/infotools"
	"github.com/gdamore/tcell/v2"
)

func (s *SpriteMetaData) SetState(state string) {
	s.State = state
}

func (s *SpriteMetaData) GetState() string {
	Preamble := "State: %s%s%s"
	msg := s.State

	color := tcell.ColorLawnGreen

	return infotools.FormatStrings(Preamble, msg, color)
}
