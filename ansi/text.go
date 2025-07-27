package ansi

func Ln() *ANSI {
	return New("\n")
}

// Text formatting functions

func Reset() *ANSI {
	return New("\033[0m")
}

func Bold() *ANSI {
	return New("\033[1m")
}

func Dim() *ANSI {
	return New("\033[2m")
}

func Italic() *ANSI {
	return New("\033[3m")
}

func Underline() *ANSI {
	return New("\033[4m")
}

func DoubleUnderline() *ANSI {
	return New("\033[21m")
}

func Blink() *ANSI {
	return New("\033[5m")
}

func RapidBlink() *ANSI {
	return New("\033[6m")
}

func Reverse() *ANSI {
	return New("\033[7m")
}

func Conceal() *ANSI {
	return New("\033[8m")
}

func StrikeThrough() *ANSI {
	return New("\033[9m")
}
