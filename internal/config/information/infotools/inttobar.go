package infotools

import "strings"

func IntToBar(total int) string {
	filledBars := total / 10
	barval := "|"
	if filledBars == 10 {
		barval = "█"
	}
	bars := strings.Repeat(barval, filledBars)

	return bars
}
