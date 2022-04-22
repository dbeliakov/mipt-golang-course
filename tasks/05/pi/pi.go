package pi

func CalculatePi(concurrent, iterations int, gen RandomPointGenerator) float64 {
	return 0
}

type RandomPointGenerator interface {
	Next() (float64, float64)
}
