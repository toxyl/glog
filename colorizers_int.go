package glog

import (
	"fmt"
)

// IntAmount colors `n` as int and appends either the
// given singular (`n` == 1) or plural (`n` > 1 || `n` == 0).
func IntAmount(n int, singular, plural string) string {
	unit := singular
	if n > 1 {
		unit = plural
	}
	color := LoggerConfig.ColorIntPositive
	if n < 0 {
		color = LoggerConfig.ColorIntNegative
	} else if n == 0 {
		color = LoggerConfig.ColorIntZero
	}
	amount := Wrap(fmt.Sprintf("%d", n), color)
	if n == 0 {
		unit = plural
	}
	return fmt.Sprintf("%s %s", amount, unit)
}

// Int colors the result cyan (`n` > 0), blue (`n` == 0) or red (`n` < 0).
func Int(n int) string {
	color := LoggerConfig.ColorIntPositive
	if n < 0 {
		color = LoggerConfig.ColorIntNegative
	} else if n == 0 {
		color = LoggerConfig.ColorIntZero
	}
	return Wrap(fmt.Sprintf("%d", n), color)
}

// Int8 colors the result cyan (`n` > 0), blue (`n` == 0) or red (`n` < 0).
func Int8(n int8) string {
	return Int(int(n))
}

// Int16 colors the result cyan (`n` > 0), blue (`n` == 0) or red (`n` < 0).
func Int16(n int16) string {
	return Int(int(n))
}

// Int32 colors the result cyan (`n` > 0), blue (`n` == 0) or red (`n` < 0).
func Int32(n int32) string {
	return Int(int(n))
}

// Int64 colors the result cyan (`n` > 0), blue (`n` == 0) or red (`n` < 0).
func Int64(n int64) string {
	return Int(int(n))
}
