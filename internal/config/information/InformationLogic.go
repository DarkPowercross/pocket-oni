package information

import "time"

func (s *InformationMetaData) InformationLogic() {
	for {
		s.UpdateThought()

		time.Sleep(10 * time.Second)
		s.UpdateHappiness()
		s.Weather.SetWeather("")
		s.UpdateComfort()
		s.UpdateWater()
		s.UpdateFood()
		s.UpdateHealth()

	}
}
