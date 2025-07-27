package colorizers

import (
	"fmt"

	"github.com/toxyl/glog/ansi"
	"github.com/toxyl/glog/config"
	"github.com/toxyl/glog/types"
)

// Percentage assumes `n` to be normalized (0..1), multiplies it with 100,
// formats it with the given `precision` and colors the result.
//
// Related config setting(s):
//
//   - `n`  < 0: `LoggerConfig.ColorFloatNegative`
//   - `n` == 0: `LoggerConfig.ColorFloatZero`
//   - `n`  > 0: `LoggerConfig.ColorFloatPositive`
func Percentage[F types.Floats](n F, precision int) string {
	color := config.LoggerConfig.ColorPercentagePositive
	if n < 0 {
		color = config.LoggerConfig.ColorPercentageNegative
	} else if n == 0 {
		color = config.LoggerConfig.ColorPercentageZero
	}
	return ansi.Wrap(fmt.Sprintf(fmt.Sprintf("%%.%df%%%%", precision), n*100.0), color).String()
}

// Float formats `n` with the given `precision` and colors the result.
//
// Related config setting(s):
//
//   - `n`  < 0: `LoggerConfig.ColorFloatNegative`
//   - `n` == 0: `LoggerConfig.ColorFloatZero`
//   - `n`  > 0: `LoggerConfig.ColorFloatPositive`
func Float[F types.Floats](n F, precision int) string {
	color := config.LoggerConfig.ColorFloatPositive
	if n < 0 {
		color = config.LoggerConfig.ColorFloatNegative
	} else if n == 0 {
		color = config.LoggerConfig.ColorFloatZero
	}
	return ansi.Wrap(fmt.Sprintf(fmt.Sprintf("%%.%df", precision), n), color).String()
}
