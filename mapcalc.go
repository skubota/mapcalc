package mapcalc

import (
	"fmt"
	"net"
	"strconv"
	"strings"
)

type Rules []Maprule

type Maprule struct {
	Ipv6   string
	Ipv4   string
	Ea_len int
}

func ip6bin(addr string) string {
	var bin string
	ip := net.ParseIP(addr)
	for _, v := range ip.To16() {
		bin = bin + fmt.Sprintf("%08b", uint8(v))
	}
	return bin
}
func ip4bin(addr string) string {
	var bin string
	ip := net.ParseIP(addr)
	for _, v := range ip.To4() {
		bin = bin + fmt.Sprintf("%08b", uint8(v))
	}
	return bin
}

func Map6to4(from string, r Rules) (string, string) {
	from_bin := ip6bin(from)
	for _, v := range r {
		ipv6 := strings.Split(v.Ipv6, "/")
		ipv6_bin := ip6bin(ipv6[0])
		ipv6_len, _ := strconv.Atoi(ipv6[1])
		if from_bin[0:ipv6_len] == ipv6_bin[0:ipv6_len] {
			ipv4 := strings.Split(v.Ipv4, "/")
			ipv4_bin := ip4bin(ipv4[0])
			ipv4_len, _ := strconv.Atoi(ipv4[1])

			addr_bin := ipv4_bin[0:ipv4_len] + from_bin[ipv6_len:ipv6_len+32-ipv4_len]
			port_bin := from_bin[ipv6_len+32-ipv4_len : (ipv6_len+32-ipv4_len)+(v.Ea_len-(32-ipv4_len))]
			port_int, _ := strconv.ParseInt(port_bin, 2, 0)
		
			addr1, _ := strconv.ParseInt(addr_bin[0:8], 2, 0)
			addr2, _ := strconv.ParseInt(addr_bin[8:16], 2, 0)
			addr3, _ := strconv.ParseInt(addr_bin[16:24], 2, 0)
			addr4, _ := strconv.ParseInt(addr_bin[24:32], 2, 0)
			return fmt.Sprintf("%d.%d.%d.%d", addr1, addr2, addr3, addr4), fmt.Sprintf("0x%x", port_int)
		}
	}
	return "", ""
}

