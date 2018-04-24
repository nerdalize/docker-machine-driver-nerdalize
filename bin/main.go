package main

import (
	"github.com/docker/machine/libmachine/drivers/plugin"
	"github.com/nerdalize/docker-machine-driver-nerdalize"
)

func main() {
	plugin.RegisterDriver(nerdalize.NewDriver("", ""))
}
