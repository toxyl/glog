package glog

import (
	"fmt"
	"strings"
)

// Uint colors the given uints.
//
// Related config setting(s):
//
//  - `n`  > 0: `LoggerConfig.ColorUintPositive`
//  - `n` == 0: `LoggerConfig.ColorUintZero`
func Uint[U Uints](n ...U) string {
	res := []string{}
	for _, num := range n {
		color := LoggerConfig.ColorUintPositive
		if num == 0 {
			color = LoggerConfig.ColorUintZero
		}
		res = append(res, Wrap(fmt.Sprintf("%d", num), color))
	}
	return strings.Join(res, ", ")
}
