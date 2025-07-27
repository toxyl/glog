package cast

import (
	"time"

	"github.com/toxyl/glog/types"
)

// Time tries to cast the given value as a time.Time value.
func Time(value any) (ok bool, v time.Time) {
	ok = false
	v = time.Now()
	switch t := value.(type) {
	case time.Time:
		v = t
		ok = true
	}
	return
}

// TimeSlice tries to cast the given value as a time.Time slice.
func TimeSlice(value any) (ok bool, v []time.Time) {
	ok = false
	v = []time.Time{}
	switch t := value.(type) {
	case []time.Time:
		v = t
		ok = true
	}
	return
}

// Duration tries to cast the given value as a time.Duration value.
func Duration(value any) (ok bool, v time.Duration) {
	ok = false
	v = time.Duration(0)
	switch t := value.(type) {
	case time.Duration:
		v = t
		ok = true
	}
	return
}

// DurationSlice tries to cast the given value as a time.Duration slice.
func DurationSlice(value any) (ok bool, v []time.Duration) {
	ok = false
	v = []time.Duration{}
	switch t := value.(type) {
	case []time.Duration:
		v = t
		ok = true
	}
	return
}

// String tries to cast the given value as a string.
func String(value any) (ok bool, v string) {
	ok = false
	v = ""
	switch t := value.(type) {
	case string:
		v = t
		ok = true
	}
	return
}

// StringSlice tries to cast the given value as a string slice.
func StringSlice(value any) (ok bool, v []string) {
	ok = false
	v = []string{}
	switch t := value.(type) {
	case []string:
		v = t
		ok = true
	}
	return
}

// InterfaceSlice tries to cast the given value as a any slice.
func InterfaceSlice(value any) (ok bool, v []any) {
	ok = false
	v = []any{}
	switch t := value.(type) {
	case []any:
		v = t
		ok = true
	}
	return
}

// Bool tries to cast the given value as a boolean.
func Bool(value any) (ok bool, v bool) {
	ok = false
	v = false
	switch t := value.(type) {
	case bool:
		v = t
		ok = true
	}
	return
}

// BoolSlice tries to cast the given value as a bool slice.
func BoolSlice(value any) (ok bool, v []bool) {
	ok = false
	v = []bool{}
	switch t := value.(type) {
	case []bool:
		v = t
		ok = true
	}
	return
}

// Int tries to cast the given value as an integer.
func Int(value any) (ok bool, v int64) {
	ok = false
	v = int64(0)
	switch t := value.(type) {
	case int:
		v = int64(t)
		ok = true
	case int8:
		v = int64(t)
		ok = true
	case int16:
		v = int64(t)
		ok = true
	case int32:
		v = int64(t)
		ok = true
	case int64:
		v = t
		ok = true
	}
	return
}

// ToInt64Slice consolidates the given values in an int64 slice.
func ToInt64Slice[I types.IntOrInterface](v ...I) []int64 {
	res := []int64{}
	for _, v := range v {
		_, vc := Int(v)
		res = append(res, vc)
	}
	return res
}

// IntSlice tries to cast the given value to a int64 slice.
func IntSlice(value any) (ok bool, v []int64) {
	ok = false
	v = []int64{}

	switch t := value.(type) {
	case []int:
		v = ToInt64Slice(t...)
		ok = true
	case []int8:
		v = ToInt64Slice(t...)
		ok = true
	case []int16:
		v = ToInt64Slice(t...)
		ok = true
	case []int32:
		v = ToInt64Slice(t...)
		ok = true
	case []int64:
		v = t
		ok = true
	}
	return
}

// Uint tries to cast the given value as an unsigned integer.
func Uint(value any) (ok bool, v uint64) {
	ok = false
	v = uint64(0)
	switch t := value.(type) {
	case uint:
		v = uint64(t)
		ok = true
	case uint8:
		v = uint64(t)
		ok = true
	case uint16:
		v = uint64(t)
		ok = true
	case uint32:
		v = uint64(t)
		ok = true
	case uint64:
		v = uint64(t)
		ok = true
	}
	return
}

// ToUint64Slice consolidates the given values in an uint64 slice.
func ToUint64Slice[U types.UintOrInterface](v ...U) []uint64 {
	res := []uint64{}
	for _, v := range v {
		_, vc := Uint(v)
		res = append(res, vc)
	}
	return res
}

// UintSlice tries to cast the given value to a uint64 slice.
func UintSlice(value any) (ok bool, v []uint64) {
	ok = false
	v = []uint64{}

	switch t := value.(type) {
	case []uint:
		v = ToUint64Slice(t...)
		ok = true
	case []uint8:
		v = ToUint64Slice(t...)
		ok = true
	case []uint16:
		v = ToUint64Slice(t...)
		ok = true
	case []uint32:
		v = ToUint64Slice(t...)
		ok = true
	case []uint64:
		v = t
		ok = true
	}
	return
}

// Float tries to cast the given value as a float.
func Float(value any) (ok bool, normalized bool, v float64) {
	ok = false
	v = float64(0)
	switch t := value.(type) {
	case float32:
		v = float64(t)
		ok = true
	case float64:
		v = t
		ok = true
	}

	normalized = ok && v >= -1.0 && v <= 1.0
	return
}

// ToFloat64Slice consolidates the given values in a float64 slice.
func ToFloat64Slice[F types.FloatOrInterface](v ...F) []float64 {
	res := []float64{}
	for _, v := range v {
		_, _, vc := Float(v)
		res = append(res, vc)
	}
	return res
}

// FloatSlice tries to cast the given value to a float64 slice.
func FloatSlice(value any) (ok bool, v []float64) {
	ok = false
	v = []float64{}

	switch t := value.(type) {
	case []float32:
		v = ToFloat64Slice(t...)
		ok = true
	case []float64:
		v = t
		ok = true
	}
	return
}
