package appenums

type Mood int

const (
	Happy Mood = iota
	Sad
	Sick
	Angry
	Eating
)

func (m Mood) String() string {
	switch m {
	case Happy:
		return "Happy"
	case Sad:
		return "Sad"
	case Sick:
		return "Sick"
	case Angry:
		return "Angry"
	case Eating:
		return "Eating"
	default:
		return "Unknown"
	}

}
