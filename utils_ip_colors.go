package glog

import (
	"fmt"
	"strings"
	"sync"
)

type ipColCache struct {
	*sync.Mutex
	data map[string]int
}

func newIPColorCache() *ipColCache {
	ipcc := &ipColCache{
		Mutex: &sync.Mutex{},
		data:  map[string]int{},
	}
	return ipcc
}

func (ipcc *ipColCache) get(ip string) int {
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

var ipColorCache *ipColCache = newIPColorCache()

func enrichAndColorIPv4(ip string, useReverseDNS bool) string {
	revDNS := "N/A"
	if useReverseDNS {
		if _, ok := LoggerConfig.reverseDNSCache[ip]; !ok {
			LoggerConfig.reverseDNSCache[ip] = ReverseDNS(ip)
		}
		revDNS = LoggerConfig.reverseDNSCache[ip]
	}
	ipColor := ipColorCache.get(ip)

	if revDNS != "N/A" {
		ip = fmt.Sprintf("%s (%s)", ip, revDNS)
	}
	ip = Wrap(ip, ipColor)
	return ip
}
