package glog

import (
	"fmt"
	"strings"
)

var ipColorCache map[string]int = map[string]int{}

func getIPColor(ip string) int {
	if v, ok := ipColorCache[ip]; ok {
		return v
	}
	parts := strings.Split(ip, ".")
	pt := 0.0
	for _, p := range parts {
		f, _ := GetFloat(p)
		pt += f
	}
	ipColorCache[ip] = int(16.0 + 215.0*(pt/4.0/255.0)) // 16 - 231 (215 total)
	return ipColorCache[ip]
}

func enrichAndColorIPv4(ip string, useReverseDNS bool) string {
	revDNS := "N/A"
	if useReverseDNS {
		if _, ok := LoggerConfig.reverseDNSCache[ip]; !ok {
			LoggerConfig.reverseDNSCache[ip] = ReverseDNS(ip)
		}
		revDNS = LoggerConfig.reverseDNSCache[ip]
	}
	ipColor := getIPColor(ip)

	if revDNS != "N/A" {
		ip = fmt.Sprintf("%s (%s)", ip, revDNS)
	}
	ip = Wrap(ip, ipColor)
	return ip
}
