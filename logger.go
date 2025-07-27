package glog

import (
	"github.com/toxyl/glog/colormap"
	"github.com/toxyl/glog/config"
	"github.com/toxyl/glog/indicator"
	"github.com/toxyl/glog/logger"
)

type Logger = logger.Logger
type Config = config.Config
type Indicator = indicator.Indicator
type TraceLine = logger.TraceLine
type TableColumn = logger.TableColumn
type Table = logger.Table
type GError = logger.GError
type GErrorRegistry = logger.GErrorRegistry

const (
	PAD_LEFT     = logger.PAD_LEFT
	PAD_CENTER   = logger.PAD_CENTER
	PAD_RIGHT    = logger.PAD_RIGHT
	DarkBlue     = colormap.DarkBlue
	Blue         = colormap.Blue
	DarkGreen    = colormap.DarkGreen
	LightBlue    = colormap.LightBlue
	OliveGreen   = colormap.OliveGreen
	Green        = colormap.Green
	Cyan         = colormap.Cyan
	Purple       = colormap.Purple
	DarkOrange   = colormap.DarkOrange
	DarkYellow   = colormap.DarkYellow
	Lime         = colormap.Lime
	DarkRed      = colormap.DarkRed
	Red          = colormap.Red
	Pink         = colormap.Pink
	Orange       = colormap.Orange
	Yellow       = colormap.Yellow
	BrightYellow = colormap.BrightYellow
	DarkGray     = colormap.DarkGray
	MediumGray   = colormap.MediumGray
	Gray         = colormap.Gray
	White        = colormap.White
)

var (
	NewDefaultConfig = config.NewDefaultConfig
	LoggerConfig     = config.LoggerConfig

	// NewLogger creates a new Logger instance with the given ID and settings.
	// If `color` is set to `-1`, a color will be chosen automatically based on the ID.
	// If `debugMode` is set to `true`, debug level logging will be enabled.
	// If `messageHandler` is not `nil`, the logger will write to the provided handler instead of the screen.
	NewLogger = logger.NewLogger

	// NewLoggerSimple creates a new Logger instance with the given ID, automatically chosen color, and without debug mode or message handler.
	NewLoggerSimple = logger.NewLoggerSimple

	NewIndicator = indicator.NewIndicator

	NewTraceLine = logger.NewTraceLine

	NewTableColumnCustom = logger.NewTableColumnCustom
	// Related config setting(s):
	//
	//   - `LoggerConfig.TablePadChar`
	NewTableColumn             = logger.NewTableColumn
	NewTableColumnLeftCustom   = logger.NewTableColumnLeftCustom
	NewTableColumnLeft         = logger.NewTableColumnLeft
	NewTableColumnRightCustom  = logger.NewTableColumnRightCustom
	NewTableColumnRight        = logger.NewTableColumnRight
	NewTableColumnCenterCustom = logger.NewTableColumnCenterCustom
	NewTableColumnCenter       = logger.NewTableColumnCenter
	NewTable                   = logger.NewTable

	NewGError         = logger.NewGError
	NewGErrorRegistry = logger.NewGErrorRegistry

	// MapColor translates `index` from glog's color table to the corresponding ANSI color index.
	// Glog uses its own color table to make smooth (automated) color transitions easier to implement.
	MapColor = colormap.MapColor
)
