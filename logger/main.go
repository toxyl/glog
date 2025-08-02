package logger

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/toxyl/glog/ansi"
	"github.com/toxyl/glog/colorizers"
	"github.com/toxyl/glog/config"
	"github.com/toxyl/glog/utils"
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
func (l *Logger) write(indicator rune, format string, a ...any) {
	prefix := ""

	if config.LoggerConfig.ShowIndicator {
		if vi, ok := config.LoggerConfig.Indicators[indicator]; ok {
			prefix = ansi.Wrap(vi.Value, vi.Color).String()
		}
	}

	if config.LoggerConfig.ShowRuntimeHumanReadable {
		prefix = fmt.Sprintf("%22s %s", colorizers.RuntimeHumanReadable(), prefix)
	}
	if config.LoggerConfig.ShowRuntimeSeconds {
		prefix = fmt.Sprintf("%22s s %s", colorizers.RuntimeSeconds(), prefix)
	}
	if config.LoggerConfig.ShowRuntimeMilliseconds {
		prefix = fmt.Sprintf("%22s ms %s", colorizers.RuntimeMilliseconds(), prefix)
	}
	if config.LoggerConfig.ShowDateTime {
		prefix = fmt.Sprintf("%s %s", colorizers.DateTime(time.Now()), prefix)
	}
	if config.LoggerConfig.ShowSubsystem {
		prefix = fmt.Sprintf("%s %s: ", prefix, ansi.Wrap(fmt.Sprintf("%-16s", l.ID), l.color).String())
	}

	msg := fmt.Sprintf(format, a...)
	if config.LoggerConfig.SplitOnNewLine {
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
		msg = ansi.StoreCursor().String() + ansi.ClearToEOL().String() + msg + ansi.RestoreCursor().String()
	} else {
		msg += "\n"
	}

	if l.onMessage != nil {
		l.onMessage(msg)
		return
	}

	if config.LoggerConfig.ColorsDisabled {
		msg = utils.StripANSI(msg)
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

		if _, err = f.WriteString(utils.StripANSI(msg)); err != nil {
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

// auto prints a message using the given indicator, but will first run all arguments
// through glog.Auto()
func (l *Logger) auto(indicator rune, format string, a ...any) {
	str := []any{}
	for _, s := range a {
		str = append(str, colorizers.Auto(s))
	}
	l.write(indicator, format, str...)
}

// Blank prints a message without any indicator such as "[ ]", "[i]", etc.
func (l *Logger) Blank(format string, a ...any) {
	l.write('_', format, a...)
}

// BlankAuto does the same as Blank but will process all arguments with glog.Auto(...) first.
func (l *Logger) BlankAuto(format string, a ...any) {
	l.auto('_', format, a...)
}

func (l *Logger) Default(format string, a ...any) {
	l.write(' ', format, a...)
}

// DefaultAuto does the same as Default but will process all arguments with glog.Auto(...) first.
func (l *Logger) DefaultAuto(format string, a ...any) {
	l.auto(' ', format, a...)
}

func (l *Logger) Info(format string, a ...any) {
	l.write('i', format, a...)
}

// InfoAuto does the same as Info but will process all arguments with glog.Auto(...) first.
func (l *Logger) InfoAuto(format string, a ...any) {
	l.auto('i', format, a...)
}

func (l *Logger) Success(format string, a ...any) {
	l.write('✓', format, a...)
}

// SuccessAuto does the same as Success but will process all arguments with glog.Auto(...) first.
func (l *Logger) SuccessAuto(format string, a ...any) {
	l.auto('✓', format, a...)
}

func (l *Logger) OK(format string, a ...any) {
	l.write('+', format, a...)
}

// OKAuto does the same as OK but will process all arguments with glog.Auto(...) first.
func (l *Logger) OKAuto(format string, a ...any) {
	l.auto('+', format, a...)
}

func (l *Logger) NotOK(format string, a ...any) {
	l.write('-', format, a...)
}

// NotOKAuto does the same as NotOK but will process all arguments with glog.Auto(...) first.
func (l *Logger) NotOKAuto(format string, a ...any) {
	l.auto('-', format, a...)
}

func (l *Logger) Error(format string, a ...any) {
	l.write('x', format, a...)
}

// ErrorAuto does the same as Error but will process all arguments with glog.Auto(...) first.
func (l *Logger) ErrorAuto(format string, a ...any) {
	l.auto('x', format, a...)
}

func (l *Logger) Warning(format string, a ...any) {
	l.write('!', format, a...)
}

// WarningAuto does the same as Warning but will process all arguments with glog.Auto(...) first.
func (l *Logger) WarningAuto(format string, a ...any) {
	l.auto('!', format, a...)
}

func (l *Logger) Debug(format string, a ...any) {
	if !l.debugMode {
		return
	}
	l.write('d', format, a...)
}

// DebugAuto does the same as Debug but will process all arguments with glog.Auto(...) first.
func (l *Logger) DebugAuto(format string, a ...any) {
	if !l.debugMode {
		return
	}
	l.auto('d', format, a...)
}

func (l *Logger) Question(format string, a ...any) {
	l.write('?', format, a...)
}

// QuestionAuto does the same as Question but will process all arguments with glog.Auto(...) first.
func (l *Logger) QuestionAuto(format string, a ...any) {
	l.auto('?', format, a...)
}

func (l *Logger) Trace(level uint) {
	if !l.traceMode || level > l.traceLevel {
		return
	}
	StackTracer.Sample(0).PrintWithLogger(l, 't')
}

func (l *Logger) ShowColors() {
	str := ""
	for i := 0; i < 8; i++ {
		str += ansi.Wrap(fmt.Sprintf("%03d ", i), i).String()
	}
	l.Default("%s", str)

	str = ""
	for i := 8; i < 16; i++ {
		str += ansi.Wrap(fmt.Sprintf("%03d ", i), i).String()
	}
	l.Default("%s", str)

	for i := 16; i < 256; i++ {
		str := ""
		for j := 0; j < 12 && i < 256; j++ {
			str += ansi.Wrap(fmt.Sprintf("%03d ", i), i).String()
			i++
		}
		i--
		l.Default("%s", str)
	}
}

func (l *Logger) Table(columns ...*TableColumn) {
	NewTable(columns...).Print(l)
}

func (l *Logger) TableWithoutHeader(columns ...*TableColumn) {
	NewTable(columns...).PrintWithoutHeader(l)
}

func (l *Logger) KeyValueTable(data map[string]any) {
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
func (l *Logger) Progress(progress float64, format string, a ...any) {
	a = append([]any{ProgressBar(progress, config.LoggerConfig.ProgressBarWidth)}, a...)
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
func (l *Logger) ProgressSuccess(progress float64, format string, a ...any) {
	a = append([]any{ansi.ClearToEOL().String(), ProgressBar(progress, config.LoggerConfig.ProgressBarWidth)}, a...)
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
func (l *Logger) ProgressError(progress float64, format string, a ...any) {
	a = append([]any{ansi.ClearToEOL().String(), ProgressBar(progress, config.LoggerConfig.ProgressBarWidth)}, a...)
	l.Error("%s%s "+format, a...)
}

// QuestionInline prints a question message without adding a newline, allowing for inline user input.
// This method uses the same visual styling as Question but doesn't advance to the next line.
func (l *Logger) QuestionInline(format string, a ...any) {
	prefix := ""

	if config.LoggerConfig.ShowIndicator {
		if vi, ok := config.LoggerConfig.Indicators['?']; ok {
			prefix = ansi.Wrap(vi.Value, vi.Color).String()
		}
	}

	if config.LoggerConfig.ShowRuntimeHumanReadable {
		prefix = fmt.Sprintf("%22s %s", colorizers.RuntimeHumanReadable(), prefix)
	}
	if config.LoggerConfig.ShowRuntimeSeconds {
		prefix = fmt.Sprintf("%22s s %s", colorizers.RuntimeSeconds(), prefix)
	}
	if config.LoggerConfig.ShowRuntimeMilliseconds {
		prefix = fmt.Sprintf("%22s ms %s", colorizers.RuntimeMilliseconds(), prefix)
	}
	if config.LoggerConfig.ShowDateTime {
		prefix = fmt.Sprintf("%s %s", colorizers.DateTime(time.Now()), prefix)
	}
	if config.LoggerConfig.ShowSubsystem {
		prefix = fmt.Sprintf("%s %s: ", prefix, ansi.Wrap(fmt.Sprintf("%-16s", l.ID), l.color).String())
	}

	msg := fmt.Sprintf(format, a...)
	if config.LoggerConfig.SplitOnNewLine {
		res := []string{}
		for ln := range strings.SplitSeq(msg, "\n") {
			res = append(res, prefix+" "+ln)
		}
		msg = strings.Join(res, "\n")
	} else {
		msg = prefix + " " + msg
	}

	// No newline added for inline questions

	if l.onMessage != nil {
		l.onMessage(msg)
		return
	}

	if config.LoggerConfig.ColorsDisabled {
		msg = utils.StripANSI(msg)
	}

	if l.file != "" {
		err := os.MkdirAll(filepath.Dir(l.file), 0770)
		if err != nil {
			panic(err)
		}

		f, err := os.OpenFile(l.file, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		if err != nil {
			panic(err)
		}

		defer f.Close()

		if _, err = f.WriteString(utils.StripANSI(msg)); err != nil {
			panic(err)
		}
	}

	if l.fileColor != "" {
		err := os.MkdirAll(filepath.Dir(l.fileColor), 0770)
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

// QuestionInlineAuto does the same as QuestionInline but will process all arguments with glog.Auto(...) first.
func (l *Logger) QuestionInlineAuto(format string, a ...any) {
	str := []any{}
	for _, s := range a {
		str = append(str, colorizers.Auto(s))
	}
	l.QuestionInline(format, str...)
}

func (l *Logger) Ask(question, optionsOrType, defaultAnswer string) string {
	l.QuestionInline("%s [%s] <default=%s>: ", question, optionsOrType, defaultAnswer)

	var userInput string
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		userInput = scanner.Text()
	}
	if len(userInput) == 0 {
		return defaultAnswer
	}
	return userInput
}

func (l *Logger) AskBool(question string, defaultAnswer bool) bool {
	options := "Y|n"
	answer := "y"
	if !defaultAnswer {
		options = "y|N"
		answer = "n"
	}
	answer = l.Ask(question, options, answer)
	return strings.ToLower(answer) == "y"
}

func (l *Logger) AskList(question string) []string {
	l.Question("%s [%s]: ", question, "list, press ENTER to end")
	res := []string{}
	i := 0
	for {
		l.QuestionInline("[%s]: ", colorizers.Auto(i))
		var userInput string
		fmt.Scanln(&userInput)
		if userInput == "" {
			break
		}
		res = append(res, userInput)
		i++
	}
	return res
}

// AskWithTimeout prompts the user with the given question and waits for a response,
// timing out after the specified duration. If no response is received, it returns
// the default answer.
func (l *Logger) AskWithTimeout(question, options, defaultAnswer string, timeout time.Duration) string {
	l.QuestionInline("%s [%s]: ", question, options)

	ch := make(chan string)

	go func() {
		var userInput string
		fmt.Scanln(&userInput)
		ch <- strings.ToLower(userInput)
	}()

	select {
	case userInput := <-ch:
		if userInput == "" {
			return defaultAnswer
		}
		return userInput
	case <-time.After(timeout):
		fmt.Println()
		l.Info("Timeout reached. Default answer (%s) selected.", defaultAnswer)
		return defaultAnswer
	}
}

func (l *Logger) AskBoolWithTimeout(question string, defaultAnswer bool, timeout time.Duration) bool {
	options := "Y|n"
	answer := "y"
	if !defaultAnswer {
		options = "y|N"
		answer = "n"
	}
	answer = l.AskWithTimeout(question, options, answer, timeout)
	return strings.ToLower(answer) == "y"
}

// NewLogger creates a new Logger instance with the given ID and settings.
// If `color` is set to `-1`, a color will be chosen automatically based on the ID.
// If `debugMode` is set to `true`, debug level logging will be enabled.
// If `messageHandler` is not `nil`, the logger will write to the provided handler instead of the screen.
func NewLogger(id string, color int, debugMode bool, messageHandler func(string)) *Logger {
	if color == -1 {
		color = utils.Scc.Get(id)
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
