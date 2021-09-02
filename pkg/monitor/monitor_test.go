package monitor

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

type mockLogger struct {
	ExpectedUrl string
	t           *testing.T
}

func (*mockLogger) Error(message string, err error) {}
func (*mockLogger) Warn(message string)             {}
func (mockLogger *mockLogger) Info(message string) {
	if !strings.Contains(message, mockLogger.ExpectedUrl) {
		mockLogger.t.Errorf("Expected log message to contain %s, got %s", mockLogger.ExpectedUrl, message)
	}
}

type mockBenchmarker struct{}

func (*mockBenchmarker) Start() {}
func (*mockBenchmarker) Measure() time.Duration {
	return time.Duration(0)
}

func TestStart(t *testing.T) {
	const expectedRequestCount = 2
	const configInterval = 1
	requestCount := 0

	response := "{\"status\": \"200 OK\"}"
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestCount++
		fmt.Fprintln(w, response)
	}))

	config := []MonitorConfig{{
		Url:      server.URL,
		Interval: configInterval,
	}}

	mockLogger := &mockLogger{ExpectedUrl: server.URL, t: t}
	mockBenchmarker := &mockBenchmarker{}

	monitor := New(mockLogger, mockBenchmarker)
	monitor.Start(config)

	// wait for server to be hit expectedRequestCount times
	time.Sleep(time.Duration(expectedRequestCount*configInterval) * time.Second)
	server.Close()

	if requestCount != expectedRequestCount {
		t.Errorf("Expected %d requests to be sent, got %d", expectedRequestCount, requestCount)
	}
}
