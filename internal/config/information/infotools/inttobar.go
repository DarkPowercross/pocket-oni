package infotools

import "strings"

func IntToBar(total float64) string {
	filledBars := total / 10
	barval := "|"
	if filledBars == 10 {
		barval = "█"
	}
	bars := strings.Repeat(barval, int(filledBars))

	return bars
}
