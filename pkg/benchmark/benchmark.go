package benchmark

import "time"

type Benchmark struct {
	startTime time.Time
}

func New() *Benchmark {
	return &Benchmark{}
}

func (b *Benchmark) Start() {
	b.startTime = time.Now()
}

func (b *Benchmark) Measure() time.Duration {
	return time.Since(b.startTime)
}
