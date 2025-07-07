package artadryal

func Sequence(N int) []float64 {
	x := make([]float64, N+1)
	x[0] = 2

	for i := 1; i <= N; i++ {
		x[i] = 2 + 1/x[i-1]
	}

	return x
}
