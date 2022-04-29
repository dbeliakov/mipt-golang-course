package pi

import (
	"math"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/goleak"
)

const (
	iterations = 59274
	expected   = 3.132233356952458
)

func TestCalculatePi(t *testing.T) {
	defer goleak.VerifyNone(t)

	for i := 1; i <= 5; i++ {
		g := NewTestGenerator(iterations)
		res := CalculatePi(i, iterations, g)
		require.True(t, math.Abs(res-expected) < 1e-4)
	}
}

type TestGenerator struct {
	t      *testing.T
	points chan []float64
}

func NewTestGenerator(count int) *TestGenerator {
	g := &TestGenerator{
		points: make(chan []float64, count),
	}
	r := rand.New(rand.NewSource(42))
	for i := 0; i < count; i++ {
		g.points <- []float64{r.Float64(), r.Float64()}
	}
	return g
}

func (g *TestGenerator) Next() (float64, float64) {
	select {
	case next := <-g.points:
		return next[0], next[1]
	default:
		require.Fail(g.t, "Count of iterations is more than needed")
	}
	return 0, 0
}
