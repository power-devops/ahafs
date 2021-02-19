package main

import (
	"fmt"

	"github.com/power-devops/ahafs"
)

func main() {
	if m, err := ahafs.NewDirMonitor("/tmp"); err != nil {
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
