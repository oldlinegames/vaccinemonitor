package main

import (
	"github.com/oldlinegames/vaccinemonitor/input"
	"github.com/oldlinegames/vaccinemonitor/monitor"
)

func main() {
	m := input.CollectInput()

	monitor.StartEngine(m)
}
