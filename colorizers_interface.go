package glog

import (
	"fmt"
	"strings"
	"time"
)

// Auto will automatically choose a highlighter based on the type of the given `interface{}`s.
func Auto(iface ...interface{}) string {
	res := []string{}
	for _, i := range iface {
		text := ""
		switch t := i.(type) {
		case nil:
			text = Wrap("nil", LoggerConfig.ColorNil)
		case []bool:
			text = Bool(t...)
		case []int:
			text = Int(t...)
		case []int8:
			text = Int8(t...)
		case []int16:
			text = Int16(t...)
		case []int32:
			text = Int32(t...)
		case []int64:
			text = Int64(t...)
		case []uint:
			text = Uint(t...)
		case []uint8:
			text = Uint8(t...)
		case []uint16:
			text = Uint16(t...)
		case []uint32:
			text = Uint32(t...)
		case []uint64:
			text = Uint64(t...)
		case []float32:
			res := []string{}
			for _, v := range t {
				if v >= -1.0 && v <= 1.0 {
					res = append(res, Percentage(float64(v), LoggerConfig.AutoFloatPrecision))
				} else {
					res = append(res, Float32(v, LoggerConfig.AutoFloatPrecision))
				}
			}
			text = strings.Join(res, ", ")
		case []float64:
			res := []string{}
			for _, v := range t {
				if v >= -1.0 && v <= 1.0 {
					res = append(res, Percentage(float64(v), LoggerConfig.AutoFloatPrecision))
				} else {
					res = append(res, Float64(v, LoggerConfig.AutoFloatPrecision))
				}
			}
			text = strings.Join(res, ", ")
		case []time.Time:
			res := []string{}
			for _, v := range t {
				res = append(res, DateTime(v))
			}
			text = strings.Join(res, ", ")
		case []time.Duration:
			res := []string{}
			for _, v := range t {
				res = append(res, Duration(uint(v.Seconds())))
			}
			text = strings.Join(res, ", ")
		case []string:
			res := []string{}
			for _, v := range t {
				res = append(res, Highlight(v))
			}
			text = strings.Join(res, ", ")
		case []interface{}:
			text = Auto(t...)
		case bool:
			text = Bool(t)
		case int:
			text = Int(t)
		case int8:
			text = Int8(t)
		case int16:
			text = Int16(t)
		case int32:
			text = Int32(t)
		case int64:
			text = Int64(t)
		case uint:
			text = Uint(t)
		case uint8:
			text = Uint8(t)
		case uint16:
			text = Uint16(t)
		case uint32:
			text = Uint32(t)
		case uint64:
			text = Uint64(t)
		case float32:
			if t >= -1.0 && t <= 1.0 {
				text = Percentage(float64(t), LoggerConfig.AutoFloatPrecision)
			} else {
				text = Float32(t, LoggerConfig.AutoFloatPrecision)
			}
		case float64:
			if t >= -1.0 && t <= 1.0 {
				text = Percentage(t, LoggerConfig.AutoFloatPrecision)
			} else {
				text = Float64(t, LoggerConfig.AutoFloatPrecision)
			}
		case time.Time:
			text = DateTime(t)
		case time.Duration:
			text = Duration(uint(t.Seconds()))
		default:
			text = Highlight(fmt.Sprint(t))
		}
		res = append(res, text)
	}
	return strings.Join(res, ", ")
}
