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

func (c *Controller) GetTemp() (temp float32) {
	file, _ := os.ReadFile("/sys/devices/virtual/thermal/thermal_zone0/temp")
	tempStr := string(file)[:2] + "." + string(file)[2:]
	f, _ := strconv.ParseFloat(tempStr, 32)
	temp = float32(f)
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
	if temp < float32(c.Config.FAN_OFF_TEMP) {
		c.SetSpeed(0)
		return
	}

	curve := createCurve(float32(c.Config.FAN_OFF_TEMP), float32(c.Config.FAN_MAX_TEMP), 255.0)
	speed := curve(temp)
	c.SetSpeed(int(speed))
}

func square(x float32) (y float32) {
	y = x * x
	return
}

func createCurve(vertex float32, x2 float32, y2 float32) (curve func(x float32) (y float32)) {
	curve = func(x float32) (y float32) {
		y = y2 / square(x2-vertex) * square(x-vertex)
		return
	}
	return
}