package information

import (
	"time"

	"github.com/Darkpowercross/pocket-oni/internal/config/information/infotools"
	"github.com/Darkpowercross/pocket-oni/internal/config/references"
	"github.com/gdamore/tcell/v2"
)

type Message struct {
	Message     string
	Colour      tcell.Color
	Flash       bool
	FlashPeriod time.Duration
	flashStart  time.Time
	LastSet     time.Time
}

func (m *Message) SetNewMessage(text string, flash bool, colour tcell.Color) {
	m.Message = text
	m.Flash = flash
	m.Colour = colour
	m.LastSet = time.Now()

	if m.Flash && m.FlashPeriod == 0 {
		m.FlashPeriod = 100 * time.Millisecond
	}

	if flash {
		m.StartFlash()
	}
}

func (m *Message) StartFlash() {
	if m.Flash {
		m.flashStart = time.Now()
	}
}

func (m *Message) GetMessage() string {
	m.ExpireIfNeeded()

	Preamble := infotools.GeneratePreamble("Message:", references.HeaderIndent)
	msg := m.Message

	color := references.DefaultFontColor
	if m.Colour != references.DefaultFontColor {
		color = m.Colour
	}

	if m.Flash {
		if m.flashStart.IsZero() {
			m.flashStart = time.Now()
		}

		elapsed := time.Since(m.flashStart)
		if (elapsed/m.FlashPeriod)%2 == 1 {
			color = tcell.ColorAntiqueWhite
		} else {
			color = m.Colour
		}
	}

	return infotools.FormatStrings(Preamble, msg, color)
}

func (m *Message) ExpireIfNeeded() bool {
	if m.Message != "" && time.Since(m.LastSet) > 3*time.Second {
		m.Message = ""
		return true
	}
	return false
}
