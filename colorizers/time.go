package colorizers

import (
	"fmt"
	"regexp"
	"time"

	"github.com/toxyl/glog/ansi"
	"github.com/toxyl/glog/config"
	"github.com/toxyl/glog/types"
	"github.com/toxyl/glog/utils"
)

var reDurationSeparators regexp.Regexp = *regexp.MustCompile(`([a-zA-Z]+)([0-9]+)`)

// Duration explodes the result of (time.Duration).String() into its segments and colors them.
//
// `seconds` can be int/uint/float (all representing seconds) or a time.Duration.
//
// Related config setting(s):
//
//   - `LoggerConfig.ColorDuration`
func Duration[D types.Durations](seconds D) string {
	str := time.Duration(seconds * D(time.Second)).String()
	str = reDurationSeparators.ReplaceAllString(str, "$1 $2")
	return ansi.Wrap(str, config.LoggerConfig.ColorDuration).String()
}

// DurationMilliseconds explodes the result of (time.Duration).String() into its segments and colors them.
//
// `milliseconds` can be int/uint/float (all representing milliseconds) or a time.Duration.
//
// Related config setting(s):
//
//   - `LoggerConfig.ColorDuration`
func DurationMilliseconds[D types.Durations](milliseconds D) string {
	str := time.Duration(milliseconds * D(time.Millisecond)).String()
	str = reDurationSeparators.ReplaceAllString(str, "$1 $2")
	return ansi.Wrap(str, config.LoggerConfig.ColorDuration).String()
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
func DurationShort[N types.Number](seconds N, scale utils.DurationScale) string {
	v, e, u := humanReadableDuration(seconds, scale)
	return wrapHumanReadable(v, "", u, e)
}

// TimeCustom formats `t` according to the given format.
func TimeCustom(t time.Time, format string) string {
	return ansi.Wrap(t.Format(format), config.LoggerConfig.ColorTime).String()
}

// Time12hr parses the time portion of `t`, formats it as AM/PM (03:04:05pm).
//
// Related config setting(s):
//
//   - `LoggerConfig.DefaultTimeFormat12hr`
func Time12hr(t time.Time) string {
	return TimeCustom(t, config.LoggerConfig.TimeFormat12hr)
}

// Time parses the time portion of `t`, formats it (15:04:05).
//
// Related config setting(s):
//
//   - `LoggerConfig.DefaultTimeFormat`
func Time(t time.Time) string {
	return TimeCustom(t, config.LoggerConfig.TimeFormat)
}

// Date parses the date portion of `t`, formats it (2006-01-02).
//
// Related config setting(s):
//
//   - `LoggerConfig.DefaultDateFormat`
func Date(t time.Time) string {
	return TimeCustom(t, config.LoggerConfig.DateFormat)
}

// DateTime parses `t`, formats it (2006-01-02 15:04:05).
//
// Related config setting(s):
//
//   - `LoggerConfig.DefaultDateTimeFormat`
func DateTime(t time.Time) string {
	return TimeCustom(t, config.LoggerConfig.DateTimeFormat)
}

// DateTime12hr parses `t`, formats it as AM/PM (2006-01-02 03:04:05pm).
//
// Related config setting(s):
//
//   - `LoggerConfig.DefaultDateTimeFormat12hr`
func DateTime12hr(t time.Time) string {
	return TimeCustom(t, config.LoggerConfig.DateTimeFormat12hr)
}

// Timestamp uses the current time, formats it as Unix timestamp (seconds).
//
// Related config setting(s):
//
//   - `LoggerConfig.ColorTime`
func Timestamp() string {
	return ansi.Wrap(fmt.Sprint(time.Now().Unix()), config.LoggerConfig.ColorTime).String()
}

// Runtime determines the number of seconds passed since program start.
//
// Related config setting(s):
//
//   - `LoggerConfig.ColorDuration`
func Runtime() string {
	return ansi.Wrap(fmt.Sprint(int(time.Since(config.LoggerConfig.CreatedAt).Seconds())), config.LoggerConfig.ColorDuration).String()
}

// RuntimeHumanReadable determines the number of seconds passed since program start.
//
// Related config setting(s):
//
//   - `LoggerConfig.ColorDuration`
func RuntimeHumanReadable() string {
	return Duration(int(time.Since(config.LoggerConfig.CreatedAt).Seconds()))
}

// RuntimeSeconds determines the number of seconds passed since program start.
//
// Related config setting(s):
//
//   - `LoggerConfig.ColorDuration`
func RuntimeSeconds() string {
	return ansi.Wrap(fmt.Sprint(int(time.Since(config.LoggerConfig.CreatedAt).Seconds())), config.LoggerConfig.ColorDuration).String()
}

// RuntimeMilliseconds determines the number of milliseconds passed since program start.
//
// Related config setting(s):
//
//   - `LoggerConfig.ColorDuration`
func RuntimeMilliseconds() string {
	return ansi.Wrap(fmt.Sprint(int(time.Since(config.LoggerConfig.CreatedAt).Milliseconds())), config.LoggerConfig.ColorDuration).String()
}
