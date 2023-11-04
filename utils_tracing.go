package glog

import (
	"fmt"
	"runtime"
	"strings"
)

const (
	STACK_TRACER_BASE_DEPTH = 3 // this controls how many frames to skip, so we can avoid to expose ourselves in every stack trace
)

type StackTraceFrame struct {
	fn   string
	file string
	line int
}

type StackTracer struct {
	frames []*StackTraceFrame
	// depth      int
	// maxDepth   int
	// baseDepth  int
	maxLenFn   int // used to determine the maximum length of a method name in the stack trace
	maxLenFile int // used to determine the maximum length of a file name in the stack trace
	numFrames  int
}

func (st *StackTracer) addFrame(frame runtime.Frame) {
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

	st.maxLenFile = Max(st.maxLenFile, len(stf.file)+2)
	st.maxLenFn = Max(st.maxLenFn, len(stf.fn)+2)

	st.frames = append(st.frames, stf)
	st.numFrames++
}

func (st *StackTracer) reset() {
	st.maxLenFile = 0
	st.maxLenFn = 0
	st.numFrames = 0
	st.frames = []*StackTraceFrame{}
}

// Sample repopulates the stack trace with the given `maxNumFrames`.
//
// `maxNumFrames` = 0 will sample all frames.
func (st *StackTracer) Sample(maxNumFrames int) *StackTracer {
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

func (st *StackTracer) getLines() []string {
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
		pr := strings.Repeat("∙", Max(0, paddedFnLen-len(stf.fn)))
		c := LoggerConfig.ColorIndicatorDebug - i
		res = append(res,
			fmt.Sprintf(
				"%s%s %s %s:%s",
				Wrap(prefix, c),
				Auto(stf.fn),
				Wrap(pr, c),
				File(stf.file),
				Int(stf.line),
			),
		)
	}

	return res
}

func (st *StackTracer) PrintWithLogger(l *Logger, indicator rune) {
	lines := st.getLines()
	for _, line := range lines {
		l.write(indicator, "%s", line)
	}
}

func NewStackTracer() *StackTracer {
	st := &StackTracer{
		frames:     []*StackTraceFrame{},
		maxLenFn:   0,
		maxLenFile: 0,
		numFrames:  0,
	}

	return st
}
