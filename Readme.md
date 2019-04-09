Package mapcalc is map-e/t calculator
===================

DESCRIPTION

 MAP is RFC7597 mechanism
	Mapping of Address and Port with Encapsulation (MAP-E)
	https://tools.ietf.org/html/rfc7597

usage

	$ go get github.com/skubota/mapcalc

godoc

<https://godoc.org/github.com/skubota/mapcalc>

sample

```go
package main

import (
	"fmt"
	"github.com/skubota/mapcalc"
)

func main() {
        r := Rules{
                Maprule{
                        Ipv6:   "2001:db8::/32",
                        Ipv4:   "10.0.0.0/16",
                        Ea_len: 24,
                },
                Maprule{
                        Ipv6:   "2001:db9::/32",
                        Ipv4:   "10.1.0.0/16",
                        Ea_len: 24,
                },
                Maprule{
                        Ipv6:   "2001:db10::/32",
                        Ipv4:   "10.2.0.0/16",
                        Ea_len: 24,
                },
        }
        addr, port := Map6to4("2001:db8:0101:ab00::1", r)
        fmt.Printf("Result: %s,%s\n", addr, port)
}
```

