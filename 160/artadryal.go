package artadryal

import "math"

func Min() int {
	for i := 100; i <= 999; i++ {
		result := i * 16
		n := int(math.Sqrt(float64(result)))

		if result == n*n {
			return n
		}
	}

	return 0
}
