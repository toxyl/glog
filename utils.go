package glog

import (
	"fmt"
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
