package glog

import (
	"os"
	"strings"
)

func Password(password string) string {
	return Wrap(password, LoggerConfig.ColorPassword)
}

func Error(err error) string {
	if err == nil {
		return Wrap("nil", LoggerConfig.ColorNil)
	}
	return Wrap(err.Error(), LoggerConfig.ColorError)
}

func Reason(reason string) string {
	return Wrap(reason, LoggerConfig.ColorReason)
}

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
