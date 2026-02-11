package information

import (
	"github.com/Darkpowercross/pocket-oni/internal/config/information/infotools"
	"github.com/Darkpowercross/pocket-oni/internal/config/references"
	"github.com/gdamore/tcell/v2"
)

func (s *InformationMetaData) GetThought() string{
	Preamble := infotools.GeneratePreamble("Thought:", references.HeaderIndent)

	msg := ""

	color := tcell.ColorGreen

	return infotools.FormatStrings(Preamble, msg, color)
}
