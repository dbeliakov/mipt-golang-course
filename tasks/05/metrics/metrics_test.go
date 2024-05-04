package metrics

import (
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
	"time"
)

// Mock time function for predictable test results
var fixedTime = time.Date(2020, 10, 10, 10, 10, 10, 0, time.UTC)
var timeShifts []time.Duration
var shiftIdx int
var timeMutex sync.Mutex

func mockNowFunc() time.Time {
	timeMutex.Lock()
	defer timeMutex.Unlock()
	if shiftIdx < len(timeShifts) {
		fixedTime = fixedTime.Add(timeShifts[shiftIdx])
		shiftIdx++
	}
	return fixedTime
}

func TestRecordAndReportDuration(t *testing.T) {
	nowFunc = mockNowFunc
	defer func() { nowFunc = time.Now }()

	timeShifts = []time.Duration{0, time.Second, 2 * time.Second}
	shiftIdx = 0

	measurer := NewDurationMeasurer(3)
	recordFunc := measurer.RecordDuration("testKey")
	recordFunc() // Ends after 1 second

	avgDuration := measurer.Report("testKey")
	// Expect average over 1 recorded duration
	assert.Equal(t, time.Second, avgDuration)
}

func TestDurationExceedingWindowSize(t *testing.T) {
	nowFunc = mockNowFunc
	defer func() { nowFunc = time.Now }()

	timeShifts = []time.Duration{0, time.Second, time.Second, 2 * time.Second, 2 * time.Second, 3 * time.Second, 3 * time.Second, 4 * time.Second, 4 * time.Second}
	shiftIdx = 0

	measurer := NewDurationMeasurer(3)
	measurer.RecordDuration("testKey")() // 1 second
	measurer.RecordDuration("testKey")() // 2 seconds
	measurer.RecordDuration("testKey")() // 3 seconds
	measurer.RecordDuration("testKey")() // 4 seconds, this should push out the oldest, which is 1 second

	avgDuration := measurer.Report("testKey")
	// Expect average over 3 durations: 2s, 3s, 4s
	expectedAvg := (2 + 3 + 4) * time.Second / 3
	assert.Equal(t, expectedAvg, avgDuration)
}

func TestReportWithNoRecordings(t *testing.T) {
	nowFunc = mockNowFunc
	defer func() { nowFunc = time.Now }()

	measurer := NewDurationMeasurer(3)
	assert.Zero(t, measurer.Report("nonExistentKey"))
}

func TestDurationsForDifferentKeys(t *testing.T) {
	nowFunc = mockNowFunc
	defer func() { nowFunc = time.Now }()

	timeShifts = []time.Duration{0, time.Second, time.Second, 2 * time.Second, 2 * time.Second, 3 * time.Second, 3 * time.Second, 4 * time.Second, 4 * time.Second, 5 * time.Second, 5 * time.Second}
	shiftIdx = 0

	measurer := NewDurationMeasurer(5)

	// Record durations for key1
	measurer.RecordDuration("key1")() // Ends after 1 second
	measurer.RecordDuration("key1")() // Ends after another 2 seconds (total 3 seconds)

	// Record durations for key2
	measurer.RecordDuration("key2")() // Ends after another second (total 4 seconds)
	measurer.RecordDuration("key2")() // Ends after another second (total 5 seconds)

	// Report average for key1 should be 2 seconds ([1s, 3s])
	avgDurationKey1 := measurer.Report("key1")
	expectedAvgKey1 := (1*time.Second + 2*time.Second) / 2
	assert.Equal(t, expectedAvgKey1, avgDurationKey1)

	// Report average for key2 should be 4.5 seconds ([4s, 5s])
	avgDurationKey2 := measurer.Report("key2")
	expectedAvgKey2 := (3*time.Second + 4*time.Second) / 2
	assert.Equal(t, expectedAvgKey2, avgDurationKey2)
}

func TestConcurrencySafety(t *testing.T) {
	nowFunc = mockNowFunc
	defer func() { nowFunc = time.Now }()

	timeShifts = []time.Duration{0, time.Millisecond, 2 * time.Millisecond, 3 * time.Millisecond}
	shiftIdx = 0

	measurer := NewDurationMeasurer(10)

	var wg sync.WaitGroup
	numRoutines := 50
	wg.Add(numRoutines * 2) // Each goroutine will record and then report

	// Start multiple goroutines that use the same key to record durations
	for i := 0; i < numRoutines; i++ {
		go func(n int) {
			defer wg.Done()
			measurer.RecordDuration("concurrentKey")()
		}(i)

		go func(n int) {
			defer wg.Done()
			measurer.Report("concurrentKey")
		}(i)
	}

	wg.Wait()
}
