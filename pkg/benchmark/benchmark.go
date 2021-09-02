package benchmark

import "time"

type Benchmark struct {
	startTime time.Time
}

func New() *Benchmark {
	return &Benchmark{}
}

// Starts the benchmark.
func (b *Benchmark) Start() {
	b.startTime = time.Now()
}

// Returns the time elapsed since the benchmark was started.
func (b *Benchmark) Measure() time.Duration {
	return time.Since(b.startTime)
}
