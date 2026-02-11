package infotools

import (
	"strings"
)

func GeneratePreamble(lable string, length int) string {
	SpaceTotal := length - len(lable)
	if SpaceTotal <= 0 {
		SpaceTotal = 1
	}
	Spaces := strings.Repeat(" ", SpaceTotal-1)

	Preamble := lable + Spaces + "| " + "%s%s%s"

	return Preamble
}
