package philosophers

import (
	"math/rand/v2"
	"time"

	"github.com/dbeliakov/mipt-golang-course/tasks/05/dining"
)

type Philosopher struct {
	seat      int
	table     *dining.Table
	l, r      *dining.Fork
	eatsCount int
}

func NewPhilosopher(table *dining.Table, seat int) *Philosopher {
	l := table.LeftFork(seat)
	r := table.RightFork(seat)

	return &Philosopher{
		seat:  seat,
		table: table,
		l:     l,
		r:     r,
	}
}

func (p *Philosopher) EatsCount() int {
	return p.eatsCount
}

func (p *Philosopher) Dine() {
	p.AcquireForks()
	defer p.ReleaseForks()
	p.Eat()
}

func (p *Philosopher) Think() {
	sleepTime := 500 + rand.Int()%500
	time.Sleep(time.Duration(sleepTime) * time.Millisecond)
}

func (p *Philosopher) Eat() {
	p.eatsCount++
}

func (p *Philosopher) AcquireForks() {
	panic("implement me")
}

func (p *Philosopher) ReleaseForks() {
	panic("implement me")
}
