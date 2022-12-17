package glog

import (
	"strings"
)

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

func (tl *TraceLine) Print(i, l, level, maxLenFunction, maxLenPath, maxLenLine int) {
	// if tl.Index == tl.Depth {
	// maxLenFunction += tl.Index * 2
	// }
	d := int(tl.Index - tl.Depth)
	pad := ""
	if d > 0 {
		if d == 1 {
			pad = strings.Repeat(" ", 1)
		} else {
			pad = strings.Repeat(" ", d*4-3)
		}
	}
	c := LoggerConfig.ColorIndicatorDebug + level + d

	tl.Logger.write('t', tl.Logger.prependFormat("%s %s %s %s:%s"),
		Wrap(pad+tl.Prefix, c),
		Auto(tl.Function),
		Wrap(strings.Repeat("∙", maxLenFunction+(l-i)*4-len(tl.Function)-6), c),
		PadLeft(File(tl.Path), maxLenPath, ' '),
		PadRight(Int(tl.Line), maxLenLine, ' '),
	)
}

func NewTraceLine(l *Logger, i, depth, maxDepth int, function, path string, line int) *TraceLine {
	prefix := "└──"
	if i == depth {
		prefix = ""
	}

	return &TraceLine{
		Index:    i,
		Depth:    Max(2, depth),
		MaxDepth: Max(3, maxDepth),
		Function: function,
		Path:     path,
		Line:     line,
		Prefix:   prefix,
		Logger:   l,
	}
}
