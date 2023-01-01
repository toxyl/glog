package glog

import "fmt"

// Percentage assumes `n` to be normalized (0..1), multiplies it with 100,
// formats it with the given `precision` and colors the result using
// `LoggerConfig.ColorPercentagePositive` (`n` > 0),
// `LoggerConfig.ColorPercentageZero` (`n` == 0) or
// `LoggerConfig.ColorPercentageNegative` (`n` < 0).
func Percentage[F Floats](n F, precision int) string {
	color := LoggerConfig.ColorPercentagePositive
	if n < 0 {
		color = LoggerConfig.ColorPercentageNegative
	} else if n == 0 {
		color = LoggerConfig.ColorPercentageZero
	}
	return Wrap(fmt.Sprintf(fmt.Sprintf("%%.%df%%%%", precision), n*100.0), color)
}

// Float formats `n` with the given `precision` and colors the result using
// `LoggerConfig.ColorFloatPositive` (`n` > 0),
// `LoggerConfig.ColorFloatZero` (`n` == 0) or
// `LoggerConfig.ColorFloatNegative` (`n` < 0).
func Float[F Floats](n F, precision int) string {
	color := LoggerConfig.ColorFloatPositive
	if n < 0 {
		color = LoggerConfig.ColorFloatNegative
	} else if n == 0 {
		color = LoggerConfig.ColorFloatZero
	}
	return Wrap(fmt.Sprintf(fmt.Sprintf("%%.%df", precision), n), color)
}
