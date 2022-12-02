package glog

import "fmt"

// Percentage assumes `n` to be normalized (0..1), multiplies it with 100,
// formats it with the given `precision` and colors the result cyan (`n` > 0), blue (`n` == 0) or red (`n` < 0).
func Percentage(n float64, precision int) string {
	color := LoggerConfig.ColorPercentagePositive
	if n < 0 {
		color = LoggerConfig.ColorPercentageNegative
	} else if n == 0 {
		color = LoggerConfig.ColorPercentageZero
	}
	return Wrap(fmt.Sprintf(fmt.Sprintf("%%.%df%%%%", precision), n*100.0), color)
}

// Float64 formats `n` with the given `precision` and colors the result cyan (`n` > 0), blue (`n` == 0) or red (`n` < 0).
func Float64(n float64, precision int) string {
	color := LoggerConfig.ColorFloatPositive
	if n < 0 {
		color = LoggerConfig.ColorFloatNegative
	} else if n == 0 {
		color = LoggerConfig.ColorFloatZero
	}
	return Wrap(fmt.Sprintf(fmt.Sprintf("%%.%df", precision), n), color)
}

// Float32 formats `n` with the given `precision` and colors the result cyan (`n` > 0), blue (`n` == 0) or red (`n` < 0).
func Float32(n float32, precision int) string {
	return Float64(float64(n), precision)
}
