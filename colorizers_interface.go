package glog

import (
	"fmt"
	"strings"
)

// Auto will automatically choose a highlighter based on the type of the given `interface{}`s.
func Auto(values ...interface{}) string {
	res := []string{}
	for _, i := range values {
		if ok, v := castBoolSlice(i); ok {
			res = append(res, Bool(v...))
			continue
		}

		if ok, v := castIntSlice(i); ok {
			res = append(res, Int(v...))
			continue
		}

		if ok, v := castUintSlice(i); ok {
			res = append(res, Uint(v...))
			continue
		}

		if ok, v := castFloatSlice(i); ok {
			for _, f := range v {
				if f >= -1.0 && f <= 1.0 {
					res = append(res, Percentage(f, LoggerConfig.AutoFloatPrecision))
				} else {
					res = append(res, Float(f, LoggerConfig.AutoFloatPrecision))
				}
			}
			continue
		}

		if ok, v := castTimeSlice(i); ok {
			for _, t := range v {
				res = append(res, DateTime(t))
			}
			continue
		}

		if ok, v := castDurationSlice(i); ok {
			for _, d := range v {
				res = append(res, Duration(d.Seconds()))
			}
			continue
		}

		if ok, v := castStringSlice(i); ok {
			res = append(res, Highlight(v...))
			continue
		}

		if ok, v := castBool(i); ok {
			res = append(res, Bool(v))
			continue
		}

		if ok, v := castInt(i); ok {
			res = append(res, Int(v))
			continue
		}

		if ok, v := castUint(i); ok {
			res = append(res, Uint(v))
			continue
		}

		if ok, normalized, v := castFloat(i); ok {
			if normalized {
				res = append(res, Percentage(v, LoggerConfig.AutoFloatPrecision))
			} else {
				res = append(res, Float(v, LoggerConfig.AutoFloatPrecision))
			}
			continue
		}

		if ok, v := castTime(i); ok {
			res = append(res, DateTime(v))
			continue
		}

		if ok, v := castDuration(i); ok {
			res = append(res, Duration(v.Seconds()))
			continue
		}

		if ok, v := castString(i); ok {
			res = append(res, Highlight(v))
			continue
		}

		if i == nil {
			res = append(res, Wrap("nil", LoggerConfig.ColorNil))
			continue
		}

		if ok, v := castInterfaceSlice(i); ok {
			res = append(res, Auto(v...))
			continue
		}

		// if we get we have an unknown type, let's output it as string
		res = append(res, Highlight(fmt.Sprint(i)))
	}
	return strings.Join(res, ", ")
}
