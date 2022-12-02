package glog

type Indicator struct {
	value string
	color int
}

func (i *Indicator) Wrap(message string) string {
	return Wrap(message, i.color)
}

func NewIndicator(value string, color int) *Indicator {
	return &Indicator{
		value: value,
		color: color,
	}
}
