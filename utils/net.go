package utils

import (
	"context"
	"net"
	"strconv"
)

// SplitHostPort takes a "IP:Port" formatted string and returns the port as integer and the ip portion as string.
// If splitting fails the function returns an empty string as IP and 0 as port.
func SplitHostPort(ipv4Addr string) (string, int) {
	ip, port, err := net.SplitHostPort(ipv4Addr)
	if err != nil {
		return "", 0
	}
	p, err := strconv.Atoi(port)
	if err != nil {
		p = 0
	}
	return ip, p
}

func ReverseDNS(ip string) string {
	host, err := (&net.Resolver{}).LookupAddr(context.Background(), ip)
	if err != nil || len(host) <= 0 {
		return "N/A"
	}
	return host[0]
}
