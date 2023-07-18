package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

type Controller struct {
	Config *Config
	CURR_SPEED int
}

func NewController() *Controller {
	c := Controller{Config: NewConfig()}
	c.Config.Load()
	return &c
}

