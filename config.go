package main

import (
	"encoding/json"
	"os"
)

type Config struct {
	FAN_OFF_TEMP    int  `json:"FAN_OFF_TEMP"`
	FAN_MAX_TEMP    int  `json:"FAN_MAX_TEMP"`
	INTERVAL        int  `json:"INTERVAL"`
	MAX_CLOCK_SPEED bool `json:"MAX_CLOCK_SPEED"`
}

func (c *Config) Load(s ...string) {
	var path string
	switch len(s) {
	case 0:
		path = "/etc/jetson-fan-ctl.conf"
	case 1:
		path = s[0]
	default:
		panic("Too many arguments")
	}

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(c)
	if err != nil {
		panic(err)
	}

	if c.FAN_OFF_TEMP > c.FAN_MAX_TEMP {
		panic("FAN_OFF_TEMP must be less than FAN_MAX_TEMP")
	}

	if c.INTERVAL < 1 {
		panic("INTERVAL must be greater than 0")
	}

	if c.INTERVAL > 60 {
		panic("INTERVAL must be less than 60")
	}
}

func NewConfig() *Config {
	c := Config{FAN_OFF_TEMP: 40, FAN_MAX_TEMP: 70, INTERVAL: 5, MAX_CLOCK_SPEED: true}
	return &c
}
