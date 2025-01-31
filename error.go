package glog

import (
	"os"
	"strings"
)

var stackTracer = NewStackTracer()

type GError struct {
	error
	isFatalError bool
	exitCode     int
}

// Related config setting(s):
//
//   - `LoggerConfig.ColorIndicatorDebug`
func (ge *GError) printStackTrace(maxLevel int, logger *Logger) {
	stackTracer.Sample(maxLevel).PrintWithLogger(logger, 'x')
}

func (ge *GError) check(err error, msg string, logger *Logger) bool {
	if err == nil {
		return false
	}
	es := err.Error()
	ges := ge.Error()
	if err == ge.error || es == ges || strings.HasSuffix(es, ": "+ges) || strings.HasPrefix(es, ges+": ") || strings.Contains(es, ": "+ges+": ") {
		logger.Error("%s: %s", msg, Error(err))
		if ge.isFatalError {
			ge.printStackTrace(0, logger)
			logger.Error("%s (exit status %s: %s)", Bold()+WrapRed("FATAL"), Int(ge.exitCode), err.Error())
			os.Exit(ge.exitCode)
		} else {
			ge.printStackTrace(1, logger)
			return true
		}
	}
	return false
}

func NewGError(err error, isFatal bool, exitCode int) *GError {
	ge := &GError{
		error:        err,
		isFatalError: isFatal,
		exitCode:     exitCode,
	}
	return ge
}

type GErrorRegistry struct {
	registeredErrors []*GError
	logger           *Logger
}

func (ger *GErrorRegistry) Append(errors ...*GError) *GErrorRegistry {
	ger.registeredErrors = append(ger.registeredErrors, errors...)
	return ger
}

func (ger *GErrorRegistry) Register(err error, isFatal bool, exitCode int) *GErrorRegistry {
	return ger.Append(NewGError(err, isFatal, exitCode))
}

func (ger *GErrorRegistry) Check(err error, msg string) {
	if err == nil {
		return
	}
	for _, ge := range ger.registeredErrors {
		if ge.check(err, msg, ger.logger) {
			break
		}
	}
}

func NewGErrorRegistry() *GErrorRegistry {
	ger := &GErrorRegistry{
		registeredErrors: []*GError{},
		logger:           NewLoggerSimple("Errors"),
	}
	return ger
}
