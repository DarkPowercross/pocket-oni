package information

import (
	"fmt"

	"github.com/Darkpowercross/pocket-oni/internal/config/information/infotools"
	"github.com/Darkpowercross/pocket-oni/internal/config/references"
)

func (s *InformationMetaData) AdjustHappiness(delta int) {
	newValue := s.Happiness + delta

	if newValue > references.MaxHappiness {
		s.Happiness = references.MaxHappiness
		return
	}

	if newValue < 0 {
		s.Happiness = 0
		return
	}

	s.Happiness = newValue
}

func (s *InformationMetaData) GetHappiness() string {
	Preamble := infotools.GeneratePreamble("Happiness:", references.InformationIndent)
	Percent := fmt.Sprintf(" %d%s", s.Happiness, "%")
	msg := infotools.IntToBar(s.Happiness) + Percent

	color := references.Happiness

	return infotools.FormatStrings(Preamble, msg, color)
}

func (s *InformationMetaData) UpdateHappiness() {
	total :=
		statImpact(s.Food) +
			statImpact(s.Water) +
			statImpact(s.Comfort)

	healthModifier := HealthImpact(s.Health)
	modifier := total + (healthModifier)

	s.AdjustHappiness(modifier)
}

func statImpact(value int) int {
	switch {
	case value >= 80:
		return 1
	case value >= 60:
		return 0
	case value >= 40:
		return -1
	case value >= 20:
		return -2
	default:
		return -4
	}
}

func WaterImpact(value int) int {
	switch {
	case value >= 80:
		return 0
	case value >= 60:
		return -2
	case value >= 40:
		return -4
	case value >= 20:
		return -5
	default:
		return -6
	}
}

func HealthImpact(value int) int {
	switch {
	case value >= 80:
		return 0
	case value >= 60:
		return -1
	case value >= 40:
		return -2
	case value >= 20:
		return -3
	default:
		return -4
	}
}
