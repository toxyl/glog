package glog

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

var (
	reANSIClose = regexp.MustCompile(`\033\[0m`)
	reANSIOpen  = regexp.MustCompile(`\033\[38;5;\d+m`)
)

func Wrap(str string, color int) string {
	return fmt.Sprintf("\033[38;5;%dm%s\033[0m", MapColor(color), str)
}

func StripANSI(str string) string {
	str = reANSIClose.ReplaceAllString(str, "")
	str = reANSIOpen.ReplaceAllString(str, "")
	return str
}

func RemoveNonPrintable(str string) string {
	return strings.TrimFunc(str, func(r rune) bool {
		return !unicode.IsGraphic(r)
	})
}
