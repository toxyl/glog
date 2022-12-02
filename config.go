package glog

import "time"

type Config struct {
	TimeFormat,
	TimeFormat12hr,
	DateFormat,
	DateTimeFormat,
	DateTimeFormat12hr string
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
	ColorPassword,
	ColorReason,
	ColorFile,
	ColorError,
	ColorIndicator,
	ColorIndicatorInfo,
	ColorIndicatorOK,
	ColorIndicatorSuccess,
	ColorIndicatorNotOK,
	ColorIndicatorError,
	ColorIndicatorWarning,
	ColorIndicatorDebug int
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
		TimeFormat:              "15:04:05",
		TimeFormat12hr:          "03:04:05pm",
		DateFormat:              "2006-01-02",
		DateTimeFormat:          "2006-01-02 15:04:05",
		DateTimeFormat12hr:      "2006-01-02 03:04:05pm",
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
		ColorTime:               Purple,
		ColorDuration:           Purple - 50,
		ColorPassword:           DarkRed,
		ColorReason:             Orange,
		ColorFile:               LightBlue,
		ColorError:              Red,
		ColorIndicator:          Gray,
		ColorIndicatorInfo:      LightBlue,
		ColorIndicatorOK:        OliveGreen,
		ColorIndicatorSuccess:   Green,
		ColorIndicatorNotOK:     DarkRed,
		ColorIndicatorError:     Red,
		ColorIndicatorWarning:   Yellow,
		ColorIndicatorDebug:     Orange,
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

	return c
}

var LoggerConfig *Config = NewDefaultConfig()
