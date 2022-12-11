package glog

import (
	"math"
	"strings"

	"github.com/toxyl/gutils"
)

var stringColorCache map[string]int = map[string]int{}

func getStringColor(str string) int {
	if v, ok := stringColorCache[str]; ok {
		return v
	}
	pt := 0.0
	bm := []rune(gutils.RemoveNonPrintable(strings.ToUpper(str)))
	l := len(bm)
	for i := 0; i < l; i++ {
		pt += (math.Max(0.0, math.Min(96.0, float64(bm[i])-33.0)) / 96.0) / float64(l)
	}
	stringColorCache[str] = int(88.0 + 143.0*pt) // 88 - 231 (143 total)
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
