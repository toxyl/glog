package glog

import (
	"math"
)

func Max[N Number](a, b N) N {
	return N(math.Max(float64(a), float64(b)))
}

func Min[N Number](a, b N) N {
	return N(math.Min(float64(a), float64(b)))
}
