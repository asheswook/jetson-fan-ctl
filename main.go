//go:build linux

package main

import (
	"fmt"
	"time"
)

func main() {
	c := NewController()
	fmt.Println("Initalized fan controller.")

	c.SetClockSpeedMax()

	fmt.Println("Started automatic fan control.")
	for {
		c.SetSpeedFromTemp()
		time.Sleep(time.Duration(c.Config.INTERVAL) * time.Second)
	}
}
