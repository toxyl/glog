package glog

import "time"

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
	ColorIndicatorQuestion int
	ColorsDisabled,
	ShowRuntimeSeconds,
	ShowRuntimeMilliseconds,
	ShowDateTime,
	ShowSubsystem,
	ShowIndicator bool
	Indicators      map[rune]*Indicator
	reverseDNSCache map[string]string
	createdAt       time.Time
}

func (c *Config) AddIndicator(indicator rune, value string, color int) {
	c.Indicators[indicator] = NewIndicator(value, color)
}

func NewDefaultConfig() *Config {
	c := &Config{
		TablePadChar:            ' ',
		AutoFloatPrecision:      2,
		TimeFormat:              "15:04:05",
		TimeFormat12hr:          "03:04:05pm",
		DateFormat:              "2006-01-02",
		DateTimeFormat:          "2006-01-02 15:04:05",
		DateTimeFormat12hr:      "2006-01-02 03:04:05pm",
		ColorPathSeparator:      207,
		ColorURLSeparators:      BrightYellow,
		ColorScheme:             DarkYellow,
		ColorUser:               DarkGreen,
		ColorPassword:           DarkRed,
		ColorURLPath:            128,
		ColorPath:               148,
		ColorQueryKey:           Orange,
		ColorQueryValue:         DarkOrange,
		ColorFragment:           LightBlue,
		ColorNil:                DarkRed,
		ColorIntNegative:        Red,
		ColorIntZero:            Blue,
		ColorIntPositive:        Cyan,
		ColorUintZero:           Blue,
		ColorUintPositive:       Cyan,
		ColorFloatNegative:      Red,
		ColorFloatZero:          Blue,
		ColorFloatPositive:      Cyan,
		ColorPercentageNegative: Red,
		ColorPercentageZero:     Blue - 2,
		ColorPercentagePositive: Cyan - 2,
		ColorBoolFalse:          Red,
		ColorBoolTrue:           Green,
		ColorTime:               26,
		ColorDuration:           208,
		ColorReason:             Orange,
		ColorFile:               LightBlue,
		ColorError:              Red,
		ColorUnitHumanReadable:  160,
		ColorIndicator:          Gray,
		ColorIndicatorInfo:      LightBlue,
		ColorIndicatorOK:        OliveGreen,
		ColorIndicatorSuccess:   Green,
		ColorIndicatorNotOK:     DarkRed,
		ColorIndicatorError:     Red,
		ColorIndicatorWarning:   Yellow,
		ColorIndicatorDebug:     Orange,
		ColorIndicatorTrace:     Orange,
		ColorIndicatorQuestion:  Lime,
		ColorsDisabled:          false,
		ShowRuntimeSeconds:      false,
		ShowRuntimeMilliseconds: true,
		ShowDateTime:            true,
		ShowSubsystem:           true,
		ShowIndicator:           true,
		Indicators:              map[rune]*Indicator{},
		reverseDNSCache:         map[string]string{},
		createdAt:               time.Now(),
	}

	c.AddIndicator('✓', "[✓]", c.ColorIndicatorSuccess)
	c.AddIndicator('+', "[+]", c.ColorIndicatorOK)
	c.AddIndicator('i', "[i]", c.ColorIndicatorInfo)
	c.AddIndicator(' ', "[ ]", c.ColorIndicator)
	c.AddIndicator('!', "[!]", c.ColorIndicatorWarning)
	c.AddIndicator('-', "[-]", c.ColorIndicatorNotOK)
	c.AddIndicator('x', "[x]", c.ColorIndicatorError)
	c.AddIndicator('d', "[D]", c.ColorIndicatorDebug)
	c.AddIndicator('?', "[?]", c.ColorIndicatorQuestion)
	c.AddIndicator('t', "[T]", c.ColorIndicatorTrace)

	return c
}

var LoggerConfig *Config = NewDefaultConfig()
