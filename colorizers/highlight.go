package colorizers

import (
	"strings"

	"github.com/toxyl/glog/ansi"
	"github.com/toxyl/glog/config"
	"github.com/toxyl/glog/utils"
)

// Highlight colorizes each given string (identical strings always get the same color).
func Highlight(message ...string) string {
	res := []string{}
	for _, msg := range message {
		res = append(res, ansi.Wrap(msg, utils.Scc.Get(msg)).String())
	}
	return strings.Join(res, ", ")
}

// HighlightInfo colorizes the `message`.
//
// Related config setting(s):
//
//   - `LoggerConfig.ColorIndicatorInfo`
func HighlightInfo(message string) string {
	return ansi.Wrap(message, config.LoggerConfig.ColorIndicatorInfo).String()
}

// HighlightOK colorizes the `message`.
//
// Related config setting(s):
//
//   - `LoggerConfig.ColorIndicatorOK`
func HighlightOK(message string) string {
	return ansi.Wrap(message, config.LoggerConfig.ColorIndicatorOK).String()
}

// HighlightSuccess colorizes the `message`.
//
// Related config setting(s):
//
//   - `LoggerConfig.ColorIndicatorSuccess`
func HighlightSuccess(message string) string {
	return ansi.Wrap(message, config.LoggerConfig.ColorIndicatorSuccess).String()
}

// HighlightNotOK colorizes the `message`.
//
// Related config setting(s):
//
//   - `LoggerConfig.ColorIndicatorNotOK`
func HighlightNotOK(message string) string {
	return ansi.Wrap(message, config.LoggerConfig.ColorIndicatorNotOK).String()
}

// HighlightError colorizes the `message`.
//
// Related config setting(s):
//
//   - `LoggerConfig.ColorIndicatorError`
func HighlightError(message string) string {
	return ansi.Wrap(message, config.LoggerConfig.ColorIndicatorError).String()
}

// HighlightWarning colorizes the `message`.
//
// Related config setting(s):
//
//   - `LoggerConfig.ColorIndicatorWarning`
func HighlightWarning(message string) string {
	return ansi.Wrap(message, config.LoggerConfig.ColorIndicatorWarning).String()
}

// HighlightDebug colorizes the `message`.
//
// Related config setting(s):
//
//   - `LoggerConfig.ColorIndicatorDebug`
func HighlightDebug(message string) string {
	return ansi.Wrap(message, config.LoggerConfig.ColorIndicatorDebug).String()
}

// HighlightQuestion colorizes the `message`.
//
// Related config setting(s):
//
//   - `LoggerConfig.ColorIndicatorQuestion`
func HighlightQuestion(message string) string {
	return ansi.Wrap(message, config.LoggerConfig.ColorIndicatorQuestion).String()
}

// HighlightTrace colorizes the `message`.
//
// Related config setting(s):
//
//   - `LoggerConfig.ColorIndicatorTrace`
func HighlightTrace(message string) string {
	return ansi.Wrap(message, config.LoggerConfig.ColorIndicatorTrace).String()
}
