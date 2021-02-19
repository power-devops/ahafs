# AHAFS Go Interface

[![Go Reference](https://pkg.go.dev/badge/github.com/power-devops/ahafs.svg)](https://pkg.go.dev/github.com/power-devops/ahafs)
[![Go Report Card](https://goreportcard.com/badge/github.com/power-devops/ahafs)](https://goreportcard.com/report/github.com/power-devops/ahafs)

AHAFS is a AIX specific interface to track changes on filesystems.

You can read more about AHAFS in [AIX 7.2 Knowledge Center](https://www.ibm.com/support/knowledgecenter/en/ssw_aix_72/osmanagement/aix_ev.html).

Some useful examples programming AHAFS with Perl and C you can find on your AIX box in /usr/samples/ahafs.

The package can be compiled on any platform, but it works only on IBM AIX. On all other platforms you'll get error "not AIX".

As for now just modFile, modFileAttr and modDirectory are known to work. All other monitors are not tested.

## Usage

```go
package main

import (
	"fmt"

	"github.com/power-devops/ahafs"
)

func main() {
	if m, err := ahafs.NewMonitor("/aha/fs/modFile.monFactory/etc/passwd.mon", ""); err != nil {
		fmt.Println("ERROR:", err)
	} else {
		c := make(chan ahafs.Event)
		go m.Watch(c)
		for {
			// waiting for event
			e := <-c
			fmt.Printf("%v\n", e)
		}
	}
}
```

## Examples

More usage examples you can find in the directory [examples](https://github.com/power-devops/ahafs/tree/main/examples).

## Contibuting

Contributions are welcome! Fork and create a pull request.

