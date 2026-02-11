package information

import (
	"fmt"

	"github.com/Darkpowercross/pocket-oni/internal/config/information/infotools"
	"github.com/Darkpowercross/pocket-oni/internal/config/references"
)

func (s *InformationMetaData) SetHealth(n int) {
	s.Health = n
}

func (s *InformationMetaData) GetHealth() string {
	Preamble := infotools.GeneratePreamble("Health:", references.InformationIndent)
	Percent := fmt.Sprintf(" %d%s", s.Health, "%")
	msg := infotools.IntToBar(s.Health) + Percent

	color := references.HealthColor

	return infotools.FormatStrings(Preamble, msg, color)
}
