package glog

import (
	"fmt"
	"strings"
)

// IntAmount colors `n` as int and appends either the
// given singular (`n` == 1) or plural (`n` > 1 || `n` == 0).
func IntAmount[I Ints](n I, singular, plural string) string {
	unit := singular
	if n > 1 {
		unit = plural
	}
	color := LoggerConfig.ColorIntPositive
	if n < 0 {
		color = LoggerConfig.ColorIntNegative
	} else if n == 0 {
		color = LoggerConfig.ColorIntZero
	}
	amount := Wrap(fmt.Sprintf("%d", n), color)
	if n == 0 {
		unit = plural
	}
	return fmt.Sprintf("%s %s", amount, unit)
}

// Int colors the result cyan (`n` > 0), blue (`n` == 0) or red (`n` < 0).
func Int[I Ints](n ...I) string {
	res := []string{}
	for _, num := range n {
		color := LoggerConfig.ColorIntPositive
		if num < 0 {
			color = LoggerConfig.ColorIntNegative
		} else if num == 0 {
			color = LoggerConfig.ColorIntZero
		}
		res = append(res, Wrap(fmt.Sprintf("%d", num), color))
	}
	return strings.Join(res, ", ")
}
