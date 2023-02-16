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

func Reset() string {
	return "\033[0m"
}

func Color(color int) string {
	return fmt.Sprintf("\033[38;5;%dm", MapColor(color))
}

func Wrap(str string, color int) string {
	return Color(color) + str + Reset()
}

func Bold() string {
	return "\033[1m"
}

func Italic() string {
	return "\033[3m"
}

func Underline() string {
	return "\033[4m"
}

func StrikeThrough() string {
	return "\033[9m"
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
