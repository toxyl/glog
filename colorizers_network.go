package glog

import (
	"fmt"
	"net"
	"net/url"
	"strings"

	"github.com/toxyl/gutils"
)

var ipColorCache map[string]int = map[string]int{}

func getIPColor(ip string) int {
	if v, ok := ipColorCache[ip]; ok {
		return v
	}
	parts := strings.Split(ip, ".")
	pt := 0.0
	for _, p := range parts {
		f, _ := gutils.GetFloat(p)
		pt += f
	}
	ipColorCache[ip] = int(16.0 + 215.0*(pt/4.0/255.0)) // 16 - 231 (215 total)
	return ipColorCache[ip]
}

func enrichAndColorIPv4(ip string, useReverseDNS bool) string {
	revDNS := "N/A"
	if useReverseDNS {
		if _, ok := LoggerConfig.reverseDNSCache[ip]; !ok {
			LoggerConfig.reverseDNSCache[ip] = gutils.ReverseDNS(ip)
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
	return Wrap(fmt.Sprint(port), int(16.0+215.0*(float64(port)/65535.0))) // 16 - 231 (215 total)
}

func IPs(ips []string, useReverseDNS bool) string {
	hs := []string{}
	for _, h := range ips {
		hs = append(hs, enrichAndColorIPv4(h, useReverseDNS))
	}
	return strings.Join(hs, ", ")
}

func URL(raw ...string) string {
	out := []string{}
	for _, r := range raw {
		isAlive := true
		u, err := url.Parse(r)
		if err != nil {
			return r
		}

		res := ""
		res += Wrap(u.Scheme+Wrap("://", LoggerConfig.ColorURLSeparators), LoggerConfig.ColorScheme)
		if u.User != nil && u.User.Username() != "" {
			res += Wrap(u.User.Username(), LoggerConfig.ColorUser)
			if p, ok := u.User.Password(); ok {
				res += Wrap(":", LoggerConfig.ColorURLSeparators) + Wrap(p, LoggerConfig.ColorPassword)
			}
			res += Wrap("@", LoggerConfig.ColorURLSeparators)
		}
		ips, _ := net.LookupIP(u.Host)
		if len(ips) > 0 {
			res += Wrap(u.Host, getIPColor(ips[0].String()))
		} else {
			isAlive = false
			res += WrapRed(u.Host)
		}
		for i, pe := range strings.Split(u.Path, "/") {
			if pe == "" {
				continue
			}
			if i > 0 {
				res += Wrap("/", LoggerConfig.ColorURLSeparators)
			}
			res += Wrap(pe, LoggerConfig.ColorURLPath)
		}
		if len(u.Path) > 0 && string(u.Path[len(u.Path)-1]) == "/" {
			res += Wrap("/", LoggerConfig.ColorURLSeparators)
		}

		q := u.RawQuery
		if q != "" {
			res += Wrap("?", LoggerConfig.ColorURLSeparators)
			pairs := []string{}
			for _, pair := range strings.Split(q, "&") {
				if pair == "" {
					continue // empty pairs don't work
				}
				e := strings.Split(pair, "=")
				if len(e) == 1 {
					// we have a single key
					pairs = append(pairs, Wrap(e[0], LoggerConfig.ColorQueryKey))
				} else {
					// we have a key-value pair
					v := strings.Join(e[1:], "=")
					pairs = append(pairs,
						Wrap(e[0], LoggerConfig.ColorQueryKey)+
							Wrap("=", LoggerConfig.ColorURLSeparators)+
							Wrap(v, LoggerConfig.ColorQueryValue),
					)
				}
			}
			res += strings.Join(pairs, Wrap("&", LoggerConfig.ColorURLSeparators))
		}

		if u.Fragment != "" {
			res += Wrap("#", LoggerConfig.ColorURLSeparators) + Wrap(u.Fragment, LoggerConfig.ColorFragment)
		}
		if !isAlive {
			res = WrapRed("(dead domain: ") + res + WrapRed(")")
		}
		out = append(out, res)
	}
	return strings.Join(out, ", ")
}
