package glog

import "strings"

// Bool colors the given booleans.
//
// Related config setting(s):
//
//  - `n` == false: `LoggerConfig.ColorBoolFalse`
//  - `n` == true:  `LoggerConfig.ColorBoolTrue`
func Bool(b ...bool) string {
	res := []string{}
	for _, bo := range b {
		color := LoggerConfig.ColorBoolFalse
		text := "false"
		if bo {
			color = LoggerConfig.ColorBoolTrue
			text = "true"
		}
		res = append(res, Wrap(text, color))
	}
	return strings.Join(res, ", ")
}
