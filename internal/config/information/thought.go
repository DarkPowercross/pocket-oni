package information

import (
	"github.com/Darkpowercross/pocket-oni/internal/config/references"
)

// func (s *InformationMetaData) GetThought() string {
// 	Preamble := infotools.GeneratePreamble("Thought:", references.HeaderIndent)

// 	msg := s.Thought

// 	color := tcell.ColorGreen

// 	return infotools.FormatStrings(Preamble, msg, color)
// }

func (s *InformationMetaData) UpdateThought() {

	text := RandomStatement()
	s.Thought.SetNewMessage(text, true, references.Good)
}

func RandomStatement() string {
	statement := references.Statements[rng.Intn(len(references.Statements))]
	return statement
}
