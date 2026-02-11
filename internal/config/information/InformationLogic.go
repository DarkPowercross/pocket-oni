package information

import "time"

func (s *InformationMetaData) InformationLogic() {
	for {
		s.UpdateHappiness()
		time.Sleep(time.Second)
	}
}
