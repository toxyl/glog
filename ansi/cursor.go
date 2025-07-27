package ansi

import "fmt"

// Cursor movement functions

func CursorUp(n int) *ANSI {
	return New(fmt.Sprintf("\033[%dA", n))
}

func CursorDown(n int) *ANSI {
	return New(fmt.Sprintf("\033[%dB", n))
}

func CursorForward(n int) *ANSI {
	return New(fmt.Sprintf("\033[%dC", n))
}

func CursorBackward(n int) *ANSI {
	return New(fmt.Sprintf("\033[%dD", n))
}

func CursorPosition(row, col int) *ANSI {
	return New(fmt.Sprintf("\033[%d;%dH", row, col))
}

func CursorToLineStart() *ANSI {
	return New("\r")
}

func CursorToLineStartAndClear() *ANSI {
	return New("\r\033[K")
}

func StoreCursor() *ANSI {
	return New("\033[s")
}

func RestoreCursor() *ANSI {
	return New("\033[u")
}

func HideCursor() *ANSI {
	return New("\033[?25l")
}

func ShowCursor() *ANSI {
	return New("\033[?25h")
}
