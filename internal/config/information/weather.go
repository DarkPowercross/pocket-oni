package information

import (
	"math/rand"
	"time"

	"github.com/Darkpowercross/pocket-oni/internal/config/appenums"
	"github.com/Darkpowercross/pocket-oni/internal/config/information/infotools"
	"github.com/Darkpowercross/pocket-oni/internal/config/references"
	"github.com/gdamore/tcell/v2"
)

type Weather struct {
	LastSet time.Time
	Name    appenums.WeatherType
}

var weatherPatterns = []appenums.WeatherType{
	appenums.Sunny,
	appenums.Rainy,
	appenums.Cloudy,
	appenums.Windy,
	appenums.Foggy,
}

var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

func (w *Weather) GetWeather() string {
	Preamble := infotools.GeneratePreamble("Weather:", references.HeaderIndent)
	msg := w.Name.String()
	color := tcell.ColorOrangeRed

	return infotools.FormatStrings(Preamble, msg, color)
}

func (w *Weather) SetWeather(name string) {
	now := time.Now()

	if name != "" {
		w.GetWeatherEnum(name)
		return
	}

	if w.LastSet.IsZero() && name != "" {
		w.GetWeatherEnum(name)
		w.LastSet = now
		return
	}

	if time.Since(w.LastSet) >= 5*time.Minute {
		if w.Name == appenums.Sunny && w.LastSet.IsZero() {
			w.LastSet = now
			return

		}
		randomWeather := weatherPatterns[rng.Intn(len(weatherPatterns))]
		w.Name = randomWeather
		w.LastSet = now
	}
}

func (w *Weather) GetWeatherEnum(name string) {
	now := time.Now()
	switch name {
	case "Sunny":
		w.Name = appenums.Sunny
		w.LastSet = now
		return
	case "Rainy":
		w.Name = appenums.Rainy
		w.LastSet = now
		return
	case "Cloudy":
		w.Name = appenums.Cloudy
		w.LastSet = now
		return
	case "Stormy":
		w.Name = appenums.Stormy
		w.LastSet = now
		return
	case "Windy":
		w.Name = appenums.Windy
		w.LastSet = now
		return
	case "Foggy":
		w.Name = appenums.Foggy
		w.LastSet = now
		return
	default:
		w.Name = appenums.Sunny
		w.LastSet = now
		return
	}
}

func (w *Weather) WeatherBuff() map[appenums.Locations]float64 {
	buffs := map[appenums.Locations]float64{
		appenums.Forest: 0,
		appenums.Ocean:  0,
		appenums.Barn:   0,
	}
	switch w.Name {
	case appenums.Sunny:
		buffs[appenums.Forest] = -5
		buffs[appenums.Ocean] = -2
		buffs[appenums.Barn] = 1
	case appenums.Rainy:
		buffs[appenums.Forest] = -5
		buffs[appenums.Ocean] = -4
		buffs[appenums.Barn] = 1.0
	case appenums.Cloudy:
		buffs[appenums.Forest] = -5
		buffs[appenums.Ocean] = -3
		buffs[appenums.Barn] = 1.0
	case appenums.Stormy:
		buffs[appenums.Forest] = -5
		buffs[appenums.Ocean] = -3
		buffs[appenums.Barn] = 1.0
	case appenums.Windy:
		buffs[appenums.Forest] = -2
		buffs[appenums.Ocean] = -1
		buffs[appenums.Barn] = 1.0
	case appenums.Foggy:
		buffs[appenums.Forest] = -2
		buffs[appenums.Ocean] = -1
		buffs[appenums.Barn] = 1.0
	default:
		buffs[appenums.Forest] = -5
		buffs[appenums.Ocean] = -2
		buffs[appenums.Barn] = 1.0
	}

	return buffs
}
