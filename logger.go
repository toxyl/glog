package glog

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

type Logger struct {
	ID         string
	color      int
	debugMode  bool
	traceMode  bool
	traceLevel uint
	file       string
	fileColor  string
	onMessage  func(string)
}

// EnableTrace enables trace mode for the logger with the given trace level.
// When trace mode is enabled, calls to Logger.Trace(level) with a severity level
// equal to or higher than the specified trace level will trigger a stack trace print.
func (l *Logger) EnableTrace(level uint) {
	l.traceMode = true
	l.traceLevel = level
}

// DisableTrace disables trace mode for the logger.
func (l *Logger) DisableTrace() {
	l.traceMode = false
}

// EnableDebug enables debug mode for the logger.
// When debug mode is enabled, calls to Logger.Debug(...) will be printed to the output.
func (l *Logger) EnableDebug() {
	l.debugMode = true
}

// DisableDebug disables debug mode for the logger.
func (l *Logger) DisableDebug() {
	l.debugMode = false
}

// EnablePlainLog enables logging to a plain text file with the given file path.
// Each log message, stripped of ANSI escapes, will be appended to the file.
// File logging can be enabled/disabled on the fly.
func (l *Logger) EnablePlainLog(path string) {
	l.file = path
}

// DisablePlainLog stops logging plaintext messages to a file.
// File logging can be enabled/disabled on the fly.
func (l *Logger) DisablePlainLog() {
	l.file = ""
}

// EnableColorLog enables logging to a text file with the given file path.
// Each log message, together with ANSI escapes, will be appended to the file.
func (l *Logger) EnableColorLog(path string) {
	l.fileColor = path
}

// DisableColorLog stops logging messages to a file.
func (l *Logger) DisableColorLog() {
	l.fileColor = ""
}

// write logs a message to the console or file with an optional indicator,
// and applies various formatting options as specified in the logger's configuration.
//
// If a progress indicator ('p') is specified, it will be displayed as a progress bar
// and continuously replaced on the same line using ANSI escape sequences.
//
// The 'format' parameter is a string that can contain verbs, as specified by the fmt package,
// and the 'a' parameter provides the corresponding arguments for each verb.
//
// If a message handler function was specified during logger creation, the formatted message
// will be passed to it instead of being printed to the console or file.
//
// If the logger is configured to use colors, the message will include ANSI escape sequences
// to apply the appropriate colors for the message elements.
//
// If a file has been specified for the logger, the message will be appended to the file.
// If a color file has also been specified, the message with color codes will be appended to that file.
// If the logger's configuration disables colors, any color codes will be stripped from the message.
//
// Related config setting(s):
//
//   - LoggerConfig.ShowIndicator
//   - LoggerConfig.ShowDateTime
//   - LoggerConfig.ShowRuntimeSeconds
//   - LoggerConfig.ShowRuntimeMilliseconds
//   - LoggerConfig.ColorsDisabled
//   - LoggerConfig.Indicators
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
	if LoggerConfig.ShowSubsystem {
		prefix = fmt.Sprintf("%s %s: ", prefix, Wrap(fmt.Sprintf("%-16s", l.ID), l.color))
	}

	msg := fmt.Sprintf(format, a...)
	if LoggerConfig.SplitOnNewLine {
		res := []string{}
		for _, ln := range strings.Split(msg, "\n") {
			res = append(res, prefix+" "+ln)
		}
		msg = strings.Join(res, "\n")
	} else {
		msg = prefix + " " + msg
	}

	if indicator == 'p' {
		// the progress indicator is special, let's add some magic:
		msg = StoreCursor() + ClearToEOL() + msg + RestoreCursor()
	} else {
		msg += "\n"
	}

	if l.onMessage != nil {
		l.onMessage(msg)
		return
	}

	if LoggerConfig.ColorsDisabled {
		msg = StripANSI(msg)
	}

	if l.file != "" {
		err := os.MkdirAll(filepath.Dir(l.file), 0770) // create target dir if it doesn't exist
		if err != nil {
			panic(err)
		}

		f, err := os.OpenFile(l.file, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		if err != nil {
			panic(err)
		}

		defer f.Close()

		if _, err = f.WriteString(StripANSI(msg)); err != nil {
			panic(err)
		}
	}

	if l.fileColor != "" {
		err := os.MkdirAll(filepath.Dir(l.fileColor), 0770) // create target dir if it doesn't exist
		if err != nil {
			panic(err)
		}

		f, err := os.OpenFile(l.fileColor, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		if err != nil {
			panic(err)
		}

		defer f.Close()

		if _, err = f.WriteString(msg); err != nil {
			panic(err)
		}
	}

	fmt.Print(msg)
}

// Blank prints a message without any indicator such as "[ ]", "[i]", etc.
func (l *Logger) Blank(format string, a ...interface{}) {
	l.write('_', format, a...)
}

func (l *Logger) Default(format string, a ...interface{}) {
	l.write(' ', format, a...)
}

func (l *Logger) Info(format string, a ...interface{}) {
	l.write('i', format, a...)
}

func (l *Logger) Success(format string, a ...interface{}) {
	l.write('✓', format, a...)
}

func (l *Logger) OK(format string, a ...interface{}) {
	l.write('+', format, a...)
}

func (l *Logger) NotOK(format string, a ...interface{}) {
	l.write('-', format, a...)
}

func (l *Logger) Error(format string, a ...interface{}) {
	l.write('x', format, a...)
}

func (l *Logger) Warning(format string, a ...interface{}) {
	l.write('!', format, a...)
}

func (l *Logger) Debug(format string, a ...interface{}) {
	if !l.debugMode {
		return
	}
	l.write('d', format, a...)
}

func (l *Logger) Question(format string, a ...interface{}) {
	l.write('?', format, a...)
}

func (l *Logger) Trace(level uint) {
	if !l.traceMode || level > l.traceLevel {
		return
	}
	stackTracer.Sample(0).PrintWithLogger(l, 't')
}

func (l *Logger) ShowColors() {
	str := ""
	for i := 0; i < 8; i++ {
		str += Wrap(fmt.Sprintf("%03d ", i), i)
	}
	l.Default("%s", str)

	str = ""
	for i := 8; i < 16; i++ {
		str += Wrap(fmt.Sprintf("%03d ", i), i)
	}
	l.Default("%s", str)

	for i := 16; i < 256; i++ {
		str := ""
		for j := 0; j < 12 && i < 256; j++ {
			str += Wrap(fmt.Sprintf("%03d ", i), i)
			i++
		}
		i--
		l.Default("%s", str)
	}
}

func (l *Logger) Table(ats ...*TableColumn) {
	NewTable(ats...).Print(l)
}

func (l *Logger) TableWithoutHeader(ats ...*TableColumn) {
	NewTable(ats...).PrintWithoutHeader(l)
}

func (l *Logger) KeyValueTable(data map[string]interface{}) {
	atsKeys := NewTableColumnLeft("Key")
	atsValues := NewTableColumnLeft("Value")
	keys := []string{}

	for k := range data {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, k := range keys {
		atsKeys.Push(k)
		atsValues.Push(data[k])
	}

	NewTable(atsKeys, atsValues).Print(l)
}

// Progress prints a progress bar followed by the given format and arguments.
// This uses ANSI escapes to continuously replace the same line on the terminal.
// Use the methods ProgressSuccess and ProgressError to end a progress.
// This will print a message of the corresponding type and advance to the next line.
//
// It uses the format string and arguments to print additional information
// alongside the progress bar. The progress parameter should be a float between
// 0 and 1, representing the progress percentage as a decimal. The progress bar
// width is determined by the LoggerConfig.ProgressBarWidth configuration setting.
//
// Related config setting(s):
//   - LoggerConfig.ProgressBarWidth
//
// See also:
//   - Logger.ProgressSuccess
//   - Logger.ProgressError
func (l *Logger) Progress(progress float64, format string, a ...interface{}) {
	a = append([]interface{}{ProgressBar(progress, LoggerConfig.ProgressBarWidth)}, a...)
	l.write('p', "%s "+format, a...)
}

// ProgressSuccess prints a success message with a progress bar followed by the given format and arguments.
// This method works in conjunction with Progress to advance the cursor to the next line
// once the progress has finished.
//
// It uses the format string and arguments to print additional information
// alongside the progress bar. The progress parameter should be a float between
// 0 and 1, representing the progress percentage as a decimal. The progress bar
// width is determined by the LoggerConfig.ProgressBarWidth configuration setting.
//
// Related config setting(s):
//   - LoggerConfig.ProgressBarWidth
//
// See also:
//   - Logger.Progress
//   - Logger.ProgressError
func (l *Logger) ProgressSuccess(progress float64, format string, a ...interface{}) {
	a = append([]interface{}{ClearToEOL(), ProgressBar(progress, LoggerConfig.ProgressBarWidth)}, a...)
	l.Success("%s%s "+format, a...)
}

// ProgressError prints an error message with a progress bar followed by the given format and arguments.
// This method works in conjunction with Progress to advance the cursor to the next line
// once the progress has finished.
//
// It uses the format string and arguments to print additional information
// alongside the progress bar. The progress parameter should be a float between
// 0 and 1, representing the progress percentage as a decimal. The progress bar
// width is determined by the LoggerConfig.ProgressBarWidth configuration setting.
//
// Related config setting(s):
//   - LoggerConfig.ProgressBarWidth
//
// See also:
//   - Logger.Progress
//   - Logger.ProgressSuccess
func (l *Logger) ProgressError(progress float64, format string, a ...interface{}) {
	a = append([]interface{}{ClearToEOL(), ProgressBar(progress, LoggerConfig.ProgressBarWidth)}, a...)
	l.Error("%s%s "+format, a...)
}

// NewLogger creates a new Logger instance with the given ID and settings.
// If `color` is set to `-1`, a color will be chosen automatically based on the ID.
// If `debugMode` is set to `true`, debug level logging will be enabled.
// If `messageHandler` is not `nil`, the logger will write to the provided handler instead of the screen.
func NewLogger(id string, color int, debugMode bool, messageHandler func(string)) *Logger {
	if color == -1 {
		color = stringColorCache.Get(id)
	}
	return &Logger{
		ID:         id,
		color:      color,
		file:       "",
		fileColor:  "",
		debugMode:  debugMode,
		onMessage:  messageHandler,
		traceMode:  false,
		traceLevel: 0,
	}
}

// NewLoggerSimple creates a new Logger instance with the given ID, automatically chosen color, and without debug mode or message handler.
func NewLoggerSimple(id string) *Logger {
	return NewLogger(id, -1, false, nil)
}
