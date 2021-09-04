package trace

import "time"

type timer struct {
	startTime time.Time
	stopTime  time.Time
}

func newTimer() *timer {
	return &timer{}
}

// Starts the benchmark.
func (t *timer) Start() {
	t.startTime = time.Now()
}

// Stops the benchmark.
func (t *timer) Stop() {
	t.stopTime = time.Now()
}

// Returns the time difference between start and stop.
func (t *timer) Measure() time.Duration {
	return t.stopTime.Sub(t.startTime)
}
