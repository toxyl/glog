package indicator

type Indicator struct {
	Value string
	Color int
}

func NewIndicator(value string, color int) *Indicator {
	return &Indicator{
		Value: value,
		Color: color,
	}
}
