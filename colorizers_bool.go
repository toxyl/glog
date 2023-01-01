package glog

import "strings"

// Bool colors the given booleans, it will use `LoggerConfig.ColorBoolFalse` for `false` and `LoggerConfig.ColorBoolTrue` for `true`.
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
