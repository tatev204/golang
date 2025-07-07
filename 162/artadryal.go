package artadryal

import "math"

func Max() int {
	for i := 9999; i >= 1000; i-- {
		result := i * 14
		n := int(math.Sqrt(float64(result)))

		if result == n*n {
			return i
		}
	}

	return 0
}
