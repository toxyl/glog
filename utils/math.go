package utils

import (
	"math"

	"github.com/toxyl/glog/types"
)

func Max[N types.Number](a, b N) N {
	return N(math.Max(float64(a), float64(b)))
}

func Min[N types.Number](a, b N) N {
	return N(math.Min(float64(a), float64(b)))
}
