package glog

import (
	"fmt"
	"time"
)

func Duration(seconds uint) string {
	return Wrap(time.Duration(seconds*uint(time.Second)).String(), LoggerConfig.ColorDuration)
}

func DurationMilliseconds(milliseconds uint) string {
	return Wrap(time.Duration(milliseconds*uint(time.Millisecond)).String(), LoggerConfig.ColorDuration)
}

// TimeCustom formats `t` according to the given format.
func TimeCustom(t time.Time, format string) string {
	return Wrap(t.Format(format), LoggerConfig.ColorTime)
}

// Time12hr parses the time portion of `t`, formats it as AM/PM (03:04:05pm).
// Overwrite `DefaultTimeFormat12hr` to use a different format.
func Time12hr(t time.Time) string {
	return TimeCustom(t, LoggerConfig.TimeFormat12hr)
}

// Time parses the time portion of `t`, formats it (15:04:05).
// Overwrite `DefaultTimeFormat` to use a different format.
func Time(t time.Time) string {
	return TimeCustom(t, LoggerConfig.TimeFormat)
}

// Date parses the date portion of `t`, formats it (2006-01-02).
// Overwrite `DefaultDateFormat` to use a different format.
func Date(t time.Time) string {
	return TimeCustom(t, LoggerConfig.DateFormat)
}

// DateTime parses `t`, formats it (2006-01-02 15:04:05).
// Overwrite `DefaultDateTimeFormat` to use a different format.
func DateTime(t time.Time) string {
	return TimeCustom(t, LoggerConfig.DateTimeFormat)
}

// DateTime12hr parses `t`, formats it as AM/PM (2006-01-02 03:04:05pm).
// Overwrite `DefaultDateTimeFormat12hr` to use a different format.
func DateTime12hr(t time.Time) string {
	return TimeCustom(t, LoggerConfig.DateTimeFormat12hr)
}

// Timestamp uses the current time, formats it as Unix timestamp (seconds).
func Timestamp() string {
	return Wrap(fmt.Sprint(time.Now().Unix()), LoggerConfig.ColorTime)
}

// Runtime determines the number of seconds passed since program start.
func Runtime() string {
	return Wrap(fmt.Sprint(int(time.Since(LoggerConfig.createdAt).Seconds())), LoggerConfig.ColorDuration)
}

// RuntimeMilliseconds determines the number of milliseconds passed since program start.
func RuntimeMilliseconds() string {
	return Wrap(fmt.Sprint(int(time.Since(LoggerConfig.createdAt).Milliseconds())), LoggerConfig.ColorDuration)
}
