package glog

import "fmt"

// Uint colors the result cyan (`n` > 0), blue (`n` == 0) or red (`n` < 0).
func Uint(n uint) string {
	color := LoggerConfig.ColorUintPositive
	if n == 0 {
		color = LoggerConfig.ColorUintZero
	}
	return Wrap(fmt.Sprintf("%d", n), color)
}

// Uint8 colors the result cyan (`n` > 0), blue (`n` == 0) or red (`n` < 0).
func Uint8(n uint8) string {
	return Uint(uint(n))
}

// Uint16 colors the result cyan (`n` > 0), blue (`n` == 0) or red (`n` < 0).
func Uint16(n uint16) string {
	return Uint(uint(n))
}

// Uint32 colors the result cyan (`n` > 0), blue (`n` == 0) or red (`n` < 0).
func Uint32(n uint32) string {
	return Uint(uint(n))
}

// Uint64 colors the result cyan (`n` > 0), blue (`n` == 0) or red (`n` < 0).
func Uint64(n uint64) string {
	return Uint(uint(n))
}
