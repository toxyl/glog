package colorizers

import (
	"strings"

	"github.com/toxyl/glog/ansi"
	"github.com/toxyl/glog/config"
)

// Bool colors the given booleans.
//
// Related config setting(s):
//
//   - `n` == false: `LoggerConfig.ColorBoolFalse`
//   - `n` == true:  `LoggerConfig.ColorBoolTrue`
func Bool(b ...bool) string {
	res := []string{}
	for _, bo := range b {
		color := config.LoggerConfig.ColorBoolFalse
		text := "false"
		if bo {
			color = config.LoggerConfig.ColorBoolTrue
			text = "true"
		}
		res = append(res, ansi.Wrap(text, color).String())
	}
	return strings.Join(res, ", ")
}
