package information

import (
	"fmt"

	"github.com/Darkpowercross/pocket-oni/internal/config/information/infotools"
	"github.com/Darkpowercross/pocket-oni/internal/config/references"
	"github.com/gdamore/tcell/v2"
)

func (s *InformationMetaData) AddFood(Food int) {
	curfood := s.Food
	if Food <= 0 {
		Food = 1
	}

	if curfood < references.Maxfood {
		if Food+curfood <= 100 {
			curfood = 100
		}
	} else {
		s.Message.SetNewMessage(
			"Oni is full",
			true,
			references.Error,
		)
	}
}

func (s *InformationMetaData) GetFood() string {
	Preamble := infotools.GeneratePreamble("Food:", references.InformationIndent)
	Percent := fmt.Sprintf(" %d%s", s.Food, "%")
	msg := infotools.IntToBar(s.Food) + Percent

	color := tcell.ColorGreen

	return infotools.FormatStrings(Preamble, msg, color)
}
