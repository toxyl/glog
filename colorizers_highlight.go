package glog

import (
	"strings"
)

// Highlight colorizes each given string (identical strings always get the same color).
func Highlight(message ...string) string {
	res := []string{}
	for _, msg := range message {
		res = append(res, Wrap(msg, stringColorCache.Get(msg)))
	}
	return strings.Join(res, ", ")
}

// HighlightInfo colorizes the `message`.
//
// Related config setting(s):
//
//  - `LoggerConfig.ColorIndicatorInfo`
func HighlightInfo(message string) string {
	return Wrap(message, LoggerConfig.ColorIndicatorInfo)
}

// HighlightOK colorizes the `message`.
//
// Related config setting(s):
//
//  - `LoggerConfig.ColorIndicatorOK`
func HighlightOK(message string) string {
	return Wrap(message, LoggerConfig.ColorIndicatorOK)
}

// HighlightSuccess colorizes the `message`.
//
// Related config setting(s):
//
//  - `LoggerConfig.ColorIndicatorSuccess`
func HighlightSuccess(message string) string {
	return Wrap(message, LoggerConfig.ColorIndicatorSuccess)
}

// HighlightNotOK colorizes the `message`.
//
// Related config setting(s):
//
//  - `LoggerConfig.ColorIndicatorNotOK`
func HighlightNotOK(message string) string {
	return Wrap(message, LoggerConfig.ColorIndicatorNotOK)
}

// HighlightError colorizes the `message`.
//
// Related config setting(s):
//
//  - `LoggerConfig.ColorIndicatorError`
func HighlightError(message string) string {
	return Wrap(message, LoggerConfig.ColorIndicatorError)
}

// HighlightWarning colorizes the `message`.
//
// Related config setting(s):
//
//  - `LoggerConfig.ColorIndicatorWarning`
func HighlightWarning(message string) string {
	return Wrap(message, LoggerConfig.ColorIndicatorWarning)
}

// HighlightDebug colorizes the `message`.
//
// Related config setting(s):
//
//  - `LoggerConfig.ColorIndicatorDebug`
func HighlightDebug(message string) string {
	return Wrap(message, LoggerConfig.ColorIndicatorDebug)
}

// HighlightQuestion colorizes the `message`.
//
// Related config setting(s):
//
//  - `LoggerConfig.ColorIndicatorQuestion`
func HighlightQuestion(message string) string {
	return Wrap(message, LoggerConfig.ColorIndicatorQuestion)
}

// HighlightTrace colorizes the `message`.
//
// Related config setting(s):
//
//  - `LoggerConfig.ColorIndicatorTrace`
func HighlightTrace(message string) string {
	return Wrap(message, LoggerConfig.ColorIndicatorTrace)
}
