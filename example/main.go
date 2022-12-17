package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/toxyl/glog"
	"github.com/toxyl/gutils"
)

var appLogger *glog.Logger = glog.NewLogger("", glog.Pink, false, nil)
var traceLogger *glog.Logger = glog.NewLogger("Trace", glog.BrightYellow, false, nil)
var messageTypesLogger *glog.Logger = glog.NewLogger("Message Type", glog.OliveGreen, false, nil)
var boolLogger *glog.Logger = glog.NewLogger("Bool", glog.Cyan, false, nil)
var intLogger *glog.Logger = glog.NewLogger("Int", glog.LightBlue, false, nil)
var uintLogger *glog.Logger = glog.NewLogger("Uint", glog.Blue, false, nil)
var floatLogger *glog.Logger = glog.NewLogger("Float", glog.DarkBlue, false, nil)
var percentageLogger *glog.Logger = glog.NewLogger("Percentage", glog.DarkGreen, false, nil)
var networkLogger *glog.Logger = glog.NewLogger("Network", glog.Green, false, nil)
var timeLogger *glog.Logger = glog.NewLogger("Time", glog.Yellow, false, nil)
var miscLogger *glog.Logger = glog.NewLogger("Misc", glog.Lime, false, func(msg string) {
	fmt.Print("With    colors: " + msg)                 // we just echo what we get
	fmt.Print("Without colors: " + glog.StripANSI(msg)) // and again, but without colors
})

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
	appLogger.Default("Sleeping a bit...")
	gutils.RandomSleep(100, 3000, time.Millisecond)
}

func demoMessageTypes() {
	appLogger.Info("%s", glog.HighlightInfo("MESSAGE TYPES"))

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
	appLogger.Default("")
	appLogger.Info("%s", glog.HighlightInfo("DATA TYPES"))

	boolLogger.Info("These are booleans: %s", glog.Bool(true, false))

	intLogger.Info(
		"This is a negative integer, a zero-value integer and a positive integer: %s",
		glog.Int(-23, 0, 32),
	)

	uintLogger.Info(
		"This is a a zero-value unsigned integer and a positive unsigned integer: %s",
		glog.Uint(0, 32),
	)

	floatLogger.Info(
		"This is a negative float64 (2 digits precision: %s), a zero-value float64 (1 digit precision: %s) and a positive float64 (3 digits precision: %s)",
		glog.Float64(-23.3223, 2),
		glog.Float64(0.0, 1),
		glog.Float64(32.2332, 3),
	)

	percentageLogger.Info(
		"This is a negative percentage (2 digits precision: %s), a zero-value percentage (1 digit precision: %s) and a positive percentage (3 digits precision: %s)",
		glog.Percentage(-0.3223, 2),
		glog.Percentage(0.0, 1),
		glog.Percentage(0.2332, 3),
	)

	appLogger.Info("You can also let glog choose a colorizer based on type: %s", glog.Auto(true, false, 1, -4, 0, 1.23, -73.64, -0.34321, 0.698765, time.Now(), 5*time.Second, "hello", "world"))
	appLogger.Info("Normalized float values (from %s to %s) will be shown as percentages (%s).", glog.Float64(-1.0, 1), glog.Float64(1.0, 1), glog.Auto(-1.0, 1.0))
}

func demoNetwork() {
	appLogger.Default("")
	appLogger.Info("%s", glog.HighlightInfo("NETWORK"))

	networkLogger.EnablePlainLog("network-plain.log")
	networkLogger.EnableColorLog("network-color.log")
	networkLogger.Info(
		"These are host:port combinations (with reverse DNS): %s, %s, %s",
		glog.AddrIPv4Port("8.8.8.8", 80, true),
		glog.AddrIPv4Port("172.217.168.206", 443, true),
		glog.AddrIPv4Port("160.20.152.105", 12343, true),
	)
	networkLogger.Info(
		"These are host:port combinations (without reverse DNS): %s, %s, %s",
		glog.AddrIPv4Port("8.8.8.8", 80, false),
		glog.AddrIPv4Port("172.217.168.206", 443, false),
		glog.AddrIPv4Port("160.20.152.105", 12343, false),
	)
	networkLogger.Info(
		"These are also host:port combinations: %s, %s, %s (with reverse DNS)",
		glog.Addr("8.8.8.8:80", true),
		glog.Addr("172.217.168.206:443", true),
		glog.Addr("160.20.152.105:12343", true),
	)
	networkLogger.Info(
		"These are also host:port combinations: %s, %s, %s (without reverse DNS)",
		glog.Addr("8.8.8.8:80", false),
		glog.Addr("172.217.168.206:443", false),
		glog.Addr("160.20.152.105:12343", false),
	)
	networkLogger.Info(
		"These are ports: %s, %s, %s, %s",
		glog.Port(80),
		glog.Port(443),
		glog.Port(8080),
		glog.Port(41230),
	)
	networkLogger.Info(
		"These are IPs (with reverse DNS): %s",
		glog.IPs([]string{"127.0.0.1", "8.8.8.8", "160.20.152.105", "255.255.255.255"}, true),
	)
	networkLogger.Info(
		"These are IPs (without reverse DNS): %s",
		glog.IPs([]string{"127.0.0.1", "8.8.8.8", "160.20.152.105", "255.255.255.255"}, false),
	)
	networkLogger.Info(
		"These are URLs: %s",
		glog.URL("https://www.google.com", "https://serverius.net", "http://some.unsafe.place-to-not-go.to"),
	)
	networkLogger.Info(
		"These are more URLs: %s",
		glog.URL(
			"https://colab.research.google.com/github/huggingface/notebooks/blob/main/diffusers/stable_diffusion.ipynb#scrollTo=yEErJFjlrSWS",
			"https://goteleport.com/docs/deploy-a-cluster/open-source/",
			"https://sam-koblenski.blogspot.com/2015/09/everyday-dsp-for-programmers-averaging.html",
			"https://random_user:random_password@random.domain/hello/world/?me=random_user&you=someone%%20else#here",
		),
	)
	networkLogger.DisablePlainLog()
	networkLogger.DisableColorLog()
}

func demoMisc() {
	appLogger.Default("")
	appLogger.Info("%s", glog.HighlightInfo("MISC"))

	miscLogger.Info(
		"These are passwords you should not use: %s, %s, %s, %s",
		glog.Password("password123"),
		glog.Password("HelloWorld"),
		glog.Password("admin"),
		glog.Password("ILoveSex"),
	)
	miscLogger.Info(
		"Here are a couple of reasons: %s, %s, %s, %s",
		glog.Reason("the weather sucked"),
		glog.Reason("my dog died"),
		glog.Reason("my dog ate my homework"),
		glog.Reason("it wasn't me"),
	)
	miscLogger.Info(
		"Sensitive files: %s, %s, %s, %s",
		glog.File("/etc/passwd"),
		glog.File("~/.ssh/id_rsa"),
		glog.File("/etc/profile"),
		glog.File("/etc/resolv.conf"),
	)

	appLogger.Default("")
	appLogger.Info("%s", glog.HighlightInfo("STRING HIGHLIGHTING WITH AUTOMATIC COLORING"))
	appLogger.Default("%s", glog.Highlight("This is a color highlighted text", "red", "green", "blue"))
	appLogger.Default("%s", glog.Highlight("This is another color highlighted text", "red", "green", "blue"))
	appLogger.Default("%s", glog.Highlight("Color is determined by each character,"))
	appLogger.Default("%s", glog.Highlight("i.e. identical strings produce identical colors."))
	appLogger.Default("%s", glog.Highlight("The algorithm is case-insensitive and only processes"))
	appLogger.Default("%s", glog.Highlight("the most common characters, others are truncated to the range."))
	appLogger.Default("%s", glog.Highlight("Highlighted text is cached, so the color does not have"))
	appLogger.Default("%s", glog.Highlight("to be recalculated every time."))

	appLogger.Default("")
	appLogger.Info("%s", glog.HighlightInfo("ANSI-AWARE VALUE PADDING"))
	appLogger.Default("Left:   %s", glog.PadLeft(glog.Auto(1.52, "hello", 10000, -0.1), 50, '-'))
	appLogger.Default("Right:  %s", glog.PadRight(glog.Auto(1.52, "hello", 10000, -0.1), 50, '+'))
	appLogger.Default("Center: %s", glog.PadCenter(glog.Auto(1.52, "hello", 10000, -0.1), 50, '='))

	appLogger.Default("")
	appLogger.Info("%s", glog.HighlightInfo("TRACING FUNCTION CALLS"))
	traceLogger.EnableTrace(3)
	fnA()
	traceLogger.DisableTrace()
	appLogger.Default("")
}

func demoDateTime() {
	// Hint: You can change the default time and date formats using these variables:
	//
	// glog.TimeFormat     = "15:04:05"
	// glog.TimeFormat12hr = "03:04:05pm"
	// glog.DateFormat     = "2006-01-02"
	//
	// Here you can find some examples for formats:
	// https://www.geeksforgeeks.org/how-to-get-current-date-and-time-in-various-format-in-golang/

	appLogger.Default("")
	appLogger.Info("%s", glog.HighlightInfo("DATE & TIME"))

	timeLogger.EnablePlainLog("time-plain.log")
	timeLogger.Info(
		"The operation took: %s, %s, %s, %s",
		glog.Duration(15),
		glog.Duration(60),
		glog.Duration(240),
		glog.Duration(644),
	)
	timeLogger.Info(
		"The operation took: %s, %s, %s, %s",
		glog.DurationMilliseconds(15),
		glog.DurationMilliseconds(60),
		glog.DurationMilliseconds(240),
		glog.DurationMilliseconds(644),
	)

	timeLogger.Info(
		"Current time (12hr): %s",
		glog.Time12hr(time.Now()),
	)
	timeLogger.Info(
		"Current time (24hr): %s",
		glog.Time(time.Now()),
	)
	timeLogger.Info(
		"Current Unix time: %s",
		glog.Timestamp(),
	)
	timeLogger.Info(
		"Today's date: %s",
		glog.Date(time.Now()),
	)
	timeLogger.Info(
		"Today's full date (12hr): %s",
		glog.DateTime12hr(time.Now()),
	)
	timeLogger.Info(
		"Today's full date (24hr): %s",
		glog.DateTime(time.Now()),
	)
	timeLogger.DisablePlainLog()
}

func demoColors() {
	appLogger.ID = "Colors"
	appLogger.ShowColors()
}

func main() {
	appLogger.Success("App booted, let's show you the demo then!")

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

	appLogger.Default(
		"App was %s seconds (%s milliseconds) active",
		glog.Runtime(),
		glog.RuntimeMilliseconds(),
	)
}
