package input

import (
	"fmt"
)

func InputArray(n int) []int {
	X := make([]int, n)
	Y := make([]int, 0)

	fmt.Printf("Enter numbers:\n", n)
	for i := 0; i < n; i++ {
		fmt.Scan(&X[i])
	}
	for _, val := range X {
		if val%2 != 0 {
			Y = append(Y, val)
		}
	}
	return Y
}
