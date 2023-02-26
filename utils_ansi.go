package glog

import "fmt"

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

func StoreCursor() string {
	return "\033[s"
}

func RestoreCursor() string {
	return "\033[u"
}
func ClearToEOL() string {
	return "\033[K"
}
