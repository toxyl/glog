package colorizers

import (
	"fmt"
	"net"
	"net/url"
	"strings"

	"github.com/toxyl/glog/ansi"
	"github.com/toxyl/glog/config"
	"github.com/toxyl/glog/types"
	"github.com/toxyl/glog/utils"
	"github.com/toxyl/math"
)

func enrichAndColorIPv4(ip string, useReverseDNS bool) string {
	revDNS := "N/A"
	if useReverseDNS {
		if _, ok := config.LoggerConfig.ReverseDNSCache[ip]; !ok {
			config.LoggerConfig.ReverseDNSCache[ip] = utils.ReverseDNS(ip)
		}
		revDNS = config.LoggerConfig.ReverseDNSCache[ip]
	}
	ipColor := utils.Icc.Get(ip)

	if revDNS != "N/A" {
		ip = fmt.Sprintf("%s (%s)", ip, revDNS)
	}
	ip = ansi.Wrap(ip, ipColor).String()
	return ip
}

func AddrIPv4Port[I types.IntOrUint](ip string, port I, useReverseDNS bool) string {
	ip = enrichAndColorIPv4(ip, useReverseDNS)
	return fmt.Sprintf("%s:%s", ip, Port(port))
}

func Addr(addrIPv4Port string, useReverseDNS bool) string {
	h, p := utils.SplitHostPort(addrIPv4Port)
	return AddrIPv4Port(h, p, useReverseDNS)
}

func ConnRemote(conn net.Conn, useReverseDNS bool) string {
	return Addr(conn.RemoteAddr().String(), useReverseDNS)
}

func ConnLocal(conn net.Conn, useReverseDNS bool) string {
	return Addr(conn.LocalAddr().String(), useReverseDNS)
}

func Port[I types.IntOrUint](port I) string {
	return ansi.Wrap(fmt.Sprint(port), int(32.0+math.Max(0.0, math.Min(183.0, 183.0*(float64(port)/65535.0))))).String()
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
		res += ansi.Wrap(u.Scheme+ansi.Wrap("://", config.LoggerConfig.ColorURLSeparators).String(), config.LoggerConfig.ColorScheme).String()
		if u.User != nil && u.User.Username() != "" {
			res += ansi.Wrap(u.User.Username(), config.LoggerConfig.ColorUser).String()
			if p, ok := u.User.Password(); ok {
				res += ansi.Wrap(":", config.LoggerConfig.ColorURLSeparators).String() + ansi.Wrap(p, config.LoggerConfig.ColorPassword).String()
			}
			res += ansi.Wrap("@", config.LoggerConfig.ColorURLSeparators).String()
		}
		if config.LoggerConfig.CheckIfURLIsAlive {
			ips, _ := net.LookupIP(u.Host)
			if len(ips) > 0 {
				res += ansi.Wrap(u.Host, utils.Icc.Get(ips[0].To4().String())).String()
			} else {
				isAlive = false
				res += WrapRed(u.Host)
			}
		} else {
			res += ansi.Wrap(u.Host, utils.Scc.Get(u.Host)).String()
		}
		for i, pe := range strings.Split(u.Path, "/") {
			if pe == "" {
				continue
			}
			if i > 0 {
				res += ansi.Wrap("/", config.LoggerConfig.ColorURLSeparators).String()
			}
			res += ansi.Wrap(pe, config.LoggerConfig.ColorURLPath).String()
		}
		if len(u.Path) > 0 && string(u.Path[len(u.Path)-1]) == "/" {
			res += ansi.Wrap("/", config.LoggerConfig.ColorURLSeparators).String()
		}

		q := u.RawQuery
		if q != "" {
			res += ansi.Wrap("?", config.LoggerConfig.ColorURLSeparators).String()
			pairs := []string{}
			for _, pair := range strings.Split(q, "&") {
				if pair == "" {
					continue // empty pairs don't work
				}
				e := strings.Split(pair, "=")
				if len(e) == 1 {
					// we have a single key
					pairs = append(pairs, ansi.Wrap(e[0], config.LoggerConfig.ColorQueryKey).String())
				} else {
					// we have a key-value pair
					v := strings.Join(e[1:], "=")
					pairs = append(pairs,
						ansi.Wrap(e[0], config.LoggerConfig.ColorQueryKey).String()+
							ansi.Wrap("=", config.LoggerConfig.ColorURLSeparators).String()+
							ansi.Wrap(v, config.LoggerConfig.ColorQueryValue).String(),
					)
				}
			}
			res += strings.Join(pairs, ansi.Wrap("&", config.LoggerConfig.ColorURLSeparators).String())
		}

		if u.Fragment != "" {
			res += ansi.Wrap("#", config.LoggerConfig.ColorURLSeparators).String() + ansi.Wrap(u.Fragment, config.LoggerConfig.ColorFragment).String()
		}
		if !isAlive {
			res = WrapRed("💀 ") + res
		}
		out = append(out, res)
	}
	return strings.Join(out, ", ")
}
