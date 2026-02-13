package appenums

type WeatherType int

const (
	Sunny WeatherType = iota
	Rainy
	Cloudy
	Stormy
	Windy
	Foggy
)

func (w WeatherType) String() string {
	switch w {
	case Sunny:
		return "Sunny"
	case Rainy:
		return "Rainy"
	case Cloudy:
		return "Cloudy"
	case Stormy:
		return "Stormy"
	case Windy:
		return "Windy"
	case Foggy:
		return "Foggy"
	default:
		return "Unknown"
	}
}
