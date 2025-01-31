package glog

import (
	"fmt"
	"regexp"
	"time"
)

var reDurationSeparators regexp.Regexp = *regexp.MustCompile(`([a-zA-Z]+)([0-9]+)`)

// Duration explodes the result of (time.Duration).String() into its segments and colors them.
//
// `seconds` can be int/uint/float (all representing seconds) or a time.Duration.
//
// Related config setting(s):
//
//   - `LoggerConfig.ColorDuration`
func Duration[D Durations](seconds D) string {
	str := time.Duration(seconds * D(time.Second)).String()
	str = reDurationSeparators.ReplaceAllString(str, "$1 $2")
	return Wrap(str, LoggerConfig.ColorDuration)
}

// DurationMilliseconds explodes the result of (time.Duration).String() into its segments and colors them.
//
// `milliseconds` can be int/uint/float (all representing milliseconds) or a time.Duration.
//
// Related config setting(s):
//
//   - `LoggerConfig.ColorDuration`
func DurationMilliseconds[D Durations](milliseconds D) string {
	str := time.Duration(milliseconds * D(time.Millisecond)).String()
	str = reDurationSeparators.ReplaceAllString(str, "$1 $2")
	return Wrap(str, LoggerConfig.ColorDuration)
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
	v, e, u := humanReadableDuration(seconds, scale)
	return wrapHumanReadable(v, "", u, e)
}

// TimeCustom formats `t` according to the given format.
func TimeCustom(t time.Time, format string) string {
	return Wrap(t.Format(format), LoggerConfig.ColorTime)
}

// Time12hr parses the time portion of `t`, formats it as AM/PM (03:04:05pm).
//
// Related config setting(s):
//
//   - `LoggerConfig.DefaultTimeFormat12hr`
func Time12hr(t time.Time) string {
	return TimeCustom(t, LoggerConfig.TimeFormat12hr)
}

// Time parses the time portion of `t`, formats it (15:04:05).
//
// Related config setting(s):
//
//   - `LoggerConfig.DefaultTimeFormat`
func Time(t time.Time) string {
	return TimeCustom(t, LoggerConfig.TimeFormat)
}

// Date parses the date portion of `t`, formats it (2006-01-02).
//
// Related config setting(s):
//
//   - `LoggerConfig.DefaultDateFormat`
func Date(t time.Time) string {
	return TimeCustom(t, LoggerConfig.DateFormat)
}

// DateTime parses `t`, formats it (2006-01-02 15:04:05).
//
// Related config setting(s):
//
//   - `LoggerConfig.DefaultDateTimeFormat`
func DateTime(t time.Time) string {
	return TimeCustom(t, LoggerConfig.DateTimeFormat)
}

// DateTime12hr parses `t`, formats it as AM/PM (2006-01-02 03:04:05pm).
//
// Related config setting(s):
//
//   - `LoggerConfig.DefaultDateTimeFormat12hr`
func DateTime12hr(t time.Time) string {
	return TimeCustom(t, LoggerConfig.DateTimeFormat12hr)
}

// Timestamp uses the current time, formats it as Unix timestamp (seconds).
//
// Related config setting(s):
//
//   - `LoggerConfig.ColorTime`
func Timestamp() string {
	return Wrap(fmt.Sprint(time.Now().Unix()), LoggerConfig.ColorTime)
}

// Runtime determines the number of seconds passed since program start.
//
// Related config setting(s):
//
//   - `LoggerConfig.ColorDuration`
func Runtime() string {
	return Wrap(fmt.Sprint(int(time.Since(LoggerConfig.createdAt).Seconds())), LoggerConfig.ColorDuration)
}

// RuntimeHumanReadable determines the number of seconds passed since program start.
//
// Related config setting(s):
//
//   - `LoggerConfig.ColorDuration`
func RuntimeHumanReadable() string {
	return Duration(int(time.Since(LoggerConfig.createdAt).Seconds()))
}

// RuntimeMilliseconds determines the number of milliseconds passed since program start.
//
// Related config setting(s):
//
//   - `LoggerConfig.ColorDuration`
func RuntimeMilliseconds() string {
	return Wrap(fmt.Sprint(int(time.Since(LoggerConfig.createdAt).Milliseconds())), LoggerConfig.ColorDuration)
}
