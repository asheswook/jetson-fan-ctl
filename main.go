//go:build linux

package main

import (
	"fmt"
	"time"
)

func main() {
	c := NewController()
	fmt.Println("Initialized fan controller.")

	go c.SetClockSpeed()

	fmt.Println("Started automatic fan control.")
	for {
		c.SetSpeedFromTemp()
		time.Sleep(time.Duration(c.Config.INTERVAL) * time.Second)
	}
}
