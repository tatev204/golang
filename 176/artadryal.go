package artadryal

func Fibonacci(N int) []int {
	if N <= 1 {
		return nil
	}
	x := make([]int, N+1)
	x[1] = 1
	x[2] = 1
	for i := 3; i <= N; i++ {
		x[i] = x[i-2] + x[i-1]
	}

	return x
}
