package glog

import (
	"math"
	"strings"
)

func getPadLength(str string, maxLength int, char rune) int {
	l1 := len(str)
	l2 := len(StripANSI(str))
	if l1 == l2 {
		return int(math.Max(0.0, float64(maxLength-l1))) // plain text
	}
	return int(math.Max(0.0, float64(maxLength-l2))) // string with ANSI escapes
}

func PadLeft[I IntOrUint](str string, maxLength I, char rune) string {
	return strings.Repeat(string(char), getPadLength(str, int(maxLength), char)) + str
}

func PadRight[I IntOrUint](str string, maxLength I, char rune) string {
	return str + strings.Repeat(string(char), getPadLength(str, int(maxLength), char))
}

func PadCenter[I IntOrUint](str string, maxLength I, char rune) string {
	padLen := getPadLength(str, int(maxLength), char)
	pl, pr := 0, 0
	if padLen%2 != 0 {
		// uneven length, let's pad one more on the right
		pl = (padLen - 1) / 2
	} else {
		// even length, easy
		pl = padLen / 2
	}
	pr = padLen - pl
	return strings.Repeat(string(char), pl) + str + strings.Repeat(string(char), pr)
}
