package monitor

import (
	"fmt"
	"net/http"
	"time"
)

type Logger interface {
	Info(message string)
	Warn(message string)
	Error(message string, err error)
}

type Benchmark interface {
	Start()
	Measure() time.Duration
}

type MonitorConfig struct {
	// The target url to monitor.
	Url string `json:"url"`
	// The interval at which to run the monitor.
	Interval int `json:"interval"`
}

type Monitor struct {
	Logger    Logger
	Benchmark Benchmark
}

func New(logger Logger, benchmark Benchmark) *Monitor {
	return &Monitor{
		Logger:    logger,
		Benchmark: benchmark,
	}
}

// Start monitoring the provided urls.
func (monitor *Monitor) Start(configs []MonitorConfig) {
	const minFrequency = 1

	for _, config := range configs {
		if config.Url == "" {
			monitor.Logger.Warn("Config contains entry with empty Url, skipping")
			continue
		}

		frequency := config.Interval
		if frequency < minFrequency {
			monitor.Logger.Warn(
				fmt.Sprintf(
					"Interval provided (%d) for %s is invalid. Defaulting to minimum value (%d)",
					frequency,
					config.Url,
					minFrequency,
				),
			)
			frequency = minFrequency
		}

		go monitor.monitorUrl(config.Url, frequency)
	}
}

func (monitor *Monitor) monitorUrl(url string, frequency int) {
	time.AfterFunc(time.Duration(frequency)*time.Second, func() { monitor.monitorUrl(url, frequency) })

	monitor.Benchmark.Start()
	response, err := request(url)
	measurement := monitor.Benchmark.Measure()

	if err != nil {
		errorMessage := fmt.Sprintf("Error fetching: %s", url)
		monitor.Logger.Error(errorMessage, err)
	} else {
		logMessage := fmt.Sprintf("%s: %s, in %s", url, response.Status, measurement)
		monitor.Logger.Info(logMessage)
	}
}

func request(url string) (*http.Response, error) {
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}
	return resp, nil
}
