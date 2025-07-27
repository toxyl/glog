package logger

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/toxyl/glog/ansi"
	"github.com/toxyl/glog/colorizers"
	"github.com/toxyl/glog/config"
	"github.com/toxyl/math"
)

const (
	STACK_TRACER_BASE_DEPTH = 3 // this controls how many frames to skip, so we can avoid to expose ourselves in every stack trace
)

type StackTraceFrame struct {
	fn   string
	file string
	line int
}

type Tracer struct {
	frames []*StackTraceFrame
	// depth      int
	// maxDepth   int
	// baseDepth  int
	maxLenFn   int // used to determine the maximum length of a method name in the stack trace
	maxLenFile int // used to determine the maximum length of a file name in the stack trace
	numFrames  int
}

func (st *Tracer) addFrame(frame runtime.Frame) {
	fnName := ""
	if frame.Func != nil {
		fnName = frame.Func.Name()
	} else {
		// if frame.Func is nil we can still get the name from frame.Function
		path := strings.Split(frame.Function, "/")
		fnName = path[len(path)-1]
	}

	if strings.Contains(fnName, "glog.(*GError") {
		return // let's not add ourselves
	}

	stf := &StackTraceFrame{
		fn:   fnName,
		file: frame.File,
		line: frame.Line,
	}

	st.maxLenFile = math.Max(st.maxLenFile, len(stf.file)+2)
	st.maxLenFn = math.Max(st.maxLenFn, len(stf.fn)+2)

	st.frames = append(st.frames, stf)
	st.numFrames++
}

func (st *Tracer) reset() {
	st.maxLenFile = 0
	st.maxLenFn = 0
	st.numFrames = 0
	st.frames = []*StackTraceFrame{}
}

// Sample repopulates the stack trace with the given `maxNumFrames`.
//
// `maxNumFrames` = 0 will sample all frames.
func (st *Tracer) Sample(maxNumFrames int) *Tracer {
	st.reset()

	allFrames := maxNumFrames == 0

	pc := make([]uintptr, STACK_TRACER_BASE_DEPTH+100)
	n := runtime.Callers(STACK_TRACER_BASE_DEPTH, pc)
	if n == 0 {
		return st // avoid processing zero frame in frames.Next() call below
	}

	frames := runtime.CallersFrames(pc[:n])

	for {
		frame, more := frames.Next() // get next frame, more == false indicates the last frame
		st.addFrame(frame)
		if !more || (!allFrames && st.numFrames >= maxNumFrames) {
			break // no more frames to process or we have produced all that were requested
		}
	}

	return st
}

func (st *Tracer) getLines() []string {
	res := []string{}
	const indent = "   "
	const marker = "└──"
	li := len(indent)

	prefix := ""
	maxPadLen := (len(st.frames) - 1) * li
	paddedFnLen := st.maxLenFn + maxPadLen

	for i, stf := range st.frames {
		if i > 1 {
			prefix = indent + prefix
			paddedFnLen -= li
		} else if i == 0 {
			prefix = ""
			paddedFnLen += li
		} else if i == 1 {
			prefix = marker
			paddedFnLen -= li
		}
		pr := strings.Repeat("∙", math.Max(0, paddedFnLen-len(stf.fn)))
		c := config.LoggerConfig.ColorIndicatorDebug - i
		res = append(res,
			fmt.Sprintf(
				"%s%s %s %s:%s",
				ansi.Wrap(prefix, c).String(),
				colorizers.Auto(stf.fn),
				ansi.Wrap(pr, c).String(),
				colorizers.File(stf.file),
				colorizers.Int(stf.line),
			),
		)
	}

	return res
}

func (st *Tracer) PrintWithLogger(l *Logger, indicator rune) {
	lines := st.getLines()
	for _, line := range lines {
		l.write(indicator, "%s", line)
	}
}

func NewTracer() *Tracer {
	st := &Tracer{
		frames:     []*StackTraceFrame{},
		maxLenFn:   0,
		maxLenFile: 0,
		numFrames:  0,
	}

	return st
}

type TraceLine struct {
	Index    int
	Depth    int
	MaxDepth int
	Function string
	Path     string
	Line     int
	Prefix   string
	Logger   *Logger
}

// Related config setting(s):
//
//   - `LoggerConfig.ColorIndicatorDebug`
func (tl *TraceLine) Print(i, l, level, maxLenFunction, maxLenPath, maxLenLine int) {
	d := int(tl.Index - tl.Depth)
	pad := ""
	if d > 0 {
		if d == 1 {
			pad = strings.Repeat(" ", 1)
		} else {
			pad = strings.Repeat(" ", d*4-3)
		}
	}
	c := config.LoggerConfig.ColorIndicatorDebug - level - d

	tl.Logger.write('t', "%s %s %s %s:%s",
		ansi.Wrap(pad+tl.Prefix, c),
		colorizers.Auto(tl.Function),
		ansi.Wrap(strings.Repeat("∙", maxLenFunction+(l-i)*4-len(tl.Function)-6), c),
		colorizers.File(tl.Path),
		colorizers.Int(tl.Line),
	)
}

func NewTraceLine(l *Logger, i, depth, maxDepth int, function, path string, line int) *TraceLine {
	prefix := "└──"
	if i == depth {
		prefix = ""
	}

	return &TraceLine{
		Index:    i,
		Depth:    math.Max(2, depth),
		MaxDepth: math.Max(3, maxDepth),
		Function: function,
		Path:     path,
		Line:     line,
		Prefix:   prefix,
		Logger:   l,
	}
}
