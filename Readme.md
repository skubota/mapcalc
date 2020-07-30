Package mapcalc is map-e/t calculator
===================

DESCRIPTION

 MAP is RFC7597 mechanism
	Mapping of Address and Port with Encapsulation (MAP-E)
	https://tools.ietf.org/html/rfc7597

[![GoDev][godev-image]][godev-url]
![Go](https://github.com/skubota/mapcalc/workflows/Go/badge.svg)

[godev-image]: https://pkg.go.dev/badge/github.com/skubota/mapcalc
[godev-url]: https://pkg.go.dev/github.com/skubota/mapcalc


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
        r := mapcalc.Rules{
                mapcalc.Maprule{
                        Ipv6:   "2001:db8::/32",
                        Ipv4:   "10.0.0.0/16",
                        Ea_length: 24,
                },
                mapcalc.Maprule{
                        Ipv6:   "2001:db9::/32",
                        Ipv4:   "10.1.0.0/16",
                        Ea_length: 24,
                },
                mapcalc.Maprule{
                        Ipv6:   "2001:db10::/32",
                        Ipv4:   "10.2.0.0/16",
                        Ea_length: 24,
                },
        }
        addr, port := mapcalc.Map6to4("2001:db8:0101:0a00::1", r)
        fmt.Printf("Result: %s,%s\n", addr, port)
}
```

