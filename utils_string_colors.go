package glog

import "github.com/toxyl/gutils"

// map to cache results to avoid repeatedly calculating the color for the same string
var stringColorCache map[string]int = map[string]int{}

// getStringColor takes an input string and maps it to a color within the ANSI range 16-231, i.e. it can generate 215 different colors.
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
