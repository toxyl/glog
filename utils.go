package glog

import (
	"fmt"
	"regexp"
)

var (
	reNonANSI = regexp.MustCompile(`\033\[38;5;\d+m(.*?)\033\[0m`)
)

func Wrap(str string, color int) string {
	return fmt.Sprintf("\033[38;5;%dm%s\033[0m", color, str)
}

func StripANSI(str string) string {
	return reNonANSI.ReplaceAllString(str, "$1")
}
