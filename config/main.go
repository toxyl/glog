package config

import (
	"time"

	"github.com/toxyl/glog/colormap"
	"github.com/toxyl/glog/indicator"
)

type Config struct {
	TablePadChar       rune
	AutoFloatPrecision int
	TimeFormat,
	TimeFormat12hr,
	DateFormat,
	DateTimeFormat,
	DateTimeFormat12hr string
	ColorURLSeparators,
	ColorPathSeparator,
	ColorPath,
	ColorScheme,
	ColorUser,
	ColorPassword,
	ColorURLPath,
	ColorQueryKey,
	ColorQueryValue,
	ColorFragment,
	ColorNil,
	ColorIntNegative,
	ColorIntZero,
	ColorIntPositive,
	ColorUintZero,
	ColorUintPositive,
	ColorFloatNegative,
	ColorFloatZero,
	ColorFloatPositive,
	ColorPercentageNegative,
	ColorPercentageZero,
	ColorPercentagePositive,
	ColorBoolFalse,
	ColorBoolTrue,
	ColorTime,
	ColorDuration,
	ColorReason,
	ColorFile,
	ColorError,
	ColorUnitHumanReadable,
	ColorIndicator,
	ColorIndicatorInfo,
	ColorIndicatorOK,
	ColorIndicatorSuccess,
	ColorIndicatorNotOK,
	ColorIndicatorError,
	ColorIndicatorWarning,
	ColorIndicatorDebug,
	ColorIndicatorTrace,
	ColorIndicatorProgress,
	ColorIndicatorQuestion int
	ColorsDisabled,
	ShowRuntimeHumanReadable,
	ShowRuntimeSeconds,
	ShowRuntimeMilliseconds,
	ShowDateTime,
	ShowSubsystem,
	ShowIndicator,
	SplitOnNewLine,
	CheckIfURLIsAlive bool
	ProgressBarWidth int
	Indicators       map[rune]*indicator.Indicator
	ReverseDNSCache  map[string]string
	CreatedAt        time.Time
}

func (c *Config) AddIndicator(id rune, value string, color int) {
	c.Indicators[id] = indicator.NewIndicator(value, color)
}

func NewDefaultConfig() *Config {
	c := &Config{
		TablePadChar:             ' ',
		AutoFloatPrecision:       2,
		TimeFormat:               "15:04:05",
		TimeFormat12hr:           "03:04:05pm",
		DateFormat:               "2006-01-02",
		DateTimeFormat:           "2006-01-02 15:04:05",
		DateTimeFormat12hr:       "2006-01-02 03:04:05pm",
		ColorPathSeparator:       207,
		ColorURLSeparators:       colormap.BrightYellow,
		ColorScheme:              colormap.DarkYellow,
		ColorUser:                colormap.DarkGreen,
		ColorPassword:            colormap.DarkRed,
		ColorURLPath:             128,
		ColorPath:                148,
		ColorQueryKey:            colormap.Orange,
		ColorQueryValue:          colormap.DarkOrange,
		ColorFragment:            colormap.LightBlue,
		ColorNil:                 colormap.DarkRed,
		ColorIntNegative:         colormap.Red,
		ColorIntZero:             colormap.Blue,
		ColorIntPositive:         colormap.Cyan,
		ColorUintZero:            colormap.Blue,
		ColorUintPositive:        colormap.Cyan,
		ColorFloatNegative:       colormap.Red,
		ColorFloatZero:           colormap.Blue,
		ColorFloatPositive:       colormap.Cyan,
		ColorPercentageNegative:  colormap.Red,
		ColorPercentageZero:      colormap.Blue - 2,
		ColorPercentagePositive:  colormap.Cyan - 2,
		ColorBoolFalse:           colormap.Red,
		ColorBoolTrue:            colormap.Green,
		ColorTime:                26,
		ColorDuration:            208,
		ColorReason:              colormap.Orange,
		ColorFile:                colormap.LightBlue,
		ColorError:               colormap.Red,
		ColorUnitHumanReadable:   160,
		ColorIndicator:           colormap.DarkGray,
		ColorIndicatorInfo:       colormap.LightBlue,
		ColorIndicatorOK:         colormap.OliveGreen,
		ColorIndicatorSuccess:    colormap.Green,
		ColorIndicatorNotOK:      colormap.DarkRed,
		ColorIndicatorError:      colormap.Red,
		ColorIndicatorWarning:    colormap.Yellow,
		ColorIndicatorDebug:      colormap.Orange,
		ColorIndicatorTrace:      colormap.Orange,
		ColorIndicatorProgress:   colormap.LightBlue,
		ColorIndicatorQuestion:   colormap.Lime,
		ColorsDisabled:           false,
		ShowRuntimeHumanReadable: false,
		ShowRuntimeSeconds:       false,
		ShowRuntimeMilliseconds:  true,
		ShowDateTime:             true,
		ShowSubsystem:            true,
		ShowIndicator:            true,
		SplitOnNewLine:           false, // false by default to not break old behavior
		CheckIfURLIsAlive:        true,  // true by default to not break old behavior
		ProgressBarWidth:         20,
		Indicators:               map[rune]*indicator.Indicator{},
		ReverseDNSCache:          map[string]string{},
		CreatedAt:                time.Now(),
	}

	c.AddIndicator('✓', "[✓]", c.ColorIndicatorSuccess)
	c.AddIndicator('+', "[+]", c.ColorIndicatorOK)
	c.AddIndicator('i', "[i]", c.ColorIndicatorInfo)
	c.AddIndicator(' ', "[ ]", c.ColorIndicator)
	c.AddIndicator('_', "   ", c.ColorIndicator)
	c.AddIndicator('!', "[!]", c.ColorIndicatorWarning)
	c.AddIndicator('-', "[-]", c.ColorIndicatorNotOK)
	c.AddIndicator('x', "[x]", c.ColorIndicatorError)
	c.AddIndicator('d', "[D]", c.ColorIndicatorDebug)
	c.AddIndicator('?', "[?]", c.ColorIndicatorQuestion)
	c.AddIndicator('t', "[T]", c.ColorIndicatorTrace)
	c.AddIndicator('p', "[∞]", c.ColorIndicatorProgress)

	return c
}

var LoggerConfig *Config = NewDefaultConfig()
