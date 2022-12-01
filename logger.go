package glog

import (
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/toxyl/gutils"
)

var reverseDNSResults map[string]string = map[string]string{}

var (
	DefaultTimeFormat         = "15:04:05"
	DefaultTimeFormat12hr     = "03:04:05pm"
	DefaultDateFormat         = "2006-01-02"
	DefaultDateTimeFormat     = fmt.Sprintf("%s %s", DefaultDateFormat, DefaultTimeFormat)
	DefaultDateTimeFormat12hr = fmt.Sprintf("%s %s", DefaultDateFormat, DefaultTimeFormat12hr)
	startTime                 = time.Now()
	showRuntime               = false
	showDateTime              = false
)

// ShowRuntimes enables prefixing log messages with the time (in seconds) since program start.
// This setting applies to all Logger instances.
func ShowRuntimes() {
	showRuntime = true
}

// HideRuntimes disables prefixing log messages with the time (in seconds) since program start.
// This setting applies to all Logger instances.
func HideRuntimes() {
	showRuntime = false
}

// ShowDateTimes enables prefixing log messages with the current date and time.
// This setting applies to all Logger instances.
func ShowDateTimes() {
	showDateTime = true
}

// HideDateTimes disables prefixing log messages with the current date and time.
// This setting applies to all Logger instances.
func HideDateTimes() {
	showDateTime = false
}

func AddrHostPort(host string, port int, useReverseDNS bool) string {
	host = enrichAndColorHost(host, useReverseDNS)
	return fmt.Sprintf("%s:%s", host, Port(port))
}

func Addr(addr string, useReverseDNS bool) string {
	h, p := gutils.SplitHostPort(addr)
	return AddrHostPort(h, p, useReverseDNS)
}

func ConnRemote(conn net.Conn, useReverseDNS bool) string {
	return Addr(conn.RemoteAddr().String(), useReverseDNS)
}

func ConnLocal(conn net.Conn, useReverseDNS bool) string {
	return Addr(conn.LocalAddr().String(), useReverseDNS)
}

func Wrap(str string, color uint) string {
	return fmt.Sprintf("\033[38;5;%dm%s\033[0m", color, str)
}

func Port(port int) string {
	// 94 - 231 (137 total)
	return Wrap(fmt.Sprint(port), uint(94.0+137.0*(float64(port)/65535.0)))
}

func enrichAndColorHost(host string, useReverseDNS bool) string {
	parts := strings.Split(host, ".")
	pt := 0.0
	for _, p := range parts {
		f, _ := gutils.GetFloat(p)
		pt += f
	}
	revDNS := "N/A"
	if useReverseDNS {
		if _, ok := reverseDNSResults[host]; !ok {
			reverseDNSResults[host] = gutils.ReverseDNS(host)
		}
		revDNS = reverseDNSResults[host]
	}
	hostColor := uint(88.0 + 143.0*(pt/4.0/255.0)) // 88 - 231 (143 total)

	if revDNS != "N/A" {
		host = fmt.Sprintf("%s (%s)", host, revDNS)
	}
	host = Wrap(host, hostColor)

	return host
}

func Hosts(hosts []string, useReverseDNS bool) string {
	hs := []string{}
	for _, h := range hosts {
		hs = append(hs, enrichAndColorHost(h, useReverseDNS))
	}
	return strings.Join(hs, ", ")
}

func Password(password string) string {
	return Wrap(password, Green)
}

func Error(err error) string {
	return Wrap(err.Error(), Orange)
}

func Reason(reason string) string {
	return Wrap(reason, Orange)
}

func File(file string) string {
	return Wrap(file, LightBlue)
}

func Highlight(message string) string {
	return Wrap(message, Cyan)
}

func Duration(seconds uint) string {
	return Wrap(time.Duration(seconds*uint(time.Second)).String(), Cyan)
}

func DurationMilliseconds(milliseconds uint) string {
	return Wrap(time.Duration(milliseconds*uint(time.Millisecond)).String(), Cyan)
}

// TimeCustom formats `t` according to the given format
// and colors the result green.
func TimeCustom(t time.Time, format string) string {
	return Wrap(t.Format(format), Green)
}

// Time12hr parses the time portion of `t`, formats it as AM/PM (03:04:05pm)
// and colors the result green.
// Overwrite `DefaultTimeFormat12hr` to use a different format.
func Time12hr(t time.Time) string {
	return TimeCustom(t, DefaultTimeFormat12hr)
}

// Time parses the time portion of `t`, formats it (15:04:05)
// and colors the result green.
// Overwrite `DefaultTimeFormat` to use a different format.
func Time(t time.Time) string {
	return TimeCustom(t, DefaultTimeFormat)
}

// Date parses the date portion of `t`, formats it (2006-01-02)
// and colors the result green.
// Overwrite `DefaultDateFormat` to use a different format.
func Date(t time.Time) string {
	return TimeCustom(t, DefaultDateFormat)
}

// DateTime parses `t`, formats it (2006-01-02 15:04:05)
// and colors the result green.
// Overwrite `DefaultDateTimeFormat` to use a different format.
func DateTime(t time.Time) string {
	return TimeCustom(t, DefaultDateTimeFormat)
}

// DateTime12hr parses `t`, formats it as AM/PM (2006-01-02 03:04:05pm)
// and colors the result green.
// Overwrite `DefaultDateTimeFormat12hr` to use a different format.
func DateTime12hr(t time.Time) string {
	return TimeCustom(t, DefaultDateTimeFormat12hr)
}

// Timestamp uses the current time, formats it as Unix timestamp (seconds)
// and colors the result green.
func Timestamp() string {
	return Wrap(fmt.Sprint(time.Now().Unix()), Green)
}

// Runtime determines the number of seconds passed since program start
// and colors the result green.
func Runtime() string {
	return Wrap(fmt.Sprint(int(time.Since(startTime).Seconds())), Green)
}

// Percentage assumes `n` to be normalized (0..1), multiplies it with 100,
// formats it with the given `precision` and colors the result cyan (`n` > 0), blue (`n` == 0) or red (`n` < 0).
func Percentage(n float64, precision int) string {
	color := uint(Cyan)
	if n < 0 {
		color = Red
	} else if n == 0 {
		color = Blue
	}
	return Wrap(fmt.Sprintf(fmt.Sprintf("%%.%df%%%%", precision), n*100.0), color)
}

// Float64 formats `n` with the given `precision` and colors the result cyan (`n` > 0), blue (`n` == 0) or red (`n` < 0).
func Float64(n float64, precision int) string {
	color := uint(Cyan)
	if n < 0 {
		color = Red
	} else if n == 0 {
		color = Blue
	}
	return Wrap(fmt.Sprintf(fmt.Sprintf("%%.%df", precision), n), color)
}

// Float32 formats `n` with the given `precision` and colors the result cyan (`n` > 0), blue (`n` == 0) or red (`n` < 0).
func Float32(n float32, precision int) string {
	return Float64(float64(n), precision)
}

// Int colors the result cyan (`n` > 0), blue (`n` == 0) or red (`n` < 0).
func Int(n int) string {
	color := uint(Cyan)
	if n < 0 {
		color = Red
	} else if n == 0 {
		color = Blue
	}
	return Wrap(fmt.Sprintf("%d", n), color)
}

// Int8 colors the result cyan (`n` > 0), blue (`n` == 0) or red (`n` < 0).
func Int8(n int8) string {
	return Int(int(n))
}

// Int16 colors the result cyan (`n` > 0), blue (`n` == 0) or red (`n` < 0).
func Int16(n int16) string {
	return Int(int(n))
}

// Int32 colors the result cyan (`n` > 0), blue (`n` == 0) or red (`n` < 0).
func Int32(n int32) string {
	return Int(int(n))
}

// Int64 colors the result cyan (`n` > 0), blue (`n` == 0) or red (`n` < 0).
func Int64(n int64) string {
	return Int(int(n))
}

// Uint colors the result cyan (`n` > 0), blue (`n` == 0) or red (`n` < 0).
func Uint(n uint) string {
	color := uint(Cyan)
	if n == 0 {
		color = Blue
	}
	return Wrap(fmt.Sprintf("%d", n), color)
}

// Uint8 colors the result cyan (`n` > 0), blue (`n` == 0) or red (`n` < 0).
func Uint8(n uint8) string {
	return Uint(uint(n))
}

// Uint16 colors the result cyan (`n` > 0), blue (`n` == 0) or red (`n` < 0).
func Uint16(n uint16) string {
	return Uint(uint(n))
}

// Uint32 colors the result cyan (`n` > 0), blue (`n` == 0) or red (`n` < 0).
func Uint32(n uint32) string {
	return Uint(uint(n))
}

// Uint64 colors the result cyan (`n` > 0), blue (`n` == 0) or red (`n` < 0).
func Uint64(n uint64) string {
	return Uint(uint(n))
}

// IntAmount colors `n` as int and appends either the
// given singular (`n` == 1) or plural (`n` > 1 || `n` == 0).
func IntAmount(n int, singular, plural string) string {
	unit := singular
	if n > 1 {
		unit = plural
	}
	color := uint(Cyan)
	if n < 0 {
		color = Red
	} else if n == 0 {
		color = Blue
	}
	amount := Wrap(fmt.Sprintf("%d", n), color)
	if n == 0 {
		unit = plural
	}
	return fmt.Sprintf("%s %s", amount, unit)
}

const (
	DarkBlue     = 17
	Blue         = 21
	DarkGreen    = 22
	LightBlue    = 27
	OliveGreen   = 34
	Green        = 46
	Cyan         = 51
	Purple       = 53
	DarkOrange   = 130
	DarkYellow   = 142
	Lime         = 154
	DarkRed      = 160
	Red          = 196
	Pink         = 201
	Orange       = 208
	Yellow       = 220
	BrightYellow = 229
	DarkGray     = 234
	MediumGray   = 240
	Gray         = 250
)

type Logger struct {
	ID            string
	color         uint
	debug         bool
	hideSubsystem bool
	hideIndicator bool
	OnMessage     func(string)
}

func (l *Logger) write(indicator rune, format string, a ...interface{}) {
	if l.hideIndicator {
		if showRuntime {
			format = fmt.Sprintf("%s | %s", Runtime(), format)
		}
		if showDateTime {
			format = fmt.Sprintf("%s | %s", DateTime(time.Now()), format)
		}

		fmt.Printf(format+"\n", a...)
		return
	}

	prefix := "[ ]"
	switch indicator {
	case 'i':
		prefix = Wrap("[i]", LightBlue)
	case '+':
		prefix = Wrap("[+]", OliveGreen)
	case '✓':
		prefix = Wrap("[✓]", Green)
	case '-':
		prefix = Wrap("[-]", DarkRed)
	case 'x':
		prefix = Wrap("[x]", Red)
	case '!':
		prefix = Wrap("[!]", Orange)
	case 'd':
		prefix = Wrap("[D]", Orange)
	case ' ':
		prefix = Wrap("[ ]", Gray)
	}

	if showRuntime {
		prefix = fmt.Sprintf("%s | %s", Runtime(), prefix)
	}
	if showDateTime {
		prefix = fmt.Sprintf("%s | %s", DateTime(time.Now()), prefix)
	}

	msg := fmt.Sprintf(prefix+" "+format+"\n", a...)
	if l.OnMessage != nil {
		l.OnMessage(msg)
		return
	}
	fmt.Print(msg)
}

func (l *Logger) prependFormat(format string) string {
	if l.hideSubsystem {
		return format
	}
	return fmt.Sprintf("%s: %s", Wrap(fmt.Sprintf("%-16s", l.ID), l.color), format)
}

func (l *Logger) Default(format string, a ...interface{}) {
	l.write(' ', l.prependFormat(format), a...)
}

func (l *Logger) Info(format string, a ...interface{}) {
	l.write('i', l.prependFormat(format), a...)
}

func (l *Logger) Success(format string, a ...interface{}) {
	l.write('✓', l.prependFormat(format), a...)
}

func (l *Logger) OK(format string, a ...interface{}) {
	l.write('+', l.prependFormat(format), a...)
}

func (l *Logger) NotOK(format string, a ...interface{}) {
	l.write('-', l.prependFormat(format), a...)
}

func (l *Logger) Error(format string, a ...interface{}) {
	l.write('x', l.prependFormat(format), a...)
}

func (l *Logger) Warning(format string, a ...interface{}) {
	l.write('!', l.prependFormat(format), a...)
}

func (l *Logger) Debug(format string, a ...interface{}) {
	if !l.debug {
		return
	}
	l.write('d', l.prependFormat(format), a...)
}

func (l *Logger) EnableDebug() *Logger {
	l.debug = true
	return l
}

func (l *Logger) DisableDebug() *Logger {
	l.debug = false
	return l
}

func (l *Logger) HideSubsystem() *Logger {
	l.hideSubsystem = true
	return l
}

func (l *Logger) HideIndicator() *Logger {
	l.hideIndicator = true
	return l
}

func (l *Logger) ShowSubsystem() *Logger {
	l.hideSubsystem = false
	return l
}

func (l *Logger) ShowIndicator() *Logger {
	l.hideIndicator = false
	return l
}

// NewLogger creates a new logger instance. Pass `nil` as `messageHandler` if you want logs to be printed to screen,
// else provide your own handler.
func NewLogger(id string, color uint, debug, hideSubsystem, hideIndicator bool, messageHandler func(string)) *Logger {
	return &Logger{
		ID:            id,
		color:         color,
		debug:         debug,
		hideSubsystem: hideSubsystem,
		hideIndicator: hideIndicator,
		OnMessage:     messageHandler,
	}
}
