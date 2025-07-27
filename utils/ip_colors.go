package utils

import (
	"strings"
	"sync"
)

type IPColorCache struct {
	*sync.Mutex
	data map[string]int
}

func newIPColorCache() *IPColorCache {
	ipcc := &IPColorCache{
		Mutex: &sync.Mutex{},
		data:  map[string]int{},
	}
	return ipcc
}

func (ipcc *IPColorCache) Get(ip string) int {
	ipcc.Lock()
	defer ipcc.Unlock()

	if v, ok := ipcc.data[ip]; ok {
		return v
	}

	parts := strings.Split(ip, ".")
	pt := 0.0
	pl := float64(len(parts))
	for _, p := range parts {
		f, _ := GetFloat(p)
		pt += f / pl
	}
	ipcc.data[ip] = int(32.0 + Max(0.0, Min(185.0, 185.0*(pt/255.0))))
	return ipcc.data[ip]
}

var Icc *IPColorCache = newIPColorCache()
