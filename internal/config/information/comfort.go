package information

import (
	"fmt"

	"github.com/Darkpowercross/pocket-oni/internal/config/information/infotools"
	"github.com/Darkpowercross/pocket-oni/internal/config/references"
)

func (s *InformationMetaData) GetComfort() string {
	Preamble := infotools.GeneratePreamble("Comfort:", references.InformationIndent)
	Percent := fmt.Sprintf(" %d%s", s.Water, "%")
	msg := infotools.IntToBar(s.Comfort) + Percent

	color := references.Comfort

	return infotools.FormatStrings(Preamble, msg, color)
}
