package appenums

type Locations int

const (
	Forest Locations = iota
	Ocean
	Barn
)

func (l Locations) String() string {
	switch l {
	case Forest:
		return "Forest"
	case Ocean:
		return "Ocean"
	case Barn:
		return "Barn"
	default:
		return "Unknown"
	}
}
