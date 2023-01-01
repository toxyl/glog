package glog

import (
	"fmt"
	"strings"
)

// Uint colors the result cyan (`n` > 0), blue (`n` == 0) or red (`n` < 0).
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
