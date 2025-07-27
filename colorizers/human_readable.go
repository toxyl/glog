package colorizers

import (
	"math"

	"github.com/toxyl/glog/ansi"
	"github.com/toxyl/glog/config"
	"github.com/toxyl/glog/types"
	"github.com/toxyl/glog/utils"
)

type prefix struct {
	symbol   string
	value    float64
	exponent int
}

type prefixScale struct {
	prefixes []prefix
}

func (ps prefixScale) process(value float64, unit string) (float64, int, string) {
	vAbs := math.Abs(value)
	prefix := ps.prefixes[0]

	for _, p := range ps.prefixes {
		if vAbs < p.value {
			break // no need to look further
		}
		prefix = p
	}
	return value / prefix.value, prefix.exponent, prefix.symbol + unit
}

func newPrefix(symbol string, base, exponent int) prefix {
	return prefix{symbol: symbol, exponent: exponent, value: math.Pow(float64(base), float64(exponent))}
}

func newPrefixRaw(symbol string, value float64, colorIndex int) prefix {
	return prefix{symbol: symbol, exponent: colorIndex, value: value}
}

var siScale prefixScale = prefixScale{
	prefixes: []prefix{
		newPrefix("q", 10, -30),
		newPrefix("r", 10, -27),
		newPrefix("y", 10, -24),
		newPrefix("z", 10, -21),
		newPrefix("a", 10, -18),
		newPrefix("f", 10, -15),
		newPrefix("p", 10, -12),
		newPrefix("n", 10, -9),
		newPrefix("µ", 10, -6),
		newPrefix("m", 10, -3),
		newPrefix("c", 10, -2),
		newPrefix(" ", 10, 0),
		newPrefix("k", 10, 3),
		newPrefix("M", 10, 6),
		newPrefix("G", 10, 9),
		newPrefix("T", 10, 12),
		newPrefix("P", 10, 15),
		newPrefix("E", 10, 18),
		newPrefix("Z", 10, 21),
		newPrefix("Y", 10, 24),
		newPrefix("R", 10, 27),
		newPrefix("Q", 10, 30),
	},
}

var durationScaleSafe prefixScale = prefixScale{
	prefixes: []prefix{
		// let's start easy-peasy with SI units and prefixes
		newPrefix("qs", 10, -30),
		newPrefix("rs", 10, -27),
		newPrefix("ys", 10, -24),
		newPrefix("zs", 10, -21),
		newPrefix("as", 10, -18),
		newPrefix("fs", 10, -15),
		newPrefix("ps", 10, -12),
		newPrefix("ns", 10, -9),
		newPrefix("µs", 10, -6),
		newPrefix("ms", 10, -3),
		// now the messy part
		newPrefixRaw("sec", utils.SECOND, 0),
		newPrefixRaw("min", utils.MINUTE, 3),
		newPrefixRaw("hr", utils.HOUR, 6),
		newPrefixRaw("d", utils.DAY, 9),
		newPrefixRaw("w", utils.WEEK, 12),
	},
}

var durationScaleAverageYear prefixScale = prefixScale{
	prefixes: append(durationScaleSafe.prefixes,
		newPrefixRaw("m", utils.MONTH_AVERAGE, 15),
		newPrefixRaw("y", utils.YEAR_AVERAGE, 18),
	),
}

var durationScaleCommonYear prefixScale = prefixScale{
	prefixes: append(durationScaleSafe.prefixes,
		newPrefixRaw("m", utils.MONTH_COMMON, 15),
		newPrefixRaw("y", utils.YEAR_COMMON, 18),
	),
}

var durationScaleLeapYear prefixScale = prefixScale{
	prefixes: append(durationScaleSafe.prefixes,
		newPrefixRaw("m", utils.MONTH_LEAP, 15),
		newPrefixRaw("y", utils.YEAR_LEAP, 18),
	),
}

var iecScale prefixScale = prefixScale{
	prefixes: []prefix{
		newPrefix("  ", 1024, 0),
		newPrefix("Ki", 1024, 1),
		newPrefix("Mi", 1024, 2),
		newPrefix("Gi", 1024, 3),
		newPrefix("Ti", 1024, 4),
		newPrefix("Pi", 1024, 5),
		newPrefix("Ei", 1024, 6),
		newPrefix("Zi", 1024, 7),
		newPrefix("Yi", 1024, 8),
		newPrefix("Ri", 1024, 9),
		newPrefix("Qi", 1024, 10),
	},
}

var shortScale prefixScale = prefixScale{
	prefixes: []prefix{
		newPrefix("", 10, 0),
		newPrefix("K", 10, 3),
		newPrefix("M", 10, 6),
		newPrefix("B", 10, 9),
		newPrefix("T", 10, 12),
		newPrefix("Q", 10, 15),
		newPrefix("QN", 10, 18),
		newPrefix("S", 10, 21),
		newPrefix("SN", 10, 24),
		newPrefix("O", 10, 27),
		newPrefix("N", 10, 30),
	},
}

var longScale prefixScale = prefixScale{
	prefixes: []prefix{
		newPrefix("", 10, 0),
		newPrefix("k", 10, 3),
		newPrefix("m", 10, 6),
		newPrefix("mrd", 10, 9),
		newPrefix("b", 10, 12),
		newPrefix("brd", 10, 15),
		newPrefix("t", 10, 18),
		newPrefix("trd", 10, 21),
		newPrefix("q", 10, 24),
		newPrefix("qrd", 10, 27),
		newPrefix("qn", 10, 30),
	},
}

func humanReadableSI[N types.Number](value N, unit string) (float64, int, string) {
	return siScale.process(float64(value), unit)
}

func humanReadableIEC[N types.Number](value N, unit string) (float64, int, string) {
	return iecScale.process(float64(value), unit)
}

func humanReadableShort[N types.Number](value N) (float64, int, string) {
	return shortScale.process(float64(value), "")
}

func humanReadableLong[N types.Number](value N) (float64, int, string) {
	return longScale.process(float64(value), "")
}

func humanReadableDuration[N types.Number](seconds N, scale utils.DurationScale) (float64, int, string) {
	if scale == utils.DURATION_SCALE_AVERAGE {
		return durationScaleAverageYear.process(float64(seconds), "")
	}
	if scale == utils.DURATION_SCALE_COMMON {
		return durationScaleCommonYear.process(float64(seconds), "")
	}
	return durationScaleLeapYear.process(float64(seconds), "")
}

func wrapHumanReadable(v float64, separator, s string, e int) string {
	if e < 0 {
		e = -e
	}
	e *= 2
	return Float(v, config.LoggerConfig.AutoFloatPrecision) + separator + ansi.Wrap(s, config.LoggerConfig.ColorUnitHumanReadable+e).String()
}

// HumanReadableShort colors the same way as Float() does but will make `n` human-readable using short scale suffixes (base1000).
//
// Related config setting(s):
//
//   - `LoggerConfig.AutoFloatPrecision`
//   - `LoggerConfig.ColorUnitHumanReadable` + exponent of the matching scale step
func HumanReadableShort[N types.Number](value N) string {
	v, e, u := humanReadableShort(value)
	return wrapHumanReadable(v, "", u, e)
}

// HumanReadableLong colors the same way as Float() does but will make `n` human-readable using long scale suffixes  (base1000).
//
// Related config setting(s):
//
//   - `LoggerConfig.AutoFloatPrecision`
//   - `LoggerConfig.ColorUnitHumanReadable` + exponent of the matching scale step
func HumanReadableLong[N types.Number](value N) string {
	v, e, u := humanReadableLong(value)
	return wrapHumanReadable(v, "", u, e)
}

// HumanReadableIEC colors the same way as Float() does but will make `n` human-readable using IEC-prefixes (base1024).
//
// Related config setting(s):
//
//   - `LoggerConfig.AutoFloatPrecision`
//   - `LoggerConfig.ColorUnitHumanReadable` + exponent of the matching scale step
func HumanReadableIEC[N types.Number](value N, unit string) string {
	v, e, u := humanReadableIEC(value, unit)
	return wrapHumanReadable(v, " ", u, e)
}

// HumanReadableSI colors the same way as Float() does but will make `n` human-readable using SI-prefixes (base1000).
//
// Related config setting(s):
//
//   - `LoggerConfig.AutoFloatPrecision`
//   - `LoggerConfig.ColorUnitHumanReadable` + exponent of the matching scale step
func HumanReadableSI[N types.Number](value N, unit string) string {
	v, e, u := humanReadableSI(value, unit)
	return wrapHumanReadable(v, " ", u, e)
}

// HumanReadableBytesIEC colors the same way as Float() does but will make `bytes` human-readable using IEC-prefixes (base1024).
//
// Related config setting(s):
//
//   - `LoggerConfig.AutoFloatPrecision`
//   - `LoggerConfig.ColorUnitHumanReadable` + exponent of the matching scale step
func HumanReadableBytesIEC[N types.Number](bytes N) string {
	return HumanReadableIEC(bytes, "B")
}

// HumanReadableBytesSI colors the same way as Float() does but will make `bytes` human-readable using SI-prefixes (base1000).
//
// Related config setting(s):
//
//   - `LoggerConfig.AutoFloatPrecision`
//   - `LoggerConfig.ColorUnitHumanReadable` + exponent of the matching scale step
func HumanReadableBytesSI[N types.Number](bytes N) string {
	return HumanReadableSI(bytes, "B")
}

// HumanReadableRateIEC colors the same way as Float() does but will make `n` human-readable using IEC-prefixes (base1024).
//
// `interval` defines the duration during which `n` accumulated.
//
// Related config setting(s):
//
//   - `LoggerConfig.AutoFloatPrecision`
//   - `LoggerConfig.ColorUnitHumanReadable` + exponent of the matching scale step
func HumanReadableRateIEC[N types.Number](n N, unit, interval string) string {
	v, e, u := humanReadableIEC(n, unit)
	return wrapHumanReadable(v, " ", u+"/"+interval, e)
}

// HumanReadableRateIEC colors the same way as Float() does but will make `n` human-readable using IEC-prefixes (base1024).
//
// `interval` defines the duration during which `n` accumulated.
//
// Related config setting(s):
//
//   - `LoggerConfig.AutoFloatPrecision`
//   - `LoggerConfig.ColorUnitHumanReadable` + exponent of the matching scale step
func HumanReadableRateSI[N types.Number](n N, unit, interval string) string {
	v, e, u := humanReadableSI(n, unit)
	return wrapHumanReadable(v, " ", u+"/"+interval, e)
}

// HumanReadableRateBytesIEC colors the same way as Float() does but will make `bytes` human-readable using IEC-prefixes (base1024).
//
// `interval` defines the duration during which `bytes` accumulated.
//
// Related config setting(s):
//
//   - `LoggerConfig.AutoFloatPrecision`
//   - `LoggerConfig.ColorUnitHumanReadable` + exponent of the matching scale step
func HumanReadableRateBytesIEC[N types.Number](bytes N, interval string) string {
	return HumanReadableRateIEC(bytes, "B", interval)
}

// HumanReadableRateBytesSI colors the same way as Float() does but will make `bytes` human-readable using SI-prefixes (base1000).
//
// `interval` defines the duration during which `bytes` accumulated.
//
// Related config setting(s):
//
//   - `LoggerConfig.AutoFloatPrecision`
//   - `LoggerConfig.ColorUnitHumanReadable` + exponent of the matching scale step
func HumanReadableRateBytesSI[N types.Number](bytes N, interval string) string {
	return HumanReadableRateSI(bytes, "B", interval)
}
