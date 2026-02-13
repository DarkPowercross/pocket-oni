package information

import (
	"fmt"

	"github.com/Darkpowercross/pocket-oni/internal/config/information/infotools"
	"github.com/Darkpowercross/pocket-oni/internal/config/references"
)

func (s *InformationMetaData) AddHealthMessage(delta float64) {
	msg := s.AddHealth(delta)
	if msg != "" {
		s.Message.SetNewMessage(
			msg,
			true,
			references.Error,
		)
	}
}

func (s *InformationMetaData) AddHealth(delta float64) string {
	newValue := s.Health + delta

	if newValue > references.MaxHealth && s.Health >= references.MaxHealth {
		return ""
	} else if newValue > references.MaxHealth {
		s.Health = references.MaxHealth
		return ""
	}

	if newValue < 0 {
		s.Health = 0
		return ""
	}

	s.Health = newValue
	return ""
}

func (s *InformationMetaData) GetHealth() string {
	Preamble := infotools.GeneratePreamble("Health:", references.InformationIndent)
	Percent := fmt.Sprintf(" %.0f%s", s.Health, "%")
	msg := infotools.IntToBar(s.Health) + Percent

	color := references.HealthColor

	return infotools.FormatStrings(Preamble, msg, color)
}

func (s *InformationMetaData) UpdateHealth() {
	HealthFoodBuff := s.HealthFoodBuff()
	WaterHealthBuff := s.HealthWaterBuff()

	total := HealthFoodBuff + (WaterHealthBuff * 2)

	s.AddHealthMessage(total)
}

func (s *InformationMetaData) HealthFoodBuff() float64 {
	switch {
	case s.Food < 80:
		return -1
	case s.Food < 60:
		return -2
	case s.Food < 40:
		return -3
	case s.Food < 20:
		return -4
	case s.Food == 0:
		return -5
	default:
		return 1
	}
}

func (s *InformationMetaData) HealthWaterBuff() float64 {
	switch {
	case s.Water < 80:
		return -0.5
	case s.Water < 60:
		return -1
	case s.Water < 40:
		return -1.5
	case s.Water < 20:
		return -2
	case s.Water == 0:
		return -25
	default:
		return 1
	}
}
