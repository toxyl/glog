package glog

import (
	"github.com/toxyl/glog/colorizers"
)

// Backwards compatibility aliases and / or equivalent replacements

var (
	Auto              = colorizers.Auto
	Bool              = colorizers.Bool
	Highlight         = colorizers.Highlight
	HighlightInfo     = colorizers.HighlightInfo
	HighlightOK       = colorizers.HighlightOK
	HighlightSuccess  = colorizers.HighlightSuccess
	HighlightNotOK    = colorizers.HighlightNotOK
	HighlightError    = colorizers.HighlightError
	HighlightWarning  = colorizers.HighlightWarning
	HighlightDebug    = colorizers.HighlightDebug
	HighlightQuestion = colorizers.HighlightQuestion
	HighlightTrace    = colorizers.HighlightTrace
	// Password colors `password` according to the config. It does not redact it!
	//
	// Related config setting(s):
	//
	//  - `LoggerConfig.ColorPassword`
	Password = colorizers.Password

	// Error colors `err.Error()` according to the config, or returns "nil" if no error was present.
	//
	// Related config setting(s):
	//
	//   - `LoggerConfig.ColorError`
	//   - `LoggerConfig.ColorNil`
	Error = colorizers.Error

	// Reason colors `reason` according to the config.
	//
	// Related config setting(s):
	//
	//   - `LoggerConfig.ColorReason`
	Reason = colorizers.Reason

	// File colors `file` according to the config using the `os.PathSeparator`.
	//
	// Related config setting(s):
	//
	//   - `LoggerConfig.ColorPathSeparator`
	//   - `LoggerConfig.ColorPath“
	File       = colorizers.File
	Addr       = colorizers.Addr
	ConnRemote = colorizers.ConnRemote
	ConnLocal  = colorizers.ConnLocal
	IPs        = colorizers.IPs
	// URL colorizes a URL and, if enabled, marks dead ones (based on a DNS check).
	//
	// Related config setting(s):
	//
	//   - `LoggerConfig.ColorURLSeparators`
	//   - `LoggerConfig.ColorScheme`
	//   - `LoggerConfig.ColorUser`
	//   - `LoggerConfig.ColorPassword`
	//   - `LoggerConfig.ColorURLPath`
	//   - `LoggerConfig.ColorQueryKey`
	//   - `LoggerConfig.ColorQueryValue`
	//   - `LoggerConfig.ColorFragment`
	//   - `LoggerConfig.CheckIfURLIsAlive`
	URL = colorizers.URL
	// TimeCustom formats `t` according to the given format.
	TimeCustom = colorizers.TimeCustom

	// Time12hr parses the time portion of `t`, formats it as AM/PM (03:04:05pm).
	//
	// Related config setting(s):
	//
	//   - `LoggerConfig.DefaultTimeFormat12hr`
	Time12hr = colorizers.Time12hr

	// Time parses the time portion of `t`, formats it (15:04:05).
	//
	// Related config setting(s):
	//
	//   - `LoggerConfig.DefaultTimeFormat`
	Time = colorizers.Time

	// Date parses the date portion of `t`, formats it (2006-01-02).
	//
	// Related config setting(s):
	//
	//   - `LoggerConfig.DefaultDateFormat`
	Date = colorizers.Date

	// DateTime parses `t`, formats it (2006-01-02 15:04:05).
	//
	// Related config setting(s):
	//
	//   - `LoggerConfig.DefaultDateTimeFormat`
	DateTime = colorizers.DateTime

	// DateTime12hr parses `t`, formats it as AM/PM (2006-01-02 03:04:05pm).
	//
	// Related config setting(s):
	//
	//   - `LoggerConfig.DefaultDateTimeFormat12hr`
	DateTime12hr = colorizers.DateTime12hr

	// Timestamp uses the current time, formats it as Unix timestamp (seconds).
	//
	// Related config setting(s):
	//
	//   - `LoggerConfig.ColorTime`
	Timestamp = colorizers.Timestamp

	// Runtime determines the number of seconds passed since program start.
	//
	// Related config setting(s):
	//
	//   - `LoggerConfig.ColorDuration`
	Runtime = colorizers.Runtime

	// RuntimeHumanReadable determines the number of seconds passed since program start.
	//
	// Related config setting(s):
	//
	//   - `LoggerConfig.ColorDuration`
	RuntimeHumanReadable = colorizers.RuntimeHumanReadable

	// RuntimeMilliseconds determines the number of milliseconds passed since program start.
	//
	// Related config setting(s):
	//
	//   - `LoggerConfig.ColorDuration`
	RuntimeMilliseconds = colorizers.RuntimeMilliseconds

	WrapDarkBlue     = colorizers.WrapDarkBlue
	WrapBlue         = colorizers.WrapBlue
	WrapDarkGreen    = colorizers.WrapDarkGreen
	WrapLightBlue    = colorizers.WrapLightBlue
	WrapOliveGreen   = colorizers.WrapOliveGreen
	WrapGreen        = colorizers.WrapGreen
	WrapCyan         = colorizers.WrapCyan
	WrapPurple       = colorizers.WrapPurple
	WrapDarkOrange   = colorizers.WrapDarkOrange
	WrapDarkYellow   = colorizers.WrapDarkYellow
	WrapLime         = colorizers.WrapLime
	WrapDarkRed      = colorizers.WrapDarkRed
	WrapRed          = colorizers.WrapRed
	WrapPink         = colorizers.WrapPink
	WrapOrange       = colorizers.WrapOrange
	WrapYellow       = colorizers.WrapYellow
	WrapBrightYellow = colorizers.WrapBrightYellow
	WrapDarkGray     = colorizers.WrapDarkGray
	WrapMediumGray   = colorizers.WrapMediumGray
	WrapGray         = colorizers.WrapGray
	WrapWhite        = colorizers.WrapWhite
)

// Percentage assumes `n` to be normalized (0..1), multiplies it with 100,
// formats it with the given `precision` and colors the result.
//
// Related config setting(s):
//
//   - `n`  < 0: `LoggerConfig.ColorFloatNegative`
//   - `n` == 0: `LoggerConfig.ColorFloatZero`
//   - `n`  > 0: `LoggerConfig.ColorFloatPositive`
func Percentage[F Floats](n F, precision int) string {
	return colorizers.Percentage(n, precision)
}

// Float formats `n` with the given `precision` and colors the result.
//
// Related config setting(s):
//
//   - `n`  < 0: `LoggerConfig.ColorFloatNegative`
//   - `n` == 0: `LoggerConfig.ColorFloatZero`
//   - `n`  > 0: `LoggerConfig.ColorFloatPositive`
func Float[F Floats](n F, precision int) string {
	return colorizers.Float(n, precision)
}

// HumanReadableShort colors the same way as Float() does but will make `n` human-readable using short scale suffixes (base1000).
//
// Related config setting(s):
//
//   - `LoggerConfig.AutoFloatPrecision`
//   - `LoggerConfig.ColorUnitHumanReadable` + exponent of the matching scale step
func HumanReadableShort[N Number](value N) string { return colorizers.HumanReadableShort(value) }

// HumanReadableLong colors the same way as Float() does but will make `n` human-readable using long scale suffixes  (base1000).
//
// Related config setting(s):
//
//   - `LoggerConfig.AutoFloatPrecision`
//   - `LoggerConfig.ColorUnitHumanReadable` + exponent of the matching scale step
func HumanReadableLong[N Number](value N) string { return colorizers.HumanReadableLong(value) }

// HumanReadableIEC colors the same way as Float() does but will make `n` human-readable using IEC-prefixes (base1024).
//
// Related config setting(s):
//
//   - `LoggerConfig.AutoFloatPrecision`
//   - `LoggerConfig.ColorUnitHumanReadable` + exponent of the matching scale step
func HumanReadableIEC[N Number](value N, unit string) string {
	return colorizers.HumanReadableIEC(value, unit)
}

// HumanReadableSI colors the same way as Float() does but will make `n` human-readable using SI-prefixes (base1000).
//
// Related config setting(s):
//
//   - `LoggerConfig.AutoFloatPrecision`
//   - `LoggerConfig.ColorUnitHumanReadable` + exponent of the matching scale step
func HumanReadableSI[N Number](value N, unit string) string {
	return colorizers.HumanReadableSI(value, unit)
}

// HumanReadableBytesIEC colors the same way as Float() does but will make `bytes` human-readable using IEC-prefixes (base1024).
//
// Related config setting(s):
//
//   - `LoggerConfig.AutoFloatPrecision`
//   - `LoggerConfig.ColorUnitHumanReadable` + exponent of the matching scale step
func HumanReadableBytesIEC[N Number](bytes N) string { return colorizers.HumanReadableBytesIEC(bytes) }

// HumanReadableBytesSI colors the same way as Float() does but will make `bytes` human-readable using SI-prefixes (base1000).
//
// Related config setting(s):
//
//   - `LoggerConfig.AutoFloatPrecision`
//   - `LoggerConfig.ColorUnitHumanReadable` + exponent of the matching scale step
func HumanReadableBytesSI[N Number](bytes N) string { return colorizers.HumanReadableBytesSI(bytes) }

// HumanReadableRateIEC colors the same way as Float() does but will make `n` human-readable using IEC-prefixes (base1024).
//
// `interval` defines the duration during which `n` accumulated.
//
// Related config setting(s):
//
//   - `LoggerConfig.AutoFloatPrecision`
//   - `LoggerConfig.ColorUnitHumanReadable` + exponent of the matching scale step
func HumanReadableRateIEC[N Number](n N, unit, interval string) string {
	return colorizers.HumanReadableRateIEC(n, unit, interval)
}

// HumanReadableRateIEC colors the same way as Float() does but will make `n` human-readable using IEC-prefixes (base1024).
//
// `interval` defines the duration during which `n` accumulated.
//
// Related config setting(s):
//
//   - `LoggerConfig.AutoFloatPrecision`
//   - `LoggerConfig.ColorUnitHumanReadable` + exponent of the matching scale step
func HumanReadableRateSI[N Number](n N, unit, interval string) string {
	return colorizers.HumanReadableRateSI(n, unit, interval)
}

// HumanReadableRateBytesIEC colors the same way as Float() does but will make `bytes` human-readable using IEC-prefixes (base1024).
//
// `interval` defines the duration during which `bytes` accumulated.
//
// Related config setting(s):
//
//   - `LoggerConfig.AutoFloatPrecision`
//   - `LoggerConfig.ColorUnitHumanReadable` + exponent of the matching scale step
func HumanReadableRateBytesIEC[N Number](bytes N, interval string) string {
	return colorizers.HumanReadableRateBytesIEC(bytes, interval)
}

// HumanReadableRateBytesSI colors the same way as Float() does but will make `bytes` human-readable using SI-prefixes (base1000).
//
// `interval` defines the duration during which `bytes` accumulated.
//
// Related config setting(s):
//
//   - `LoggerConfig.AutoFloatPrecision`
//   - `LoggerConfig.ColorUnitHumanReadable` + exponent of the matching scale step
func HumanReadableRateBytesSI[N Number](bytes N, interval string) string {
	return colorizers.HumanReadableRateBytesSI(bytes, interval)
}

// IntAmount colors `n` as int and appends either the
// given singular (`n` == 1) or plural (`n` > 1 || `n` == 0).
//
// Related config setting(s):
//
//   - `n`  > 0: `LoggerConfig.ColorIntPositive`
//   - `n` == 0: `LoggerConfig.ColorIntZero`
//   - `n`  < 0: `LoggerConfig.ColorIntNegative`
func IntAmount[I IntOrUint](n I, singular, plural string) string {
	return colorizers.IntAmount(n, singular, plural)
}

// Int colors the given ints.
//
// Related config setting(s):
//
//   - `n`  > 0: `LoggerConfig.ColorIntPositive`
//   - `n` == 0: `LoggerConfig.ColorIntZero`
//   - `n`  < 0: `LoggerConfig.ColorIntNegative`
func Int[I IntOrUint](n ...I) string {
	return colorizers.Int(n...)
}

func AddrIPv4Port[I IntOrUint](ip string, port I, useReverseDNS bool) string {
	return colorizers.AddrIPv4Port(ip, port, useReverseDNS)
}

func Port[I IntOrUint](port I) string {
	return colorizers.Port(port)
}

// Duration explodes the result of (time.Duration).String() into its segments and colors them.
//
// `seconds` can be int/uint/float (all representing seconds) or a time.Duration.
//
// Related config setting(s):
//
//   - `LoggerConfig.ColorDuration`
func Duration[D Durations](seconds D) string { return colorizers.Duration(seconds) }

// DurationMilliseconds explodes the result of (time.Duration).String() into its segments and colors them.
//
// `milliseconds` can be int/uint/float (all representing milliseconds) or a time.Duration.
//
// Related config setting(s):
//
//   - `LoggerConfig.ColorDuration`
func DurationMilliseconds[D Durations](milliseconds D) string {
	return colorizers.DurationMilliseconds(milliseconds)
}

// DurationShort colors the same way as Float() does but will make `n` human-readable using time suffixes.
//
// This function, unlike `Duration` and `DurationMilliseconds`, only accepts int/uint/float values.
// time.Duration is not allowed to avoid overflow after a couple hundred years and
// to preserve the ability to process negative values.
// Use (time.Duration).Seconds() instead.
//
// Use `scale` to define how to treat intervals of months (always 1/12 of a year) and years:
//
//   - `YEAR_COMMON` = 365 days
//   - `YEAR_LEAP` = 366 days
//   - `YEAR_AVERAGE` = (3 * `YEAR_COMMON` + 1 * `YEAR_LEAP`) / 4
//
// Related config setting(s):
//
//   - `LoggerConfig.AutoFloatPrecision`
func DurationShort[N Number](seconds N, scale DurationScale) string {
	return colorizers.DurationShort(seconds, scale)
}

// Uint colors the given uints.
//
// Related config setting(s):
//
//   - `n`  > 0: `LoggerConfig.ColorUintPositive`
//   - `n` == 0: `LoggerConfig.ColorUintZero`
func Uint[U Uints](n ...U) string {
	return colorizers.Uint(n...)
}
