package main

import (
	"runtime"

	Config "github.com/toledompm/http_monitor/internal/config"
	Monitor "github.com/toledompm/http_monitor/internal/monitor"
	"github.com/toledompm/http_monitor/pkg/logger"
)

func main() {
	config, err := Config.ReadConfig()
	if err != nil {
		logger.Error("Error opening config file", err)
		runtime.Goexit()
	}

	Monitor.Start(config)
	runtime.Goexit()
}
