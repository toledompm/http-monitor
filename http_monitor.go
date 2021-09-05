package main

import (
	"os"
	"runtime"

	"github.com/toledompm/http_monitor/internal/config"
	"github.com/toledompm/http_monitor/internal/monitor"
	"github.com/toledompm/http_monitor/pkg/logger"
)

func main() {
	configFilePath := os.Args[1]
	config, err := config.ReadConfig(configFilePath)
	if err != nil {
		logger.Error("Error opening config file", err)
		runtime.Goexit()
	}

	monitor.Start(config)
	runtime.Goexit()
}
