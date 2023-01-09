package glog

import "fmt"

// Percentage assumes `n` to be normalized (0..1), multiplies it with 100,
// formats it with the given `precision` and colors the result.
//
// Related config setting(s):
//
//  - `n`  < 0: `LoggerConfig.ColorFloatNegative`
//  - `n` == 0: `LoggerConfig.ColorFloatZero`
//  - `n`  > 0: `LoggerConfig.ColorFloatPositive`
func Percentage[F Floats](n F, precision int) string {
	color := LoggerConfig.ColorPercentagePositive
	if n < 0 {
		color = LoggerConfig.ColorPercentageNegative
	} else if n == 0 {
		color = LoggerConfig.ColorPercentageZero
	}
	return Wrap(fmt.Sprintf(fmt.Sprintf("%%.%df%%%%", precision), n*100.0), color)
}

// Float formats `n` with the given `precision` and colors the result.
//
// Related config setting(s):
//
//  - `n`  < 0: `LoggerConfig.ColorFloatNegative`
//  - `n` == 0: `LoggerConfig.ColorFloatZero`
//  - `n`  > 0: `LoggerConfig.ColorFloatPositive`
func Float[F Floats](n F, precision int) string {
	color := LoggerConfig.ColorFloatPositive
	if n < 0 {
		color = LoggerConfig.ColorFloatNegative
	} else if n == 0 {
		color = LoggerConfig.ColorFloatZero
	}
	return Wrap(fmt.Sprintf(fmt.Sprintf("%%.%df", precision), n), color)
}
