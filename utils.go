package glog

import (
	"github.com/toxyl/glog/ansi"
	"github.com/toxyl/glog/logger"
	"github.com/toxyl/glog/types"
	"github.com/toxyl/glog/utils"
)

// Backwards compatibility aliases and / or equivalent replacements

type Ints = types.Ints
type IntOrInterface = types.IntOrInterface
type Uints = types.Uints
type UintOrInterface = types.UintOrInterface
type IntOrUint = types.IntOrUint
type Floats = types.Floats
type FloatOrInterface = types.FloatOrInterface
type Number = types.Number
type NumberOrInterface = types.NumberOrInterface
type Durations = types.Durations
type StringColorCache = utils.StringColorCache
type StackTraceFrame = logger.StackTraceFrame
type StackTracer = logger.Tracer
type DurationScale = utils.DurationScale
type PathType = utils.PathType

const (
	INVALID_PATH            = utils.INVALID_PATH
	FILE_PATH               = utils.FILE_PATH
	URL_PATH                = utils.URL_PATH
	STACK_TRACER_BASE_DEPTH = logger.STACK_TRACER_BASE_DEPTH // this controls how many frames to skip, so we can avoid to expose ourselves in every stack trace

	SECOND        = utils.SECOND
	MINUTE        = utils.MINUTE
	HOUR          = utils.HOUR
	DAY           = utils.DAY
	WEEK          = utils.WEEK
	YEAR_COMMON   = utils.YEAR_COMMON
	YEAR_LEAP     = utils.YEAR_LEAP
	YEAR_AVERAGE  = utils.YEAR_AVERAGE
	MONTH_COMMON  = utils.MONTH_COMMON
	MONTH_LEAP    = utils.MONTH_LEAP
	MONTH_AVERAGE = utils.MONTH_AVERAGE

	DURATION_SCALE_AVERAGE = utils.DURATION_SCALE_AVERAGE
	DURATION_SCALE_COMMON  = utils.DURATION_SCALE_COMMON
	DURATION_SCALE_LEAP    = utils.DURATION_SCALE_LEAP
)

var (
	// Math
	GetFloat     = utils.GetFloat
	GetRandomInt = utils.GetRandomInt

	// Net
	SplitHostPort = utils.SplitHostPort
	ReverseDNS    = utils.ReverseDNS

	// Paths
	IsURL       = utils.IsURL
	IsFile      = utils.IsFile
	IsValidPath = utils.IsValidPath

	// Plaintext
	StripANSI          = utils.StripANSI
	ReplaceRunes       = utils.ReplaceRunes
	ReplaceEmojis      = utils.ReplaceEmojis
	RemoveNonPrintable = utils.RemoveNonPrintable

	PlaintextString       = utils.PlaintextString
	PlaintextStringLength = utils.PlaintextStringLength

	// Progress
	ProgressBar = logger.ProgressBar

	// StackTracer
	NewStackTracer = logger.NewTracer

	// Time
	RandomSleep = utils.RandomSleep

	// Text formatting
	Reset     = ansi.Reset().String
	Bold      = ansi.Bold().String
	Dim       = ansi.Dim().String
	Italic    = ansi.Italic().String
	Underline = ansi.Underline().String

	DoubleUnderline = ansi.DoubleUnderline().String
	Blink           = ansi.Blink().String
	RapidBlink      = ansi.RapidBlink().String
	Reverse         = ansi.Reverse().String
	Conceal         = ansi.Conceal().String
	StrikeThrough   = ansi.StrikeThrough().String

	// Cursor movement
	CursorToLineStart         = ansi.CursorToLineStart().String
	CursorToLineStartAndClear = ansi.CursorToLineStartAndClear().String
	StoreCursor               = ansi.StoreCursor().String
	RestoreCursor             = ansi.RestoreCursor().String
	HideCursor                = ansi.HideCursor().String
	ShowCursor                = ansi.ShowCursor().String
	// Screen control
	ClearScreen           = ansi.ClearScreen().String
	ClearScreenFromCursor = ansi.ClearScreenFromCursor().String
	ClearScreenToCursor   = ansi.ClearScreenToCursor().String
	ClearLine             = ansi.ClearLine().String
	ClearToEOL            = ansi.ClearToEOL().String
	ClearLineToCursor     = ansi.ClearLineToCursor().String
	EnableLineWrap        = ansi.EnableLineWrap().String
	DisableLineWrap       = ansi.DisableLineWrap().String
)

// Colors
func Color(color int) string                      { return ansi.Color(color).String() }
func Wrap(str string, color int) string           { return ansi.Wrap(str, color).String() }
func BackgroundColor(color int) string            { return ansi.BackgroundColor(color).String() }
func WrapBackground(str string, color int) string { return ansi.WrapBackground(str, color).String() }

// Cursor movement
func CursorUp(n int) string              { return ansi.CursorUp(n).String() }
func CursorDown(n int) string            { return ansi.CursorDown(n).String() }
func CursorForward(n int) string         { return ansi.CursorForward(n).String() }
func CursorBackward(n int) string        { return ansi.CursorBackward(n).String() }
func CursorPosition(row, col int) string { return ansi.CursorPosition(row, col).String() }

// Screen control
func ScrollUp(n int) string   { return ansi.ScrollUp(n).String() }
func ScrollDown(n int) string { return ansi.ScrollDown(n).String() }

// Math
func Max[N Number](a, b N) N { return utils.Max(a, b) }
func Min[N Number](a, b N) N { return utils.Min(a, b) }

// Padding

func PadLeft[I IntOrUint](str string, maxLength I, char rune) string {
	return utils.PadLeft(str, maxLength, char)
}

func PadRight[I IntOrUint](str string, maxLength I, char rune) string {
	return utils.PadRight(str, maxLength, char)
}

func PadCenter[I IntOrUint](str string, maxLength I, char rune) string {
	return utils.PadCenter(str, maxLength, char)
}
