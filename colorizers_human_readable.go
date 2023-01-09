package glog

func wrapHumanReadable(v float64, separator, s string, e int) string {
	if e < 0 {
		e = -e
	}
	e *= 2
	return Float(v, LoggerConfig.AutoFloatPrecision) + separator + Wrap(s, LoggerConfig.ColorUnitHumanReadable+e)
}

// HumanReadableShort colors the same way as Float() does but will make `n` human-readable using short scale suffixes (base1000).
//
// Related config setting(s):
//
//  - `LoggerConfig.AutoFloatPrecision`
//  - `LoggerConfig.ColorUnitHumanReadable` + exponent of the matching scale step
func HumanReadableShort[N Number](value N) string {
	v, e, u := humanReadableShort(value)
	return wrapHumanReadable(v, "", u, e)
}

// HumanReadableLong colors the same way as Float() does but will make `n` human-readable using long scale suffixes  (base1000).
//
// Related config setting(s):
//
//  - `LoggerConfig.AutoFloatPrecision`
//  - `LoggerConfig.ColorUnitHumanReadable` + exponent of the matching scale step
func HumanReadableLong[N Number](value N) string {
	v, e, u := humanReadableLong(value)
	return wrapHumanReadable(v, "", u, e)
}

// HumanReadableIEC colors the same way as Float() does but will make `n` human-readable using IEC-prefixes (base1024).
//
// Related config setting(s):
//
//  - `LoggerConfig.AutoFloatPrecision`
//  - `LoggerConfig.ColorUnitHumanReadable` + exponent of the matching scale step
func HumanReadableIEC[N Number](value N, unit string) string {
	v, e, u := humanReadableIEC(value, unit)
	return wrapHumanReadable(v, " ", u, e)
}

// HumanReadableSI colors the same way as Float() does but will make `n` human-readable using SI-prefixes (base1000).
//
// Related config setting(s):
//
//  - `LoggerConfig.AutoFloatPrecision`
//  - `LoggerConfig.ColorUnitHumanReadable` + exponent of the matching scale step
func HumanReadableSI[N Number](value N, unit string) string {
	v, e, u := humanReadableSI(value, unit)
	return wrapHumanReadable(v, " ", u, e)
}

// HumanReadableBytesIEC colors the same way as Float() does but will make `bytes` human-readable using IEC-prefixes (base1024).
//
// Related config setting(s):
//
//  - `LoggerConfig.AutoFloatPrecision`
//  - `LoggerConfig.ColorUnitHumanReadable` + exponent of the matching scale step
func HumanReadableBytesIEC[N Number](bytes N) string {
	return HumanReadableIEC(bytes, "B")
}

// HumanReadableBytesSI colors the same way as Float() does but will make `bytes` human-readable using SI-prefixes (base1000).
//
// Related config setting(s):
//
//  - `LoggerConfig.AutoFloatPrecision`
//  - `LoggerConfig.ColorUnitHumanReadable` + exponent of the matching scale step
func HumanReadableBytesSI[N Number](bytes N) string {
	return HumanReadableSI(bytes, "B")
}

// HumanReadableRateIEC colors the same way as Float() does but will make `n` human-readable using IEC-prefixes (base1024).
//
// `interval` defines the duration during which `n` accumulated.
//
// Related config setting(s):
//
//  - `LoggerConfig.AutoFloatPrecision`
//  - `LoggerConfig.ColorUnitHumanReadable` + exponent of the matching scale step
func HumanReadableRateIEC[N Number](n N, unit, interval string) string {
	v, e, u := humanReadableIEC(n, unit)
	return wrapHumanReadable(v, " ", u+"/"+interval, e)
}

// HumanReadableRateIEC colors the same way as Float() does but will make `n` human-readable using IEC-prefixes (base1024).
//
// `interval` defines the duration during which `n` accumulated.
//
// Related config setting(s):
//
//  - `LoggerConfig.AutoFloatPrecision`
//  - `LoggerConfig.ColorUnitHumanReadable` + exponent of the matching scale step
func HumanReadableRateSI[N Number](n N, unit, interval string) string {
	v, e, u := humanReadableSI(n, unit)
	return wrapHumanReadable(v, " ", u+"/"+interval, e)
}

// HumanReadableRateBytesIEC colors the same way as Float() does but will make `bytes` human-readable using IEC-prefixes (base1024).
//
// `interval` defines the duration during which `bytes` accumulated.
//
// Related config setting(s):
//
//  - `LoggerConfig.AutoFloatPrecision`
//  - `LoggerConfig.ColorUnitHumanReadable` + exponent of the matching scale step
func HumanReadableRateBytesIEC[N Number](bytes N, interval string) string {
	return HumanReadableRateIEC(bytes, "B", interval)
}

// HumanReadableRateBytesSI colors the same way as Float() does but will make `bytes` human-readable using SI-prefixes (base1000).
//
// `interval` defines the duration during which `bytes` accumulated.
//
// Related config setting(s):
//
//  - `LoggerConfig.AutoFloatPrecision`
//  - `LoggerConfig.ColorUnitHumanReadable` + exponent of the matching scale step
func HumanReadableRateBytesSI[N Number](bytes N, interval string) string {
	return HumanReadableRateSI(bytes, "B", interval)
}
