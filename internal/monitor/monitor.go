package monitor

import (
	"fmt"
	"time"

	"github.com/toledompm/http_monitor/pkg/logger"
	"github.com/toledompm/http_monitor/pkg/trace"
)

type MonitorConfig struct {
	// The target url to monitor.
	Url string `json:"url"`
	// The interval at which to run the monitor.
	Interval int `json:"interval"`
}

// Start monitoring the provided urls.
func Start(configs []MonitorConfig) {
	const minFrequency = 1

	for _, config := range configs {
		if config.Url == "" {
			logger.Warn("Config contains entry with empty Url, skipping")
			continue
		}

		frequency := config.Interval
		if frequency < minFrequency {
			logger.Warn(
				fmt.Sprintf(
					"Interval provided (%d) for %s is invalid. Defaulting to minimum value (%d)",
					frequency,
					config.Url,
					minFrequency,
				),
			)
			frequency = minFrequency
		}

		go monitorUrl(config.Url, frequency)
	}
}

func monitorUrl(url string, frequency int) {
	time.AfterFunc(time.Duration(frequency)*time.Second, func() { monitorUrl(url, frequency) })

	traceResult, err := trace.TraceGetRequest(url)

	if err != nil {
		errorMessage := fmt.Sprintf("Error fetching: %s", url)
		logger.Error(errorMessage, err)
	} else {
		logMessage := formatTraceResult(url, traceResult)
		logger.Info(logMessage)
	}
}

func formatTraceResult(url string, traceResult *trace.TraceResult) string {
	return fmt.Sprintf(
		"%s: %s - %s\nDNS Lookup: %s\nConnection: %s\nTLS Handshake: %s\nTime to first byte: %s",
		url,
		traceResult.Response.Status,
		traceResult.TotalTime,
		traceResult.DnsTime,
		traceResult.ConnectionTime,
		traceResult.TlsTime,
		traceResult.FirstResponseByteTime,
	)
}
