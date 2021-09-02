package main

import (
	"runtime"

	Benchmark "github.com/toledompm/http_monitor/pkg/benchmark"
	Logger "github.com/toledompm/http_monitor/pkg/logger"
	Monitor "github.com/toledompm/http_monitor/pkg/monitor"

	Config "github.com/toledompm/http_monitor/internal/config"
)

func main() {
	benchmark := Benchmark.New()
	logger := Logger.New()

	monitor := Monitor.New(logger, benchmark)

	config, err := Config.ReadConfig()
	if err != nil {
		logger.Error("Error opening config file", err)
		runtime.Goexit()
	}

	monitor.Start(config)
	runtime.Goexit()
}
