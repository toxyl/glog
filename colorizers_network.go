package glog

import (
	"fmt"
	"net"
	"net/url"
	"strings"
)

func AddrIPv4Port[I IntOrUint](ip string, port I, useReverseDNS bool) string {
	ip = enrichAndColorIPv4(ip, useReverseDNS)
	return fmt.Sprintf("%s:%s", ip, Port(port))
}

func Addr(addrIPv4Port string, useReverseDNS bool) string {
	h, p := SplitHostPort(addrIPv4Port)
	return AddrIPv4Port(h, p, useReverseDNS)
}

func ConnRemote(conn net.Conn, useReverseDNS bool) string {
	return Addr(conn.RemoteAddr().String(), useReverseDNS)
}

func ConnLocal(conn net.Conn, useReverseDNS bool) string {
	return Addr(conn.LocalAddr().String(), useReverseDNS)
}

func Port[I IntOrUint](port I) string {
	return Wrap(fmt.Sprint(port), int(32.0+Max(0.0, Min(183.0, 183.0*(float64(port)/65535.0)))))
}

func IPs(ips []string, useReverseDNS bool) string {
	hs := []string{}
	for _, h := range ips {
		hs = append(hs, enrichAndColorIPv4(h, useReverseDNS))
	}
	return strings.Join(hs, ", ")
}

// URL colorizes a URL and, if enabled, marks dead ones (based on a DNS check).
//
// Related config setting(s):
//
//   - `LoggerConfig.ColorURLSeparators`
//   - `LoggerConfig.ColorScheme`
//   - `LoggerConfig.ColorUser`
//   - `LoggerConfig.ColorPassword`
//   - `LoggerConfig.ColorURLPath`
//   - `LoggerConfig.ColorQueryKey`
//   - `LoggerConfig.ColorQueryValue`
//   - `LoggerConfig.ColorFragment`
//   - `LoggerConfig.CheckIfURLIsAlive`
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
		if LoggerConfig.CheckIfURLIsAlive {
			ips, _ := net.LookupIP(u.Host)
			if len(ips) > 0 {
				res += Wrap(u.Host, ipColorCache.get(ips[0].To4().String()))
			} else {
				isAlive = false
				res += WrapRed(u.Host)
			}
		} else {
			res += Wrap(u.Host, stringColorCache.Get(u.Host))
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
			res = WrapRed("💀 ") + res
		}
		out = append(out, res)
	}
	return strings.Join(out, ", ")
}
