package glog

import (
	"strings"

	"github.com/toxyl/gutils"
)

var stringColorCache map[string]int = map[string]int{}

func getStringColor(str string) int {
	if v, ok := stringColorCache[str]; ok {
		return v
	}
	pt := 0.0
	bm := []rune(gutils.RemoveNonPrintable(str))
	l := len(bm)
	for i := 0; i < l; i++ {
		pt += (Max(0.0, Min(94.0, float64(bm[i])-32.0)) / 94.0) / float64(l)
	}
	stringColorCache[str] = int(16.0 + 215.0*pt) // 16 - 231 (215 total)
	return stringColorCache[str]
}

func Highlight(message ...string) string {
	res := []string{}
	for _, msg := range message {
		res = append(res, Wrap(msg, getStringColor(msg)))
	}
	return strings.Join(res, ", ")
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

func HighlightQuestion(message string) string {
	return LoggerConfig.Indicators['?'].Wrap(message)
}

func HighlightTrace(message string) string {
	return LoggerConfig.Indicators['t'].Wrap(message)
}
