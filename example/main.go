package main

import (
	"errors"
	"fmt"
	"math"
	"os"
	"strings"
	"time"

	"github.com/toxyl/glog"
	"github.com/toxyl/glog/ansi"
)

var appLogger *glog.Logger = glog.NewLogger("", glog.Pink, false, nil)
var traceLogger *glog.Logger = glog.NewLoggerSimple("Trace")
var colorLogger *glog.Logger = glog.NewLoggerSimple("Colors")
var stringLogger *glog.Logger = glog.NewLogger("Strings", glog.Purple, false, nil)
var messageTypesLogger *glog.Logger = glog.NewLoggerSimple("Message Type")
var boolLogger *glog.Logger = glog.NewLogger("Bool", glog.Cyan, false, nil)
var intLogger *glog.Logger = glog.NewLogger("Int", glog.LightBlue, false, nil)
var uintLogger *glog.Logger = glog.NewLogger("Uint", glog.Blue, false, nil)
var floatLogger *glog.Logger = glog.NewLogger("Float", glog.DarkBlue, false, nil)
var percentageLogger *glog.Logger = glog.NewLogger("Percentage", glog.DarkGreen, false, nil)
var humanReadableLogger *glog.Logger = glog.NewLoggerSimple("Human Readable")
var autoLogger *glog.Logger = glog.NewLoggerSimple("Auto")
var networkLogger *glog.Logger = glog.NewLogger("Network", glog.Green, false, nil)
var timeLogger *glog.Logger = glog.NewLoggerSimple("Time")
var progressLogger *glog.Logger = glog.NewLoggerSimple("Progress")
var tableLogger *glog.Logger = glog.NewLogger("Tables", glog.MediumGray, false, nil)
var miscLogger *glog.Logger = glog.NewLogger("Misc", -1, false, func(msg string) {
	fmt.Print("With    colors: " + msg)                 // we just echo what we get
	fmt.Print("Without colors: " + glog.StripANSI(msg)) // and again, but without colors
})

var ansiLogger *glog.Logger = glog.NewLogger("ANSI", glog.Yellow, false, nil)
var cursorLogger *glog.Logger = glog.NewLogger("Cursor", glog.Cyan, false, nil)
var screenLogger *glog.Logger = glog.NewLogger("Screen", glog.Red, false, nil)

func fnE() {
	traceLogger.Trace(1)
}

func fnD() {
	traceLogger.Trace(4)
	fnE()
}

func fnC() {
	traceLogger.Trace(3)
	fnD()
}

func fnB() {
	traceLogger.Trace(2)
	fnC()
}

func fnA() {
	traceLogger.Trace(1)
	fnB()
}

func sleep() {
	if ignoreSleeps {
		return
	}
	appLogger.Default("Sleeping a bit...")
	i := 0
	n := glog.GetRandomInt(1, 30)
	for i = 0; i < n; i++ {
		appLogger.Progress(float64(i)/float64(n), "done waiting...")
		time.Sleep(100 * time.Millisecond)
	}
	if n%2 == 0 {
		appLogger.ProgressSuccess(float64(i)/float64(n), "Done!")
		return
	}
	appLogger.ProgressError(float64(i)/float64(n), "Failed!")
}

func printSection(title string) {
	appLogger.Blank("")
	appLogger.Blank("%s", glog.HighlightInfo(title))
}

func demoMessageTypes() {
	printSection("MESSAGE TYPES")

	messageTypesLogger.Info("This is %s message", glog.HighlightInfo("an info"))
	messageTypesLogger.Success("This is %s message", glog.HighlightSuccess("a success"))
	messageTypesLogger.OK("This is %s message", glog.HighlightOK("an all-OK"))
	messageTypesLogger.NotOK("This is %s message", glog.HighlightNotOK("a not-OK"))
	messageTypesLogger.Warning("This is %s message", glog.HighlightWarning("a warning"))
	messageTypesLogger.Error("This is %s", glog.HighlightError("an error"))
	messageTypesLogger.Error("This is %s", glog.Error(errors.New("another error")))
	messageTypesLogger.Debug("This is %s that you are not going to see because we initialized the logger with debug set to false.", glog.HighlightDebug("a debug message"))
	messageTypesLogger.EnableDebug()
	messageTypesLogger.Debug("The last debug message you didn't see, but this %s you will see.", glog.HighlightDebug("debug message"))
	messageTypesLogger.DisableDebug()
	messageTypesLogger.Question("Do you have %s? Maybe the next sections will answer them.", glog.HighlightQuestion("a question"))
}

func demoDataTypes() {
	printSection("DATA TYPES")

	boolLogger.Default("These are booleans: %s", glog.Bool(true, false))

	intLogger.Default(
		"This is a negative integer, a zero-value integer and a positive integer: %s",
		glog.Int(-23, 0, 32),
	)

	uintLogger.Default(
		"This is a a zero-value unsigned integer and a positive unsigned integer: %s",
		glog.Uint(uint(0), uint(32)),
	)

	floatLogger.Default(
		"This is a negative float64 (2 digits precision: %s), a zero-value float64 (1 digit precision: %s) and a positive float64 (3 digits precision: %s)",
		glog.Float(-23.3223, 2),
		glog.Float(0.0, 1),
		glog.Float(32.2332, 3),
	)

	percentageLogger.Default(
		"This is a negative percentage (2 digits precision: %s), a zero-value percentage (1 digit precision: %s) and a positive percentage (3 digits precision: %s)",
		glog.Percentage(-0.3223, 2),
		glog.Percentage(0.0, 1),
		glog.Percentage(0.2332, 3),
	)

	printSection("AUTOMATIC TYPE DETECTION")
	autoLogger.Default("You can also let glog choose a colorizer based on type: %s", glog.Auto(true, false, 1, -4, 0, 1.23, -73.64, -0.34321, 0.698765, time.Now(), 5*time.Second, "hello", "world"))
	autoLogger.Default("Normalized float values (from %s to %s) will be shown as percentages (%s).", glog.Float(-1.0, 1), glog.Float(1.0, 1), glog.Auto(-1.0, 1.0))

	printSection("HUMAN READABLE NUMBERS")
	printSection("Meters:")
	for v := -30; v <= 30; v++ {
		humanReadableLogger.Default(
			"10^%s m -> %s %s",
			glog.PadRight(glog.Int(v), 3, ' '),
			glog.PadLeft(glog.HumanReadableSI(math.Pow10(v), "m"), 18, ' '),
			glog.PadLeft(glog.HumanReadableSI(-math.Pow10(v), "m"), 18, ' '),
		)
	}

	printSection("Short Scale (uppercase) vs. Long Scale (lowercase):")
	for v := 0; v <= 30; v++ {
		humanReadableLogger.Default(
			"10^%s | short: %s %s | long: %s %s",
			glog.PadRight(glog.Int(v), 3, ' '),
			glog.PadLeft(glog.HumanReadableShort(math.Pow10(v)), 23, ' '),
			glog.PadLeft(glog.HumanReadableShort(-math.Pow10(v)), 23, ' '),
			glog.PadLeft(glog.HumanReadableLong(math.Pow10(v)), 23, ' '),
			glog.PadLeft(glog.HumanReadableLong(-math.Pow10(v)), 23, ' '),
		)
	}

	printSection("Data Rate Bytes, IEC vs. SI (base1000):")
	for v := 0; v <= 10; v++ {
		dr := math.Pow(1000, float64(v))
		humanReadableLogger.Default(
			"1000^%s | IEC: %s %s %s %s | SI: %s %s %s %s",
			glog.PadRight(glog.Int(v), 3, ' '),
			glog.PadLeft(glog.HumanReadableRateBytesIEC(-dr, "s"), 18, ' '),
			glog.PadLeft(glog.HumanReadableRateBytesIEC(dr, "s"), 18, ' '),
			glog.PadLeft(glog.HumanReadableRateBytesIEC(dr*60, "m"), 18, ' '),
			glog.PadLeft(glog.HumanReadableRateBytesIEC(dr*60*60, "h"), 18, ' '),
			glog.PadLeft(glog.HumanReadableRateBytesSI(-dr, "s"), 18, ' '),
			glog.PadLeft(glog.HumanReadableRateBytesSI(dr, "s"), 18, ' '),
			glog.PadLeft(glog.HumanReadableRateBytesSI(dr*60, "m"), 18, ' '),
			glog.PadLeft(glog.HumanReadableRateBytesSI(dr*60*60, "h"), 18, ' '),
		)
	}

	printSection("Data Rate Bytes, IEC vs. SI (base1024):")
	for v := 0; v <= 10; v++ {
		dr := math.Pow(1024, float64(v))
		humanReadableLogger.Default(
			"1024^%s | IEC: %s %s %s %s | SI: %s %s %s %s",
			glog.PadRight(glog.Int(v), 3, ' '),
			glog.PadLeft(glog.HumanReadableRateBytesIEC(-dr, "s"), 18, ' '),
			glog.PadLeft(glog.HumanReadableRateBytesIEC(dr, "s"), 18, ' '),
			glog.PadLeft(glog.HumanReadableRateBytesIEC(dr*60, "m"), 18, ' '),
			glog.PadLeft(glog.HumanReadableRateBytesIEC(dr*60*60, "h"), 18, ' '),
			glog.PadLeft(glog.HumanReadableRateBytesSI(-dr, "s"), 18, ' '),
			glog.PadLeft(glog.HumanReadableRateBytesSI(dr, "s"), 18, ' '),
			glog.PadLeft(glog.HumanReadableRateBytesSI(dr*60, "m"), 18, ' '),
			glog.PadLeft(glog.HumanReadableRateBytesSI(dr*60*60, "h"), 18, ' '),
		)
	}

	printSection("Data Rate Requests, IEC vs. SI (base1000):")
	for v := 0; v <= 10; v++ {
		humanReadableLogger.Default(
			"1000^%s | IEC: %s %s | SI: %s %s",
			glog.PadRight(glog.Int(v), 3, ' '),
			glog.PadLeft(glog.HumanReadableRateIEC(math.Pow(1000, float64(v)), "Req", "m"), 18, ' '),
			glog.PadLeft(glog.HumanReadableRateIEC(-math.Pow(1000, float64(v)), "Req", "m"), 18, ' '),
			glog.PadLeft(glog.HumanReadableRateSI(math.Pow(1000, float64(v)), "Req", "m"), 18, ' '),
			glog.PadLeft(glog.HumanReadableRateSI(-math.Pow(1000, float64(v)), "Req", "m"), 18, ' '),
		)
	}

	printSection("Data Rate Requests, IEC vs. SI (base1024):")
	for v := 0; v <= 10; v++ {
		humanReadableLogger.Default(
			"1024^%s | IEC: %s %s | SI: %s %s",
			glog.PadRight(glog.Int(v), 3, ' '),
			glog.PadLeft(glog.HumanReadableRateIEC(math.Pow(1024, float64(v)), "Req", "m"), 18, ' '),
			glog.PadLeft(glog.HumanReadableRateIEC(-math.Pow(1024, float64(v)), "Req", "m"), 18, ' '),
			glog.PadLeft(glog.HumanReadableRateSI(math.Pow(1024, float64(v)), "Req", "m"), 18, ' '),
			glog.PadLeft(glog.HumanReadableRateSI(-math.Pow(1024, float64(v)), "Req", "m"), 18, ' '),
		)
	}

	printSection("Data Amount, IEC vs. SI (base1000):")
	for v := 0; v <= 10; v++ {
		humanReadableLogger.Default(
			"1000^%s | IEC: %s %s | SI: %s %s",
			glog.PadRight(glog.Int(v), 3, ' '),
			glog.PadLeft(glog.HumanReadableBytesIEC(math.Pow(1000, float64(v))), 18, ' '),
			glog.PadLeft(glog.HumanReadableBytesIEC(-math.Pow(1000, float64(v))), 18, ' '),
			glog.PadLeft(glog.HumanReadableBytesSI(math.Pow(1000, float64(v))), 18, ' '),
			glog.PadLeft(glog.HumanReadableBytesSI(-math.Pow(1000, float64(v))), 18, ' '),
		)
	}

	printSection("Data Amount, IEC vs. SI (base1024):")
	for v := 0; v <= 10; v++ {
		humanReadableLogger.Default(
			"1024^%s | IEC: %s %s | SI: %s %s",
			glog.PadRight(glog.Int(v), 3, ' '),
			glog.PadLeft(glog.HumanReadableBytesIEC(math.Pow(1024, float64(v))), 18, ' '),
			glog.PadLeft(glog.HumanReadableBytesIEC(-math.Pow(1024, float64(v))), 18, ' '),
			glog.PadLeft(glog.HumanReadableBytesSI(math.Pow(1024, float64(v))), 18, ' '),
			glog.PadLeft(glog.HumanReadableBytesSI(-math.Pow(1024, float64(v))), 18, ' '),
		)
	}
}

func demoNetwork() {
	printSection("NETWORK")

	networkLogger.EnablePlainLog("network-plain.log")
	networkLogger.EnableColorLog("network-color.log")

	glog.LoggerConfig.CheckIfURLIsAlive = false
	networkLogger.Default(
		`
These are host:port combinations (with reverse DNS): %s, %s
These are host:port combinations (without reverse DNS): %s, %s
These are also host:port combinations: %s, %s (with reverse DNS)
These are also host:port combinations: %s, %s (without reverse DNS)
These are ports: %s, %s, %s, %s
These are IPs (with reverse DNS): %s
These are IPs (without reverse DNS): %s
These are URLs: %s
These are more URLs: %s
Even more URLs: %s
		`,
		glog.AddrIPv4Port("8.8.8.8", 80, true),
		glog.AddrIPv4Port("172.217.168.206", 443, true),
		glog.AddrIPv4Port("8.8.8.8", 80, false),
		glog.AddrIPv4Port("172.217.168.206", 443, false),
		glog.Addr("8.8.8.8:80", true),
		glog.Addr("172.217.168.206:443", true),
		glog.Addr("8.8.8.8:80", false),
		glog.Addr("172.217.168.206:443", false),
		glog.Port(80),
		glog.Port(443),
		glog.Port(8080),
		glog.Port(41230),
		glog.IPs([]string{"127.0.0.1", "8.8.8.8", "255.255.255.255"}, true),
		glog.IPs([]string{"127.0.0.1", "8.8.8.8", "255.255.255.255"}, false),
		glog.URL("https://www.google.com", "http://some.unsafe.place-to-not-go.to"),
		glog.URL(
			"https://sam-koblenski.blogspot.com/2015/09/everyday-dsp-for-programmers-averaging.html",
			"https://random_user:random_password@random.domain/hello/world/?me=random_user&you=someone%%20else#here",
		),
		glog.URL(
			"https://colab.research.google.com/github/huggingface/notebooks/blob/main/diffusers/stable_diffusion.ipynb#scrollTo=yEErJFjlrSWS",
			"https://goteleport.com/docs/deploy-a-cluster/open-source/",
			"https://this.domain.should.be.dead",
		),
	)
	networkLogger.DisablePlainLog()
	networkLogger.DisableColorLog()
}

func demoMisc() {
	printSection("MISC")

	miscLogger.Default(
		`
These are passwords you should not use: %s, %s, %s, %s
Here are a couple of reasons: %s, %s, %s, %s
Sensitive files: %s, %s, %s, %s
`,
		glog.Password("password123"),
		glog.Password("HelloWorld"),
		glog.Password("admin"),
		glog.Password("ILoveSex"),
		glog.Reason("the weather sucked"),
		glog.Reason("my dog died"),
		glog.Reason("my dog ate my homework"),
		glog.Reason("it wasn't me"),
		glog.File("/etc/passwd"),
		glog.File("~/.ssh/id_rsa"),
		glog.File("/etc/profile"),
		glog.File("/etc/resolv.conf"),
	)

	printSection("STRING HIGHLIGHTING WITH AUTOMATIC COLORING")
	stringLogger.Default(`
%s
%s
%s
%s
%s
%s
%s
%s
%s
%s
`,
		glog.Highlight("This is a color highlighted text", "red", "green", "blue"),
		glog.Highlight("This is another color highlighted text", "red", "green", "blue"),
		glog.Highlight("Color is determined by each character,"),
		glog.Highlight("i.e. identical strings produce identical colors."),
		glog.Highlight("The algorithm is case-sensitive and only processes ASCII characters,"),
		glog.Highlight("others are truncated to the ASCII range."),
		glog.Highlight("Highlighted text is cached, so the color does not have"),
		glog.Highlight("to be recalculated every time."),
		glog.Highlight("a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "1", "2", "3", "4", "5", "6", "7", "8", "9", "0"),
		glog.Highlight("A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "!", "@", "#", "$", "%", "^", "&", "*", "(", ")"),
	)

	printSection("ANSI-AWARE VALUE PADDING")
	stringLogger.Default(`
"Left:   %s"
"Right:  %s", 
"Center: %s", 
`,
		glog.PadLeft(glog.Auto(1.52, "hello", 10000, -0.1), 50, '-'),
		glog.PadRight(glog.Auto(1.52, "hello", 10000, -0.1), 50, '+'),
		glog.PadCenter(glog.Auto(1.52, "hello", 10000, -0.1), 50, '='),
	)

	printSection("TRACING FUNCTION CALLS")
	traceLogger.EnableTrace(3)
	fnA()
	traceLogger.DisableTrace()

	printSection("PROGRESS BARS")
	progressLogger.Info(`
%s complete
%s complete 
%s complete 
%s complete 
%s complete 
%s complete 
%s complete 
%s complete 
`,
		glog.ProgressBar(float64(glog.GetRandomInt(0, 10000))/10000.0, 40),
		glog.ProgressBar(float64(glog.GetRandomInt(0, 10000))/10000.0, 40),
		glog.ProgressBar(float64(glog.GetRandomInt(0, 10000))/10000.0, 40),
		glog.ProgressBar(float64(glog.GetRandomInt(0, 10000))/10000.0, 40),
		glog.ProgressBar(float64(glog.GetRandomInt(0, 10000))/10000.0, 40),
		glog.ProgressBar(float64(glog.GetRandomInt(0, 10000))/10000.0, 40),
		glog.ProgressBar(float64(glog.GetRandomInt(0, 10000))/10000.0, 40),
		glog.ProgressBar(float64(glog.GetRandomInt(0, 10000))/10000.0, 40),
	)
}

func demoDateTime() {
	// Hint: You can change the default time and date formats using these variables:
	//
	// glog.LoggerConfig.TimeFormat     = "15:04:05"
	// glog.LoggerConfig.TimeFormat12hr = "03:04:05pm"
	// glog.LoggerConfig.DateFormat     = "2006-01-02"
	//
	// Here you can find some examples for formats:
	// https://www.geeksforgeeks.org/how-to-get-current-date-and-time-in-various-format-in-golang/

	printSection("DATE, TIME & DURATION")
	printSection("Duration:")
	timeLogger.EnablePlainLog("time-plain.log")
	timeLogger.Default(
		`The operations took: %s, %s, %s, %s
                    %s, %s, %s, %s`,
		glog.Duration(15),
		glog.Duration(60),
		glog.Duration(240.5),
		glog.Duration(644),
		glog.DurationMilliseconds(15),
		glog.DurationMilliseconds(60.5),
		glog.DurationMilliseconds(240),
		glog.DurationMilliseconds(644),
	)

	printSection("Duration (short):")
	for _, v := range []float64{
		glog.SECOND / 100000,
		glog.SECOND / 10000,
		glog.SECOND / 1000,
		glog.SECOND / 10,
		glog.SECOND,
		glog.SECOND * 30,
		glog.MINUTE,
		glog.MINUTE * 30,
		glog.HOUR,
		glog.HOUR * 12,
		glog.DAY,
		glog.WEEK,
		glog.WEEK * 4,
		glog.MONTH_AVERAGE,
		glog.YEAR_AVERAGE,
		glog.YEAR_AVERAGE * 10,
		glog.YEAR_AVERAGE * 100,
		glog.YEAR_AVERAGE * 1000,
	} {
		timeLogger.Default(
			"%ss = Average year: %s, Common year: %s, Leap year: %s",
			glog.PadLeft(glog.Float(v, 4), 24, ' '),
			glog.PadLeft(glog.DurationShort(v, glog.DURATION_SCALE_AVERAGE), 18, ' '),
			glog.PadLeft(glog.DurationShort(v, glog.DURATION_SCALE_COMMON), 18, ' '),
			glog.PadLeft(glog.DurationShort(v, glog.DURATION_SCALE_LEAP), 18, ' '),
		)
	}

	printSection("Time Related:")
	timeLogger.Default(
		`
Time:
	12hr: %s
	24hr: %s
	Unix: %s

Date:
	Today: %s

Date & Time:
	12hr: %s
	24hr: %s
`,
		glog.Time12hr(time.Now()),
		glog.Time(time.Now()),
		glog.Timestamp(),
		glog.Date(time.Now()),
		glog.DateTime12hr(time.Now()),
		glog.DateTime(time.Now()),
	)

	timeLogger.DisablePlainLog()
}

func demoColors() {
	printSection("COLORS")
	colorLogger.ShowColors()
}

func demoTables() {
	printSection("TABLES")

	tableLogger.KeyValueTable(
		map[string]any{
			"key 1": 123,
			"key 2": 0.123,
			"key 3": -1,
			"key 4": "hello world",
			"key 5": []string{"this", "is", "useful", "for", "debugging"},
		},
	)

	// let's make a table first
	col1 := glog.NewTableColumnLeft("Left").Push(
		10,
		" ",
		"---",
		[]any{"hello world", nil, 2, []string{"nesting", "works", "too"}, 0.30},
		" ",
		"---",
		" ",
		nil,
		-85,
		80,
		0.001,
		"https://www.google.com",
		"http://some.unsafe.place-to-not-go.to",
		"localhost",
		"/a/file/path",
	)
	col2 := glog.NewTableColumnCenter("Center").Push(
		false,
		"my little pony",
		50,
		60,
		time.Now(),
		"---",
		"  ",
		90,
		" ",
		"---",
		" ",
	)
	col3 := glog.NewTableColumnRight("Right").Push(
		10,
		20,
		true,
		[]int{40, 50},
		60,
		"---",
		"   ",
		10*time.Second,
		"---",
		" ",
		"   ",
		"",
		"care to log in?",
		100,
	)
	col4 := glog.NewTableColumnCenterCustom("Pad Char", '∙', nil).Push(
		-10,
		[]any{0, 5, 1.4, "test"},
		"---",
		" ",
		30.0/2.9,
		"---",
		"    ",
		nil,
		"---",
		false,
		"so long and",
		"thanks for all",
		"the fish",
		time.Now(),
		90,
		100,
	)
	col5 := glog.NewTableColumnCenterCustom("No Highlight", ' ', fmt.Sprint).Push(
		-10,
		" ",
		"---",
		glog.Auto([]any{0, 5, 1.4, "test"}),
		30.0/2.9,
		"---",
		"     ",
		nil,
		false,
		"so long and",
		"thanks for all",
		"the fish",
		time.Now(),
		90,
		100,
	)

	printSection("TABLE W/ HEADER ROW")
	tbl := glog.NewTable(col1, col2, col3, col4, col5)
	tbl.Print(tableLogger)

	printSection("TABLE W/O HEADER ROW")
	tbl.PrintWithoutHeader(tableLogger)

	printSection("TABLE CSV")
	tableLogger.Blank(tbl.CSV(','))

	printSection("TABLE TSV")
	tableLogger.Blank(tbl.CSV('\t'))

	printSection("TABLE YAML")
	if data, err := tbl.YAML(); err == nil {
		tableLogger.Blank(data)
	}

	printSection("TABLE JSON")
	if data, err := tbl.JSON(); err == nil {
		tableLogger.Blank(strings.ReplaceAll(data, "],[", "],\n["))
	}
}

var (
	errTestGo  = fmt.Errorf("test error")
	errTestGo2 = fmt.Errorf("test error 2")
	testErrors = glog.NewGErrorRegistry().Append(
		glog.NewGError(errTestGo, false, 0),
		glog.NewGError(errTestGo2, true, 20),
	).Register(
		fmt.Errorf("this is not fatal"),
		false,
		0,
	)
)

func somewhatShorterFunctionName() {
	testErrors.Check(errTestGo2, "should DIE! (fatal)")
}

func horriblyLongFunctionNameForTesting() {
	testErrors.Check(errTestGo, "should not die (non-fatal)")
	somewhatShorterFunctionName()
}

func demoErrorHandling() {
	printSection("ERROR HANDLING")
	testErrors.Check(errTestGo, "should not die (non-fatal)")
	testErrors.Check(fmt.Errorf("test error"), "should not die (non-fatal)")
	testErrors.Check(fmt.Errorf("some other: test error"), "should not die (non-fatal)")
	testErrors.Check(fmt.Errorf("test error: more errors"), "should not die (non-fatal)")
	testErrors.Check(fmt.Errorf("it can be somewhere: test error: in the chain"), "should not die (non-fatal)")
	testErrors.Check(fmt.Errorf("this is not fatal"), "should not die (non-fatal)")
	testErrors.Check(errTestGo, "should not die (non-fatal)")
	printSection("")
}

func demoFatalErrorHandling() {
	printSection("FATAL ERROR HANDLING")
	fn := func() {
		testErrors.Check(errTestGo, "should not die (non-fatal)")
		fn := func() {
			horriblyLongFunctionNameForTesting()
		}
		fn()
	}
	fn()
}

func demoANSIUtilities() {
	printSection("ANSI TEXT FORMATTING")

	ansiLogger.Default("Basic formatting: %s%s%s%s%s%s%s%s%s",
		ansi.Bold().Apply("Bold")+" ",
		ansi.Dim().Apply("Dim")+" ",
		ansi.Italic().Apply("Italic")+" ",
		ansi.Underline().Apply("Underline")+" ",
		ansi.DoubleUnderline().Apply("DoubleUnderline")+" ",
		ansi.StrikeThrough().Apply("StrikeThrough")+" ",
		ansi.Reverse().Apply("Reverse")+" ",
		ansi.Conceal().Apply("Conceal")+" ",
		ansi.Reset().String(),
	)

	printSection("BACKGROUND COLORS")
	ansiLogger.Default("Background colors: %s %s %s %s %s",
		ansi.WrapBackground("Red Background", glog.Red),
		ansi.WrapBackground("Green Background", glog.Green),
		ansi.WrapBackground("Blue Background", glog.Blue),
		ansi.WrapBackground("Yellow Background", glog.Yellow),
		ansi.WrapBackground("Cyan Background", glog.Cyan),
	)
}

func demoCursorMovement() {
	printSection("CURSOR MOVEMENT")

	cursorLogger.Default("Let's play Tic-Tac-Toe")

	// Demonstrate cursor movement
	cA, cB, cC := glog.NewTableColumnCenter("a"), glog.NewTableColumnCenter("b"), glog.NewTableColumnCenter("c")
	cA.Push("-", "-", "-")
	cB.Push("-", "-", "-")
	cC.Push("-", "-", "-")
	cursorLogger.TableWithoutHeader(cA, cB, cC)
	cursorLogger.Default("Line 1: Original text")

	// Move cursor up and play tic-tac-toe
	ansi.CursorUp(5).ToStdout()
	ansi.CursorForward(56).ToStdout()
	ansi.New("x").ToStdout()

	time.Sleep(500 * time.Millisecond)

	ansi.CursorForward(7).ToStdout()
	ansi.New("o").ToStdout()

	time.Sleep(500 * time.Millisecond)

	ansi.CursorDown(1).ToStdout()
	ansi.CursorBackward(5).ToStdout()
	ansi.New("x").ToStdout()

	time.Sleep(500 * time.Millisecond)

	ansi.CursorForward(3).ToStdout()
	ansi.New("o").ToStdout()

	time.Sleep(500 * time.Millisecond)

	ansi.CursorDown(1).ToStdout()
	ansi.CursorBackward(1).ToStdout()
	ansi.New("x").ToStdout()

	time.Sleep(500 * time.Millisecond)

	// Move cursor back to end
	ansi.CursorDown(2).ToStdout()
	ansi.CursorBackward(11).ToStdout()
	for _, c := range "Sorry `o`, `x` has won!" {
		ansi.New(string(c)).ToStdout()
		time.Sleep(500 * time.Millisecond)
	}
	ansi.ClearToEOL().ToStdout()
	ansi.Ln().ToStdout()     // Progress to next line
	cursorLogger.Default("") // Add a blank line after our demo

	printSection("CURSOR POSITIONING")
	cursorLogger.Default("") // Add a blank line after our demo
	ansi.CursorUp(1).ToStdout()
	ansi.CursorForward(56).ToStdout()
	ansi.StoreCursor().ToStdout()
	ansi.New("Hello there! This line will be overwritten.").ToStdout()
	time.Sleep(2 * time.Second)
	ansi.RestoreCursor().ToStdout()
	for _, c := range "Overwritten!" {
		ansi.New(string(c)).ToStdout()
		time.Sleep(500 * time.Millisecond)
	}
	ansi.ClearToEOL().ToStdout()
	ansi.Ln().ToStdout() // Progress to next line
	time.Sleep(2 * time.Second)

	// Move to end
	ansi.CursorDown(3).ToStdout()
	cursorLogger.Default("") // Add space after demo
}

func demoScreenControl() {
	printSection("SCREEN CONTROL")

	screenLogger.Default("Demonstrating screen clearing capabilities:")
	screenLogger.Default("This text will be cleared in 2 seconds...")

	// Wait a bit to show the text
	time.Sleep(2 * time.Second)

	// Clear from cursor to end of screen
	ansi.ClearScreenFromCursor().ToStdout()
	screenLogger.Default("Screen cleared from cursor to end!")

	time.Sleep(1 * time.Second)

	// Clear entire screen
	ansi.ClearScreen().ToStdout()
	screenLogger.Default("Entire screen cleared!")
	screenLogger.Default("Cursor positioned at top-left")

	printSection("LINE CLEARING")
	screenLogger.Default("Line 1: This will be partially cleared")
	screenLogger.Default("Line 2: This will be partially cleared")
	screenLogger.Default("Line 3: This will be partially cleared")

	// Move cursor up and clear lines
	ansi.CursorUp(3).ToStdout()
	ansi.CursorForward(15).ToStdout()
	ansi.ClearToEOL().ToStdout()
	ansi.New("CLEARED!").ToStdout()

	ansi.CursorDown(1).ToStdout()
	ansi.CursorForward(15).ToStdout()
	ansi.ClearToEOL().ToStdout()
	ansi.New("CLEARED!").ToStdout()

	ansi.CursorDown(1).ToStdout()
	ansi.CursorForward(15).ToStdout()
	ansi.ClearToEOL().ToStdout()
	ansi.New("CLEARED!").ToStdout()

	ansi.CursorDown(1).ToStdout()
	screenLogger.Default("") // Add space after demo
}

func demoInteractiveProgress() {
	printSection("INTERACTIVE PROGRESS DEMO")

	progressLogger.Default("Starting interactive progress demo...")

	// Hide cursor for cleaner output
	ansi.HideCursor().ToStdout()
	defer ansi.ShowCursor().ToStdout()

	for i := 0; i <= 100; i += 5 {
		// Move cursor up and clear line
		ansi.CursorUp(1).ToStdout()
		ansi.CursorToLineStartAndClear().ToStdout()
		time.Sleep(100 * time.Millisecond)
	}

	ansi.CursorDown(1).ToStdout()
	progressLogger.Default("Progress demo completed!")
}

func demoTextEffects() {
	printSection("TEXT EFFECTS DEMO")

	ansiLogger.Default("Demonstrating various text effects:")

	// Concealed text (password-like)
	ansiLogger.Default("Concealed text: %s%s",
		ansi.Conceal().Apply("secret123"),
		" - Useful for password prompts")
}

func demoLineWrapping() {
	printSection("LINE WRAPPING CONTROL")

	screenLogger.Default("Demonstrating line wrapping control:")

	// Disable line wrapping
	ansi.DisableLineWrap().ToStdout()
	screenLogger.Default("This is a very long line that should not wrap even if it exceeds the terminal width. It will continue on the same line until it reaches the edge of the terminal window.")

	// Re-enable line wrapping
	ansi.EnableLineWrap().ToStdout()
	screenLogger.Default("This line should wrap normally if it's too long for the terminal width.")

	screenLogger.Default("Line wrapping restored to normal behavior.")
}

func demoQuestionInline() {
	printSection("(ASKING) QUESTIONS")
	log := glog.NewLoggerSimple("Questions")

	log.Question("This is a question without user input,\nit could be used to mark items that need review.")
	log.BlankAuto(
		"Hi %s, you are %s and you can be reached at %s.",
		log.Ask("What is your name?", "string", "John Doe"),
		log.Ask("Enter your age: ", "int", "30"),
		log.Ask("Enter your email: ", "email", "john@example.com"),
	)
	correct := log.AskBoolWithTimeout("Is that correct?", true, 5*time.Second)
	if correct {
		log.Success("You said the info is correct.")
	} else {
		log.NotOK("Do I have to lookup your data for you?\nJust kidding!")
	}
	log.Blank("")
}

var (
	ignoreSleeps = false
)

func main() {
	if len(os.Args) == 2 && os.Args[1] == "no-sleeps" {
		ignoreSleeps = true
	}

	appLogger.Success("App booted, let's show you the demo then!")

	glog.LoggerConfig.SplitOnNewLine = true
	appLogger.Info("There is a new feature that you can enable if you like.")
	appLogger.BlankAuto(
		`
By default it is disabled to not break previous hacks
that implement or exploit this behavior. 
For this example it's now on.

Set %s to %s 
to enable splitting messages on new lines and 
prefixing each line according to the %s.
`,
		"glog.LoggerConfig.SplitOnNewLine",
		true,
		"LoggerConfig",
	) // ignore warning, *Auto functions convert everything to strings....

	demoQuestionInline()
	sleep()

	demoMessageTypes()
	sleep()

	demoDataTypes()
	sleep()

	demoNetwork()
	sleep()

	demoDateTime()
	sleep()

	demoMisc()
	sleep()

	demoColors()
	sleep()

	// New ANSI utility demonstrations
	demoANSIUtilities()
	sleep()

	demoCursorMovement()
	sleep()

	demoScreenControl()
	sleep()

	demoInteractiveProgress()
	sleep()

	demoTextEffects()
	sleep()

	demoLineWrapping()
	sleep()

	demoTables()
	sleep()

	demoErrorHandling()
	sleep()

	appLogger.Success(
		"App was %s seconds (%s milliseconds) active",
		glog.Runtime(),
		glog.RuntimeMilliseconds(),
	)

	demoFatalErrorHandling()
}
