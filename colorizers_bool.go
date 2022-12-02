package glog

// Bool colors the result green if `b` is true, else it colors it red.
func Bool(b bool) string {
	color := LoggerConfig.ColorBoolFalse
	text := "false"
	if b {
		color = LoggerConfig.ColorBoolTrue
		text = "true"
	}
	return Wrap(text, color)
}
