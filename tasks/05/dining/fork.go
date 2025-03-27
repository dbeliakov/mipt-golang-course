package dining

type Fork struct {
	// your code here
}

func NewFork() *Fork {
	return &Fork{
		// your code here
	}
}

func (f *Fork) Acquire() {
	panic("implement me")
}

func (f *Fork) Release() {
	panic("implement me")
}

type Table struct {
	seats int
	forks []*Fork
}

func NewTable(seats int) *Table {
	forks := make([]*Fork, 0, seats)
	for i := 0; i < seats; i++ {
		forks = append(forks, NewFork())
	}

	return &Table{
		seats: seats,
		forks: forks,
	}
}

func (t *Table) SeatsCount() int {
	return t.seats
}

func (t *Table) LeftForkIdx(seatNum int) int {
	return seatNum
}

func (t *Table) RightForkIdx(seatNum int) int {
	rightIdx := seatNum - 1
	if rightIdx < 0 {
		rightIdx = t.seats - 1
	}

	return rightIdx
}

func (t *Table) LeftFork(seatNum int) *Fork {
	return t.forks[t.LeftForkIdx(seatNum)]
}

func (t *Table) RightFork(seatNum int) *Fork {
	return t.forks[t.RightForkIdx(seatNum)]
}
