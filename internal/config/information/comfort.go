package information

import (
	"fmt"

	"github.com/Darkpowercross/pocket-oni/internal/config/appenums"
	"github.com/Darkpowercross/pocket-oni/internal/config/information/infotools"
	"github.com/Darkpowercross/pocket-oni/internal/config/references"
)

func (s *InformationMetaData) AddComfortMessage(delta float64) {
	msg := s.AddComfort(delta)
	if msg != "" {
		s.Message.SetNewMessage(
			msg,
			true,
			references.Error,
		)
	}
}
func (s *InformationMetaData) AddComfort(delta float64) string {
	newValue := s.Comfort + delta

	if newValue > references.MaxComfort && s.Comfort >= references.MaxComfort {
		return "Feeling Uncomfortable"
	} else if newValue > references.MaxComfort {
		s.Comfort = references.MaxComfort
		return ""
	}

	if newValue < 0 {
		s.Comfort = 0
		return ""
	}

	s.Comfort = newValue
	return ""
}

func (s *InformationMetaData) GetComfort() string {
	Preamble := infotools.GeneratePreamble("Comfort:", references.InformationIndent)
	Percent := fmt.Sprintf(" %.0f%s", s.Comfort, "%")
	msg := infotools.IntToBar(s.Comfort) + Percent

	color := references.Comfort

	return infotools.FormatStrings(Preamble, msg, color)
}

func (s *InformationMetaData) SetComfort(delta float64) {
	newValue := s.Comfort + delta

	if newValue > references.MaxComfort {
		s.Comfort = references.MaxComfort
		return
	}

	if newValue < 0 {
		s.Comfort = 0
		return
	}

	s.Comfort = newValue
}

func (s *InformationMetaData) UpdateComfort() {
	wastebuff := 0.0
	locationbuff := float64(s.Weather.WeatherBuff()[s.Location])
	if s.Location == appenums.Barn {
		wastebuff = s.WasteBuff()
	}

	total := wastebuff + locationbuff
	s.AddComfort(total)

}

func (s *InformationMetaData) WasteBuff() float64 {
	value := s.Waste
	switch {
	case value >= 80:
		return -3
	case value >= 60:
		return -2
	case value >= 40:
		return -1
	case value >= 20:
		return 0
	default:
		return 1
	}
}
