package colorizers

import (
	"os"
	"strings"

	"github.com/toxyl/glog/ansi"
	"github.com/toxyl/glog/config"
)

// Password colors `password` according to the config. It does not redact it!
//
// Related config setting(s):
//
//   - `LoggerConfig.ColorPassword`
func Password(password string) string {
	return ansi.Wrap(password, config.LoggerConfig.ColorPassword).String()
}

// Error colors `err.Error()` according to the config, or returns "nil" if no error was present.
//
// Related config setting(s):
//
//   - `LoggerConfig.ColorError`
//   - `LoggerConfig.ColorNil`
func Error(err error) string {
	if err == nil {
		return ansi.Wrap("nil", config.LoggerConfig.ColorNil).String()
	}
	return ansi.Wrap(err.Error(), config.LoggerConfig.ColorError).String()
}

// Reason colors `reason` according to the config.
//
// Related config setting(s):
//
//   - `LoggerConfig.ColorReason`
func Reason(reason string) string {
	return ansi.Wrap(reason, config.LoggerConfig.ColorReason).String()
}

// File colors `file` according to the config using the `os.PathSeparator`.
//
// Related config setting(s):
//
//   - `LoggerConfig.ColorPathSeparator`
//   - `LoggerConfig.ColorPath“
func File(file string) string {
	ops := string(os.PathSeparator)
	res := ""
	for i, pe := range strings.Split(file, ops) {
		if pe == "" {
			continue
		}
		if i > 0 {
			res += ansi.Wrap(ops, config.LoggerConfig.ColorPathSeparator).String()
		}
		res += ansi.Wrap(pe, config.LoggerConfig.ColorPath).String()
	}
	if len(file) > 0 && string(file[len(file)-1]) == ops {
		res += ansi.Wrap(ops, config.LoggerConfig.ColorPathSeparator).String()
	}

	return res
}
