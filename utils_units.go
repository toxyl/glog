package glog

import "math"

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

type DurationScale int

const (
	SECOND        = 1.0
	MINUTE        = SECOND * 60.0
	HOUR          = MINUTE * 60.0
	DAY           = HOUR * 24.0
	WEEK          = DAY * 7.0
	YEAR_COMMON   = DAY * 365.0
	YEAR_LEAP     = DAY * 366.0
	YEAR_AVERAGE  = (3.0*YEAR_COMMON + YEAR_LEAP) / 4.0
	MONTH_COMMON  = YEAR_COMMON / 12.0
	MONTH_LEAP    = YEAR_LEAP / 12.0
	MONTH_AVERAGE = YEAR_AVERAGE / 12.0

	DURATION_SCALE_AVERAGE = DurationScale(0)
	DURATION_SCALE_COMMON  = DurationScale(1)
	DURATION_SCALE_LEAP    = DurationScale(2)
)

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
		newPrefixRaw("sec", SECOND, 0),
		newPrefixRaw("min", MINUTE, 3),
		newPrefixRaw("hr", HOUR, 6),
		newPrefixRaw("d", DAY, 9),
		newPrefixRaw("w", WEEK, 12),
	},
}

var durationScaleAverageYear prefixScale = prefixScale{
	prefixes: append(durationScaleSafe.prefixes,
		newPrefixRaw("m", MONTH_AVERAGE, 15),
		newPrefixRaw("y", YEAR_AVERAGE, 18),
	),
}

var durationScaleCommonYear prefixScale = prefixScale{
	prefixes: append(durationScaleSafe.prefixes,
		newPrefixRaw("m", MONTH_COMMON, 15),
		newPrefixRaw("y", YEAR_COMMON, 18),
	),
}

var durationScaleLeapYear prefixScale = prefixScale{
	prefixes: append(durationScaleSafe.prefixes,
		newPrefixRaw("m", MONTH_LEAP, 15),
		newPrefixRaw("y", YEAR_LEAP, 18),
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

func humanReadableSI[N Number](value N, unit string) (float64, int, string) {
	return siScale.process(float64(value), unit)
}

func humanReadableIEC[N Number](value N, unit string) (float64, int, string) {
	return iecScale.process(float64(value), unit)
}

func humanReadableShort[N Number](value N) (float64, int, string) {
	return shortScale.process(float64(value), "")
}

func humanReadableLong[N Number](value N) (float64, int, string) {
	return longScale.process(float64(value), "")
}

func humanReadableDuration[N Number](seconds N, scale DurationScale) (float64, int, string) {
	if scale == DURATION_SCALE_AVERAGE {
		return durationScaleAverageYear.process(float64(seconds), "")
	}
	if scale == DURATION_SCALE_COMMON {
		return durationScaleCommonYear.process(float64(seconds), "")
	}
	return durationScaleLeapYear.process(float64(seconds), "")
}
