package glog

import "strings"

// Bool colors the result green if `b` is true, else it colors it red.
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
