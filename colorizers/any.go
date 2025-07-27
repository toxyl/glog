package colorizers

import (
	"fmt"
	"strings"

	"github.com/toxyl/glog/ansi"
	"github.com/toxyl/glog/cast"
	"github.com/toxyl/glog/config"
	"github.com/toxyl/glog/utils"
)

// Auto will automatically choose a highlighter based on the type of the given `any`s.
//
// Related config setting(s):
//
//   - `LoggerConfig.AutoFloatPrecision`
//   - `LoggerConfig.ColorNil`
func Auto(values ...any) string {
	res := []string{}
	for _, i := range values {
		if ok, v := cast.BoolSlice(i); ok {
			res = append(res, Bool(v...))
			continue
		}

		if ok, v := cast.IntSlice(i); ok {
			res = append(res, Int(v...))
			continue
		}

		if ok, v := cast.UintSlice(i); ok {
			res = append(res, Uint(v...))
			continue
		}

		if ok, v := cast.FloatSlice(i); ok {
			for _, f := range v {
				if f >= -1.0 && f <= 1.0 {
					res = append(res, Percentage(f, config.LoggerConfig.AutoFloatPrecision))
				} else {
					res = append(res, Float(f, config.LoggerConfig.AutoFloatPrecision))
				}
			}
			continue
		}

		if ok, v := cast.TimeSlice(i); ok {
			for _, t := range v {
				res = append(res, DateTime(t))
			}
			continue
		}

		if ok, v := cast.DurationSlice(i); ok {
			for _, d := range v {
				res = append(res, Duration(d.Seconds()))
			}
			continue
		}

		if ok, v := cast.StringSlice(i); ok {
			res = append(res, Highlight(v...))
			continue
		}

		if ok, v := cast.Bool(i); ok {
			res = append(res, Bool(v))
			continue
		}

		if ok, v := cast.Int(i); ok {
			res = append(res, Int(v))
			continue
		}

		if ok, v := cast.Uint(i); ok {
			res = append(res, Uint(v))
			continue
		}

		if ok, normalized, v := cast.Float(i); ok {
			if normalized {
				res = append(res, Percentage(v, config.LoggerConfig.AutoFloatPrecision))
			} else {
				res = append(res, Float(v, config.LoggerConfig.AutoFloatPrecision))
			}
			continue
		}

		if ok, v := cast.Time(i); ok {
			res = append(res, DateTime(v))
			continue
		}

		if ok, v := cast.Duration(i); ok {
			res = append(res, Duration(v.Seconds()))
			continue
		}

		if ok, v := cast.String(i); ok {
			// maybe it's a path
			switch utils.IdentifyPath(v) {
			case utils.INVALID_PATH:
				v = Highlight(v)
			case utils.FILE_PATH:
				v = File(v)
			case utils.URL_PATH:
				v = URL(v)
			}
			res = append(res, v)
			continue
		}

		if i == nil {
			res = append(res, ansi.Wrap("nil", config.LoggerConfig.ColorNil).String())
			continue
		}

		if ok, v := cast.InterfaceSlice(i); ok {
			res = append(res, Auto(v...))
			continue
		}

		// if we get we have an unknown type, let's output it as string
		res = append(res, Highlight(fmt.Sprint(i)))
	}
	return strings.Join(res, ", ")
}
