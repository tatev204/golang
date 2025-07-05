package artadryal

func Factorial(N int) int {
	result := 1

	for i := N; i >= 1; i -= 2 {
		result *= i
	}

	return result
}
