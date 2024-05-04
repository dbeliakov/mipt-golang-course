package metrics

import (
	"time"
)

type DurationMeasurer struct {
}

func NewDurationMeasurer(windowSize int) *DurationMeasurer {
	return &DurationMeasurer{}
}

func (m *DurationMeasurer) RecordDuration(key string) func() {
	return func() {}
}

func (m *DurationMeasurer) Report(key string) time.Duration {
	return 0
}

var nowFunc = time.Now
