package glog

import (
	"fmt"
	"time"
)

type Logger struct {
	ID        string
	color     int
	debugMode bool
	onMessage func(string)
}

func (l *Logger) EnableDebug() {
	l.debugMode = true
}

func (l *Logger) DisableDebug() {
	l.debugMode = false
}

func (l *Logger) write(indicator rune, format string, a ...interface{}) {
	prefix := ""
	if LoggerConfig.ShowIndicator {
		if vi, ok := LoggerConfig.Indicators[indicator]; ok {
			prefix = Wrap(vi.value, vi.color)
		}
	}

	if LoggerConfig.ShowRuntimeSeconds {
		prefix = fmt.Sprintf("%22s s %s", Runtime(), prefix)
	}
	if LoggerConfig.ShowRuntimeMilliseconds {
		prefix = fmt.Sprintf("%22s ms %s", RuntimeMilliseconds(), prefix)
	}
	if LoggerConfig.ShowDateTime {
		prefix = fmt.Sprintf("%s %s", DateTime(time.Now()), prefix)
	}

	msg := fmt.Sprintf(prefix+" "+format+"\n", a...)

	if l.onMessage != nil {
		l.onMessage(msg)
		return
	}

	if LoggerConfig.ColorsDisabled {
		msg = StripANSI(msg)
	}
	fmt.Print(msg)
}

func (l *Logger) prependFormat(format string) string {
	if !LoggerConfig.ShowSubsystem {
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
	if !l.debugMode {
		return
	}
	l.write('d', l.prependFormat(format), a...)
}

// NewLogger creates a new logger instance. Pass `nil` as `messageHandler` if you want logs to be printed to screen,
// else provide your own handler.
func NewLogger(id string, color int, debugMode bool, messageHandler func(string)) *Logger {
	return &Logger{
		ID:        id,
		color:     color,
		debugMode: debugMode,
		onMessage: messageHandler,
	}
}
