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
		newPrefix("μ", 10, -6),
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
