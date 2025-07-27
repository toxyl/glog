package colorizers

import (
	"fmt"
	"strings"

	"github.com/toxyl/glog/ansi"
	"github.com/toxyl/glog/config"
	"github.com/toxyl/glog/types"
)

// IntAmount colors `n` as int and appends either the
// given singular (`n` == 1) or plural (`n` > 1 || `n` == 0).
//
// Related config setting(s):
//
//   - `n`  > 0: `LoggerConfig.ColorIntPositive`
//   - `n` == 0: `LoggerConfig.ColorIntZero`
//   - `n`  < 0: `LoggerConfig.ColorIntNegative`
func IntAmount[I types.IntOrUint](n I, singular, plural string) string {
	unit := singular
	if n > 1 {
		unit = plural
	}
	color := config.LoggerConfig.ColorIntPositive
	if n < 0 {
		color = config.LoggerConfig.ColorIntNegative
	} else if n == 0 {
		color = config.LoggerConfig.ColorIntZero
	}
	amount := ansi.Wrap(fmt.Sprintf("%d", n), color).String()
	if n == 0 {
		unit = plural
	}
	return fmt.Sprintf("%s %s", amount, unit)
}

// Int colors the given ints.
//
// Related config setting(s):
//
//   - `n`  > 0: `LoggerConfig.ColorIntPositive`
//   - `n` == 0: `LoggerConfig.ColorIntZero`
//   - `n`  < 0: `LoggerConfig.ColorIntNegative`
func Int[I types.IntOrUint](n ...I) string {
	res := []string{}
	for _, num := range n {
		color := config.LoggerConfig.ColorIntPositive
		if num < 0 {
			color = config.LoggerConfig.ColorIntNegative
		} else if num == 0 {
			color = config.LoggerConfig.ColorIntZero
		}
		res = append(res, ansi.Wrap(fmt.Sprintf("%d", num), color).String())
	}
	return strings.Join(res, ", ")
}
