package main

import (
	"github.com/docker/machine/libmachine/drivers/plugin"
	"github.com/mesaks/docker-machine-workstation/pkg/drivers/workstation"
)

func main() {
	plugin.RegisterDriver(workstation.NewDriver("", ""))
}
