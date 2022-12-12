package glog

import (
	"fmt"
	"math"
	"regexp"
	"strings"
)

var (
	reNonANSI = regexp.MustCompile(`\033\[38;5;\d+m`)
)

func Wrap(str string, color int) string {
	return fmt.Sprintf("\033[38;5;%dm%s\033[0m", color, str)
}

func StripANSI(str string) string {
	str = strings.ReplaceAll(str, "\033[0m", "")
	str = reNonANSI.ReplaceAllString(str, "")
	return str
}

func getPadLength(str string, maxLength int, char rune) int {
	l1 := len(str)
	l2 := len(StripANSI(str))
	if l1 == l2 {
		return int(math.Max(0.0, float64(maxLength-l1))) // plain text
	}
	return int(math.Max(0.0, float64(maxLength-l2))) // string with ANSI escapes

}

func PadLeft(str string, maxLength int, char rune) string {
	return strings.Repeat(string(char), getPadLength(str, maxLength, char)) + str
}

func PadRight(str string, maxLength int, char rune) string {
	return str + strings.Repeat(string(char), getPadLength(str, maxLength, char))
}

func PadCenter(str string, maxLength int, char rune) string {
	padLen := getPadLength(str, maxLength, char)
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
