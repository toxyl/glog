package glog

import "time"

// castTime tries to cast the given value as a time.Time value.
func castTime(value interface{}) (ok bool, v time.Time) {
	ok = false
	v = time.Now()
	switch t := value.(type) {
	case time.Time:
		v = t
		ok = true
	}
	return
}

// castTimeSlice tries to cast the given value as a time.Time slice.
func castTimeSlice(value interface{}) (ok bool, v []time.Time) {
	ok = false
	v = []time.Time{}
	switch t := value.(type) {
	case []time.Time:
		v = t
		ok = true
	}
	return
}

// castDuration tries to cast the given value as a time.Duration value.
func castDuration(value interface{}) (ok bool, v time.Duration) {
	ok = false
	v = time.Duration(0)
	switch t := value.(type) {
	case time.Duration:
		v = t
		ok = true
	}
	return
}

// castDurationSlice tries to cast the given value as a time.Duration slice.
func castDurationSlice(value interface{}) (ok bool, v []time.Duration) {
	ok = false
	v = []time.Duration{}
	switch t := value.(type) {
	case []time.Duration:
		v = t
		ok = true
	}
	return
}

// castString tries to cast the given value as a string.
func castString(value interface{}) (ok bool, v string) {
	ok = false
	v = ""
	switch t := value.(type) {
	case string:
		v = t
		ok = true
	}
	return
}

// castStringSlice tries to cast the given value as a string slice.
func castStringSlice(value interface{}) (ok bool, v []string) {
	ok = false
	v = []string{}
	switch t := value.(type) {
	case []string:
		v = t
		ok = true
	}
	return
}

// castInterfaceSlice tries to cast the given value as a interface{} slice.
func castInterfaceSlice(value interface{}) (ok bool, v []interface{}) {
	ok = false
	v = []interface{}{}
	switch t := value.(type) {
	case []interface{}:
		v = t
		ok = true
	}
	return
}

// castBool tries to cast the given value as a boolean.
func castBool(value interface{}) (ok bool, v bool) {
	ok = false
	v = false
	switch t := value.(type) {
	case bool:
		v = t
		ok = true
	}
	return
}

// castBoolSlice tries to cast the given value as a bool slice.
func castBoolSlice(value interface{}) (ok bool, v []bool) {
	ok = false
	v = []bool{}
	switch t := value.(type) {
	case []bool:
		v = t
		ok = true
	}
	return
}

// castInt tries to cast the given value as an integer.
func castInt(value interface{}) (ok bool, v int64) {
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

// castToInt64Slice consolidates the given values in an int64 slice.
func castToInt64Slice[I IntOrInterface](v ...I) []int64 {
	res := []int64{}
	for _, v := range v {
		_, vc := castInt(v)
		res = append(res, vc)
	}
	return res
}

// castIntSlice tries to cast the given value to a int64 slice.
func castIntSlice(value interface{}) (ok bool, v []int64) {
	ok = false
	v = []int64{}

	switch t := value.(type) {
	case []int:
		v = castToInt64Slice(t...)
		ok = true
	case []int8:
		v = castToInt64Slice(t...)
		ok = true
	case []int16:
		v = castToInt64Slice(t...)
		ok = true
	case []int32:
		v = castToInt64Slice(t...)
		ok = true
	case []int64:
		v = t
		ok = true
	}
	return
}

// castUint tries to cast the given value as an unsigned integer.
func castUint(value interface{}) (ok bool, v uint64) {
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

// castToUint64Slice consolidates the given values in an uint64 slice.
func castToUint64Slice[U UintOrInterface](v ...U) []uint64 {
	res := []uint64{}
	for _, v := range v {
		_, vc := castUint(v)
		res = append(res, vc)
	}
	return res
}

// castUintSlice tries to cast the given value to a uint64 slice.
func castUintSlice(value interface{}) (ok bool, v []uint64) {
	ok = false
	v = []uint64{}

	switch t := value.(type) {
	case []uint:
		v = castToUint64Slice(t...)
		ok = true
	case []uint8:
		v = castToUint64Slice(t...)
		ok = true
	case []uint16:
		v = castToUint64Slice(t...)
		ok = true
	case []uint32:
		v = castToUint64Slice(t...)
		ok = true
	case []uint64:
		v = t
		ok = true
	}
	return
}

// castFloat tries to cast the given value as a float.
func castFloat(value interface{}) (ok bool, normalized bool, v float64) {
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

// castToFloat64Slice consolidates the given values in a float64 slice.
func castToFloat64Slice[F FloatOrInterface](v ...F) []float64 {
	res := []float64{}
	for _, v := range v {
		_, _, vc := castFloat(v)
		res = append(res, vc)
	}
	return res
}

// castFloatSlice tries to cast the given value to a float64 slice.
func castFloatSlice(value interface{}) (ok bool, v []float64) {
	ok = false
	v = []float64{}

	switch t := value.(type) {
	case []float32:
		v = castToFloat64Slice(t...)
		ok = true
	case []float64:
		v = t
		ok = true
	}
	return
}
