package mapcalc

import (
	"fmt"
	"testing"
)

func TestMap6to4(t *testing.T) {
	r := Rules{
		Maprule{
			Ipv6:   "2001:db8::/32",
			Ipv4:   "10.0.0.0/16",
			Ea_length: 24,
		},
		Maprule{
			Ipv6:   "2001:db9::/32",
			Ipv4:   "10.1.0.0/16",
			Ea_length: 24,
		},
		Maprule{
			Ipv6:   "2001:db10::/32",
			Ipv4:   "10.2.0.0/16",
			Ea_length: 24,
		},
	}
	addr, port := Map6to4("2001:db8:0101:ab00::1", r)
	fmt.Printf("Result: %s,%s\n", addr, port)
	if addr != "10.0.1.1" {
		t.Errorf("address is not correct :%s", addr)
	}
}

