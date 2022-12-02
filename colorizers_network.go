package glog

import (
	"fmt"
	"net"
	"strings"

	"github.com/toxyl/gutils"
)

func enrichAndColorIPv4(ip string, useReverseDNS bool) string {
	parts := strings.Split(ip, ".")
	pt := 0.0
	for _, p := range parts {
		f, _ := gutils.GetFloat(p)
		pt += f
	}
	revDNS := "N/A"
	if useReverseDNS {
		if _, ok := LoggerConfig.reverseDNSCache[ip]; !ok {
			LoggerConfig.reverseDNSCache[ip] = gutils.ReverseDNS(ip)
		}
		revDNS = LoggerConfig.reverseDNSCache[ip]
	}
	ipColor := int(88.0 + 143.0*(pt/4.0/255.0)) // 88 - 231 (143 total)

	if revDNS != "N/A" {
		ip = fmt.Sprintf("%s (%s)", ip, revDNS)
	}
	ip = Wrap(ip, ipColor)

	return ip
}

func AddrIPv4Port(ip string, port int, useReverseDNS bool) string {
	ip = enrichAndColorIPv4(ip, useReverseDNS)
	return fmt.Sprintf("%s:%s", ip, Port(port))
}

func Addr(addrIPv4Port string, useReverseDNS bool) string {
	h, p := gutils.SplitHostPort(addrIPv4Port)
	return AddrIPv4Port(h, p, useReverseDNS)
}

func ConnRemote(conn net.Conn, useReverseDNS bool) string {
	return Addr(conn.RemoteAddr().String(), useReverseDNS)
}

func ConnLocal(conn net.Conn, useReverseDNS bool) string {
	return Addr(conn.LocalAddr().String(), useReverseDNS)
}

func Port(port int) string {
	// 94 - 231 (137 total)
	return Wrap(fmt.Sprint(port), int(94.0+137.0*(float64(port)/65535.0)))
}

func IPs(ips []string, useReverseDNS bool) string {
	hs := []string{}
	for _, h := range ips {
		hs = append(hs, enrichAndColorIPv4(h, useReverseDNS))
	}
	return strings.Join(hs, ", ")
}
