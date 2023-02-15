package glog

import (
	"sync"
)

type StringColorCache struct {
	data map[string]int
	lock *sync.Mutex
}

// map to cache results to avoid repeatedly calculating the color for the same string
var stringColorCache StringColorCache = StringColorCache{
	data: map[string]int{},
	lock: &sync.Mutex{},
}

// Get takes an input string and maps it to a color within the ANSI range 16-231, i.e. it can generate 215 different colors.
func (scc *StringColorCache) Get(str string) int {
	scc.lock.Lock()
	defer scc.lock.Unlock()
	if v, ok := scc.data[str]; ok {
		return v
	}
	pt := 0.0
	bm := []rune(RemoveNonPrintable(str))
	l := len(bm)
	for i := 0; i < l; i++ {
		pt += (Max(0.0, Min(94.0, float64(bm[i])-32.0)) / 94.0) / float64(l)
	}
	scc.data[str] = int(16.0 + 215.0*pt) // 16 - 231 (215 total)
	return scc.data[str]
}
