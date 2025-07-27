package ansi

import (
	"fmt"

	"github.com/toxyl/glog/colormap"
)

// Foreground colors
func Color(color int) *ANSI {
	return New(fmt.Sprintf("\033[38;5;%dm", colormap.MapColor(color)))
}

func Wrap(text string, color int) *ANSI {
	return New(Color(color).sequence + text + Reset().sequence)
}

// Background colors
func BackgroundColor(color int) *ANSI {
	return New(fmt.Sprintf("\033[48;5;%dm", colormap.MapColor(color)))
}

func WrapBackground(text string, color int) *ANSI {
	return New(BackgroundColor(color).sequence + text + Reset().sequence)
}

// Convenience functions for common colors
func Red() *ANSI {
	return Color(1)
}

func Green() *ANSI {
	return Color(2)
}

func Blue() *ANSI {
	return Color(4)
}

func Yellow() *ANSI {
	return Color(3)
}

func Magenta() *ANSI {
	return Color(5)
}

func Cyan() *ANSI {
	return Color(6)
}

func White() *ANSI {
	return Color(7)
}

func Black() *ANSI {
	return Color(0)
}
