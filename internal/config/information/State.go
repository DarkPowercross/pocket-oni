package information

import (
	"github.com/Darkpowercross/pocket-oni/internal/config/appenums"
	"github.com/Darkpowercross/pocket-oni/internal/config/information/infotools"
	"github.com/Darkpowercross/pocket-oni/internal/config/references"
	"github.com/gdamore/tcell/v2"
)

func (s *InformationMetaData) SetState(state appenums.Mood) {
	s.State = state
}

func (s *InformationMetaData) GetState() string {
	Preamble := infotools.GeneratePreamble("State:", references.HeaderIndent)
	msg := s.State.String()

	color := tcell.ColorLawnGreen

	return infotools.FormatStrings(Preamble, msg, color)
}
