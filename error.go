package glog

import (
	"fmt"
	"os"
	"runtime"
	"strings"
)

type GErrorTrace struct {
	depth  int
	line   int
	file   string
	method string
}

type GError struct {
	error
	isFatalError bool
	exitCode     int
	stack        []*GErrorTrace
	maxFuncLen   int
	maxPathLen   int
	maxLineLen   int
}

func (ge *GError) trace() {
	res := []*GErrorTrace{}
	depth := 2
	maxDepth := depth + 22
	pcCaller1 := make([]uintptr, maxDepth+depth)
	nCaller1 := runtime.Callers(int(depth), pcCaller1)
	ml := int(Min(len(pcCaller1), Min(nCaller1, int(maxDepth))))
	ge.maxFuncLen, ge.maxPathLen, ge.maxLineLen = 0, 0, 0

	for i := int(depth); i < ml; i++ {
		frames := runtime.CallersFrames(pcCaller1[i-2 : i-1])
		frameCaller, _ := frames.Next()
		fullPathFuncCaller := strings.Split(frameCaller.Function, "/")
		fnName := fullPathFuncCaller[len(fullPathFuncCaller)-1]
		fnFile := frameCaller.File
		fnLine := frameCaller.Line
		if fnName == "glog.(*GErrorRegistry).Check" {
			continue // skip checks from the error registry
		}
		res = append(res, &GErrorTrace{
			line:   fnLine,
			file:   fnFile,
			method: fnName,
			depth:  i - depth,
		})

		ge.maxFuncLen = Max(ge.maxFuncLen, len(fnName))
		ge.maxPathLen = Max(ge.maxPathLen, len(fnFile))
		ge.maxLineLen = Max(ge.maxLineLen, len(fmt.Sprint(fnLine)))
	}

	ge.stack = res
}

//
// Related config setting(s):
//
//  - `LoggerConfig.ColorIndicatorDebug`
func (ge *GError) printStackTrace(maxLevel int, logger *Logger) {
	for i, g := range ge.stack {
		if i == 0 {
			continue
		}
		if i >= maxLevel {
			break
		}
		d := g.depth
		pad := ""
		if d > 0 {
			if d == 1 {
				pad = strings.Repeat(" ", 1)
			} else {
				pad = strings.Repeat(" ", d*4-3)
			}
		}
		c := LoggerConfig.ColorIndicatorDebug - d
		prefix := "└──"
		if d == 0 {
			prefix = ""
		}
		prefix = pad + prefix
		l := Max(0, (len(prefix)+(ge.maxFuncLen-len(g.method))+2)-((d-1)*2)*4)

		logger.write('x', logger.prependFormat("%s %s %s %s:%s"),
			Wrap(prefix, c),
			Auto(g.method),
			Wrap(strings.Repeat("∙", l), c),
			File(g.file),
			Int(g.line),
		)
	}
}

func (ge *GError) check(err error, msg string, logger *Logger) bool {
	if err == nil {
		return false
	}
	ge.trace()
	es := err.Error()
	ges := ge.Error()
	if err == ge.error || es == ges || strings.HasSuffix(es, ": "+ges) || strings.HasPrefix(es, ges+": ") || strings.Contains(es, ": "+ges+": ") {
		logger.Error("%s: %s", msg, Error(err))
		if ge.isFatalError {
			ge.printStackTrace(22, logger)
			os.Exit(ge.exitCode)
		} else {
			ge.printStackTrace(3, logger)
			return true
		}
	}
	return false
}

func NewGError(err error, isFatal bool, exitCode int) *GError {
	ge := &GError{
		error:        err,
		stack:        []*GErrorTrace{},
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
	for _, ge := range errors {
		ger.registeredErrors = append(ger.registeredErrors, ge)
	}
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
