package philosophers

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/dbeliakov/mipt-golang-course/tasks/05/dining"
	"github.com/stretchr/testify/assert"
)

const seatCount = 10

func TestDining(t *testing.T) {
	table := dining.NewTable(seatCount)

	philosophers := make([]*Philosopher, 0, seatCount)
	for i := 0; i < seatCount; i++ {
		philosophers = append(philosophers, NewPhilosopher(table, i))
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()
	var wg sync.WaitGroup

	for i := 0; i < seatCount; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()
			for c := 0; c < 100; c++ {
				select {
				case <-ctx.Done():
					return
				default:
					philosophers[i].Dine()
					philosophers[i].Think()
				}
			}
		}(i)
	}

	wg.Wait()
	for _, p := range philosophers {
		assert.Greater(t, p.EatsCount(), 0)
	}
}
