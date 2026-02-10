package app

type WeatherTypes int

const (
	Sunny WeatherTypes = iota
	Scorching
	Raining
	Storming
)

var Weather = []WeatherTypes{Sunny, Scorching, Raining, Storming}

type Weathers struct {
	Name  WeatherTypes
	Buffs Buffs
}

type Buffs struct {
	Barn   int
	Forest int
	Ocean  int
}

var Weathermap = map[WeatherTypes]Weathers{
	Sunny: {
		Name: Sunny,
		Buffs: Buffs{
			Barn:   100,
			Forest: 100,
			Ocean:  100,
		},
	},
}

// func (w *Weathers) GenerateWeather() Weathers{

// 	for item := range Weather{
// 		return Weathermap[item]
// 	}
// }


