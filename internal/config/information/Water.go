package information

import (
	"fmt"

	"github.com/Darkpowercross/pocket-oni/internal/config/information/infotools"
	"github.com/Darkpowercross/pocket-oni/internal/config/references"
	"github.com/gdamore/tcell/v2"
)

func (s *InformationMetaData) AddWater(water int) {
	curwater := s.Water
	if water <= 0 {
		water = 1
	}

	if curwater < references.Maxwater {
		if water+curwater <= 100 {
			curwater = 100
		}
		curwater += water
	} else {
		s.Message.SetNewMessage(
			"Oni can't be watered",
			true,
			references.Error,
		)
	}
}

func (s *InformationMetaData) GetWater() string {
	Preamble := infotools.GeneratePreamble("Water:", references.InformationIndent)
	Percent := fmt.Sprintf(" %d%s", s.Water, "%")
	msg := infotools.IntToBar(s.Water) + Percent

	color := tcell.ColorBlue

	return infotools.FormatStrings(Preamble, msg, color)
}
