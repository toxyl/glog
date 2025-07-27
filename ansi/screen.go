package ansi

import "fmt"

// Screen control functions

func ClearScreen() *ANSI {
	return New("\033[2J")
}

func ClearScreenFromCursor() *ANSI {
	return New("\033[J")
}

func ClearScreenToCursor() *ANSI {
	return New("\033[1J")
}

func ClearLine() *ANSI {
	return New("\033[2K")
}

func ClearToEOL() *ANSI {
	return New("\033[K")
}

func ClearLineToCursor() *ANSI {
	return New("\033[1K")
}

func ScrollUp(n int) *ANSI {
	return New(fmt.Sprintf("\033[%dS", n))
}

func ScrollDown(n int) *ANSI {
	return New(fmt.Sprintf("\033[%dT", n))
}

func EnableLineWrap() *ANSI {
	return New("\033[?7h")
}

func DisableLineWrap() *ANSI {
	return New("\033[?7l")
}
