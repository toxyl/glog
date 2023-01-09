package glog

import (
	"fmt"
	"strings"
)

// IntAmount colors `n` as int and appends either the
// given singular (`n` == 1) or plural (`n` > 1 || `n` == 0).
//
// Related config setting(s):
//
//  - `n`  > 0: `LoggerConfig.ColorIntPositive`
//  - `n` == 0: `LoggerConfig.ColorIntZero`
//  - `n`  < 0: `LoggerConfig.ColorIntNegative`
func IntAmount[I IntOrUint](n I, singular, plural string) string {
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

// Int colors the given ints.
//
// Related config setting(s):
//
//  - `n`  > 0: `LoggerConfig.ColorIntPositive`
//  - `n` == 0: `LoggerConfig.ColorIntZero`
//  - `n`  < 0: `LoggerConfig.ColorIntNegative`
func Int[I IntOrUint](n ...I) string {
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
