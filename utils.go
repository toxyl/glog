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
	return fmt.Sprintf("\033[38;5;%dm%s\033[0m", MapColor(color), str)
}

func StripANSI(str string) string {
	str = strings.ReplaceAll(str, "\033[0m", "")
	str = reNonANSI.ReplaceAllString(str, "")
	return str
}

func Max[N Number](a, b N) N {
	return N(math.Max(float64(a), float64(b)))
}

func Min[N Number](a, b N) N {
	return N(math.Min(float64(a), float64(b)))
}
