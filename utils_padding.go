package glog

import (
	"strings"
)

func getPadLength(str string, maxLength int) int {
	return Max(0, maxLength-plaintextStringLengthForPadding(str))
}

func PadLeft[I IntOrUint](str string, maxLength I, char rune) string {
	return strings.Repeat(string(char), getPadLength(str, int(maxLength))) + str
}

func PadRight[I IntOrUint](str string, maxLength I, char rune) string {
	return str + strings.Repeat(string(char), getPadLength(str, int(maxLength)))
}

func PadCenter[I IntOrUint](str string, maxLength I, char rune) string {
	padLen := getPadLength(str, int(maxLength))
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
