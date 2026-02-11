package infotools

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
)

func FormatStrings(Preamble string, Message any, Messagecolor tcell.Color) string {
	msg := fmt.Sprint(Message)
	if Preamble == "" {
		return msg
	}

	return fmt.Sprintf(
		Preamble,
		"["+Messagecolor.String()+"]",
		Message,
		"[-]",
	)
}
