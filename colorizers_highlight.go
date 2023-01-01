package glog

import (
	"strings"
)

// Highlight colorizes each given string (identical strings always get the same color).
func Highlight(message ...string) string {
	res := []string{}
	for _, msg := range message {
		res = append(res, Wrap(msg, getStringColor(msg)))
	}
	return strings.Join(res, ", ")
}

// HighlightInfo colorizes the `message` using `LoggerConfig.ColorIndicatorInfo`.
func HighlightInfo(message string) string {
	return Wrap(message, LoggerConfig.ColorIndicatorInfo)
}

// HighlightOK colorizes the `message` using `LoggerConfig.ColorIndicatorOK`.
func HighlightOK(message string) string {
	return Wrap(message, LoggerConfig.ColorIndicatorOK)
}

// HighlightSuccess colorizes the `message` using `LoggerConfig.ColorIndicatorSuccess`.
func HighlightSuccess(message string) string {
	return Wrap(message, LoggerConfig.ColorIndicatorSuccess)
}

// HighlightNotOK colorizes the `message` using `LoggerConfig.ColorIndicatorNotOK`.
func HighlightNotOK(message string) string {
	return Wrap(message, LoggerConfig.ColorIndicatorNotOK)
}

// HighlightError colorizes the `message` using `LoggerConfig.ColorIndicatorError`.
func HighlightError(message string) string {
	return Wrap(message, LoggerConfig.ColorIndicatorError)
}

// HighlightWarning colorizes the `message` using `LoggerConfig.ColorIndicatorWarning`.
func HighlightWarning(message string) string {
	return Wrap(message, LoggerConfig.ColorIndicatorWarning)
}

// HighlightDebug colorizes the `message` using `LoggerConfig.ColorIndicatorDebug`.
func HighlightDebug(message string) string {
	return Wrap(message, LoggerConfig.ColorIndicatorDebug)
}

// HighlightQuestion colorizes the `message` using `LoggerConfig.ColorIndicatorQuestion`.
func HighlightQuestion(message string) string {
	return Wrap(message, LoggerConfig.ColorIndicatorQuestion)
}

// HighlightTrace colorizes the `message` using `LoggerConfig.ColorIndicatorTrace`.
func HighlightTrace(message string) string {
	return Wrap(message, LoggerConfig.ColorIndicatorTrace)
}
