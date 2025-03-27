package dining

import (
	"bytes"
	"os"
	"runtime"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAcquireRelease(t *testing.T) {
	f := NewFork()

	f.Acquire()
	f.Release()
}

func TestSequentialAcquireRelease(t *testing.T) {
	f := NewFork()

	f.Acquire()
	f.Release()

	f.Acquire()
	f.Release()
}

func TestNoSharedState(t *testing.T) {
	f1 := NewFork()
	f1.Acquire()

	f2 := NewFork()
	f2.Acquire()

	f2.Release()
	f1.Release()
}

func TestMutualExlusion(t *testing.T) {
	f := NewFork()
	l := false

	go func() {
		f.Acquire()
		l = true
		time.Sleep(time.Second * 3)
		l = false
		f.Release()
	}()

	time.Sleep(time.Second)
	f.Acquire()
	assert.False(t, l)
	f.Release()
}

func TestNoBusyWait(t *testing.T) {
	fork := NewFork()
	fork.Acquire()
	defer fork.Release()

	for i := 0; i < 100; i++ {
		go func() {
			fork.Acquire()
			defer fork.Release()
		}()
	}

	verifyNoBusyGoroutines(t)
}

func verifyNoBusyGoroutines(t *testing.T) {
	time.Sleep(time.Millisecond * 100)

	for i := 0; i < 100; i++ {
		time.Sleep(time.Millisecond)

		var stacks []byte
		for n := 1 << 20; true; n *= 2 {
			stacks = make([]byte, n)
			m := runtime.Stack(stacks, true)

			if m < n {
				stacks = stacks[:m]
				break
			}
		}

		busy := bytes.Count(stacks, []byte("[running]"))
		busy += bytes.Count(stacks, []byte("[runnable]"))
		busy += bytes.Count(stacks, []byte("[sleep]"))

		if !assert.Less(t, busy, 2) {
			_, _ = os.Stderr.Write(stacks)
			break
		}
	}

}
