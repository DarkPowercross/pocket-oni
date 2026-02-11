package information

import (
	"github.com/Darkpowercross/pocket-oni/internal/config/information/infotools"
	"github.com/Darkpowercross/pocket-oni/internal/config/references"
	"github.com/gdamore/tcell/v2"
)

func (s *SpriteMetaData) AddFood(Food int) {
	curfood := s.Food
	if Food <= 0 {
		Food = 1
	}

	if curfood < references.Maxfood {
		if Food+curfood <= 100 {
			curfood = 100
		}
		curfood += Food
	} else {
		s.Message.SetNewMessage(
			"Oni is full",
			true,
			tcell.ColorIndianRed,
		)
	}
}

func (s *SpriteMetaData) GetFood() string {
	Preamble := "Food: %s%s%s"
	msg := infotools.IntToBar(s.Food)

	color := tcell.ColorGreen

	return infotools.FormatStrings(Preamble, msg, color)
}
