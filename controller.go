package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

type Controller struct {
	Config     *Config
	CURR_SPEED int
}

func NewController() *Controller {
	c := Controller{Config: NewConfig()}
	c.Config.Load()
	return &c
}

func (c *Controller) SetSpeed(speed int) {
	c.CURR_SPEED = speed
	err := os.WriteFile("/sys/devices/pwm-fan/target_pwm", []byte(strconv.Itoa(speed)), 0644)
	if err != nil {
		panic(err)
	}
}

func (c *Controller) GetTemp() (temp int) {
	file, _ := os.ReadFile("/sys/devices/virtual/thermal/thermal_zone0/temp")
	temp, _ = strconv.Atoi(string(file))
	return
}

func (c *Controller) SetClockSpeedMax() {
	cmd := exec.Command("/usr/bin/jetson_clocks")
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}

func (c *Controller) SetClockSpeed() {
	if c.Config.MAX_CLOCK_SPEED {
		c.SetClockSpeedMax()
		fmt.Println("Enabled Jetson performance mode.")
	}
}

func (c *Controller) SetSpeedFromTemp() {
	temp := c.GetTemp()
	if temp < c.Config.FAN_OFF_TEMP {
		c.SetSpeed(0)
		return
	}

	curve := createCurve(c.Config.FAN_OFF_TEMP, c.Config.FAN_MAX_TEMP, 255)
	speed := curve(temp)
	c.SetSpeed(speed)
}

func square(x int) (y int) {
	y = x * x
	return
}

func createCurve(vertex int, x2 int, y2 int) (curve func(x int) (y int)) {
	curve = func(x int) (y int) {
		y = y2 / square(x2 - vertex) * square(x - vertex)
		return
	}
	return
}