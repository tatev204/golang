package input

import (
	"fmt"
)

func InputArray(n int) []int {
	X := make([]int, n)
	Y := make([]int, 0)

	fmt.Printf("Enter %d numbers:\n", n)
	for i := 0; i < n; i++ {
		fmt.Scan(&X[i])
	}

	for i, val := range X {
		if val != i {
			Y = append(Y, val)
		}
	}
	return Y
}
