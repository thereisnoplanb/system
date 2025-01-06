package DayOfWeek

type Enum int

const (
	Sunday    Enum = 0
	Monday    Enum = 1
	Tuesday   Enum = 2
	Wednesday Enum = 3
	Thursday  Enum = 4
	Friday    Enum = 5
	Saturday  Enum = 6
)

var names map[Enum]string = map[Enum]string{
	Sunday:    "Sunday",
	Monday:    "Monday",
	Tuesday:   "Tuesday",
	Wednesday: "Wednesday",
	Thursday:  "Thursday",
	Friday:    "Friday",
	Saturday:  "Staurday",
}

func (enum Enum) String() string {
	if name, ok := names[enum]; ok {
		return name
	}
	return "Unknown"
}

func (enum Enum) IsDefined() bool {
	_, ok := names[enum]
	return ok
}

func GetValues() []Enum {
	return []Enum{
		Sunday,
		Monday,
		Tuesday,
		Wednesday,
		Thursday,
		Friday,
		Saturday,
	}
}
