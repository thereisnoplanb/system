package Month

type Enum int

const (
	January Enum = 1 + iota
	February
	March
	April
	May
	June
	July
	August
	September
	October
	November
	December
)

var names map[Enum]string = map[Enum]string{
	January:   "January",
	February:  "February",
	March:     "March",
	April:     "April",
	May:       "May",
	June:      "June",
	July:      "July",
	August:    "August",
	September: "September",
	October:   "October",
	November:  "November",
	December:  "December",
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
		January,
		February,
		March,
		April,
		May,
		June,
		July,
		August,
		September,
		October,
		November,
		December,
	}
}
