package artadryal

import "math"

func Min(n int) int {

	for i := 100; i <= 999; i++ {
		result := math.Sqrt(float64(i))

		if result > float64(n) {
			return i
		}
	}

	return 0
}
