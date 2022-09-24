package glog

import (
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/toxyl/gutils"
)

var reverseDNSResults map[string]string = map[string]string{}

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

func Int(n int) string {
	return Wrap(fmt.Sprintf("%d", n), Cyan)
}

// IntAmount colors `n` as int and appends either the
// given singular (`n` == 1) or plural (`n` > 1 || `n` == 0).
func IntAmount(n int, singular, plural string) string {
	unit := singular
	if n > 1 {
		unit = plural
	}
	amount := Wrap(fmt.Sprintf("%d", n), Cyan)
	if n == 0 {
		amount = Wrap("0", Orange)
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

// NewLogger creates a new logger instance. Pass `nil` as `messageHandler` if you want logs to printed to screen,
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
