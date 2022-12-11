package glog

import (
	"fmt"
	"strings"
)

// Uint colors the result cyan (`n` > 0), blue (`n` == 0) or red (`n` < 0).
func Uint(n ...uint) string {
	res := []string{}
	for _, num := range n {
		color := LoggerConfig.ColorUintPositive
		if num == 0 {
			color = LoggerConfig.ColorUintZero
		}
		res = append(res, Wrap(fmt.Sprintf("%d", num), color))
	}
	return strings.Join(res, ", ")
}

// Uint8 colors the result cyan (`n` > 0), blue (`n` == 0) or red (`n` < 0).
func Uint8(n ...uint8) string {
	uints := []uint{}
	for _, num := range n {
		uints = append(uints, uint(num))
	}
	return Uint(uints...)
}

// Uint16 colors the result cyan (`n` > 0), blue (`n` == 0) or red (`n` < 0).
func Uint16(n ...uint16) string {
	uints := []uint{}
	for _, num := range n {
		uints = append(uints, uint(num))
	}
	return Uint(uints...)
}

// Uint32 colors the result cyan (`n` > 0), blue (`n` == 0) or red (`n` < 0).
func Uint32(n ...uint32) string {
	uints := []uint{}
	for _, num := range n {
		uints = append(uints, uint(num))
	}
	return Uint(uints...)
}

// Uint64 colors the result cyan (`n` > 0), blue (`n` == 0) or red (`n` < 0).
func Uint64(n ...uint64) string {
	uints := []uint{}
	for _, num := range n {
		uints = append(uints, uint(num))
	}
	return Uint(uints...)
}
