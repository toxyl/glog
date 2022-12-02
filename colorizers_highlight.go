package glog

func Highlight(message string) string {
	return LoggerConfig.Indicators[' '].Wrap(message)
}

func HighlightInfo(message string) string {
	return LoggerConfig.Indicators['i'].Wrap(message)
}

func HighlightOK(message string) string {
	return LoggerConfig.Indicators['+'].Wrap(message)
}

func HighlightSuccess(message string) string {
	return LoggerConfig.Indicators['✓'].Wrap(message)
}

func HighlightNotOK(message string) string {
	return LoggerConfig.Indicators['-'].Wrap(message)
}

func HighlightError(message string) string {
	return LoggerConfig.Indicators['x'].Wrap(message)
}

func HighlightWarning(message string) string {
	return LoggerConfig.Indicators['!'].Wrap(message)
}

func HighlightDebug(message string) string {
	return LoggerConfig.Indicators['d'].Wrap(message)
}
