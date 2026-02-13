package information

import (
	"github.com/Darkpowercross/pocket-oni/internal/config/references"
	"github.com/gdamore/tcell/v2"
)

func (s *InformationMetaData) AddWaste(n float64) {
	curwaste := s.Waste
	if n <= 0 {
		n = 1
	}

	if curwaste < references.Maxwaste {
		if n+curwaste >= 100 {
			curwaste = 100
			s.Message.SetNewMessage(
				"Clean Barn",
				true,
				tcell.ColorIndianRed)
		}
	}
}

func (s *InformationMetaData) CleanWaste() {
	if s.Waste == 0 {
		s.Message.SetNewMessage("Nothing to clean!", true, tcell.ColorYellow)
		return
	}

	s.Waste = 0
	s.Comfort += 5
	s.Happiness += 3

	if s.Comfort > 100 {
		s.Comfort = 100
	}
	if s.Happiness > 100 {
		s.Happiness = 100
	}

	s.Message.SetNewMessage("Barn cleaned!", true, tcell.ColorGreen)
}
