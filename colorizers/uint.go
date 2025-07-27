package colorizers

import (
	"fmt"
	"strings"

	"github.com/toxyl/glog/ansi"
	"github.com/toxyl/glog/config"
	"github.com/toxyl/glog/types"
)

// Uint colors the given uints.
//
// Related config setting(s):
//
//   - `n`  > 0: `LoggerConfig.ColorUintPositive`
//   - `n` == 0: `LoggerConfig.ColorUintZero`
func Uint[U types.Uints](n ...U) string {
	res := []string{}
	for _, num := range n {
		color := config.LoggerConfig.ColorUintPositive
		if num == 0 {
			color = config.LoggerConfig.ColorUintZero
		}
		res = append(res, ansi.Wrap(fmt.Sprintf("%d", num), color).String())
	}
	return strings.Join(res, ", ")
}
