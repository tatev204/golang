package artadryal

func Realnum(N int) []float64 {
	if N <= 3 {
		return nil
	}
	x := make([]float64, N+1)
	x[1] = 1
	x[2] = 2
	x[3] = 3
	for i := 4; i <= N; i++ {
		x[i] = x[i-2] + x[i-1] - 2*x[i-3]
	}

	return x
}
