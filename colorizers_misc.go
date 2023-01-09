package glog

import (
	"os"
	"strings"
)

// Password colors `password` according to the config. It does not redact it!
//
// Related config setting(s):
//
//  - `LoggerConfig.ColorPassword`
func Password(password string) string {
	return Wrap(password, LoggerConfig.ColorPassword)
}

// Error colors `err.Error()` according to the config, or returns "nil" if no error was present.
//
// Related config setting(s):
//
//  - `LoggerConfig.ColorError`
//  - `LoggerConfig.ColorNil`
func Error(err error) string {
	if err == nil {
		return Wrap("nil", LoggerConfig.ColorNil)
	}
	return Wrap(err.Error(), LoggerConfig.ColorError)
}

// Reason colors `reason` according to the config.
//
// Related config setting(s):
//
//  - `LoggerConfig.ColorReason`
func Reason(reason string) string {
	return Wrap(reason, LoggerConfig.ColorReason)
}

// File colors `file` according to the config using the `os.PathSeparator`.
//
// Related config setting(s):
//
//  - `LoggerConfig.ColorPathSeparator`
//  - `LoggerConfig.ColorPath``
func File(file string) string {
	ops := string(os.PathSeparator)
	res := ""
	for i, pe := range strings.Split(file, ops) {
		if pe == "" {
			continue
		}
		if i > 0 {
			res += Wrap(ops, LoggerConfig.ColorPathSeparator)
		}
		res += Wrap(pe, LoggerConfig.ColorPath)
	}
	if len(file) > 0 && string(file[len(file)-1]) == ops {
		res += Wrap(ops, LoggerConfig.ColorPathSeparator)
	}

	return res
}
