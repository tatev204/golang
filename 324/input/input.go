package input

import (
	"fmt"
)

func InputArray(n int) []int {
	X := make([]int, n)

	fmt.Printf("Enter %d numbers:\n", n)
	for i := 0; i < n; i++ {
		fmt.Scan(&X[i])
	}
	return X
}
func LargeEqualSmallElements(X []int) []int {
	Y := make([]int, 0)

	for _, val := range X {
		if val > X[0] {
			Y = append(Y, val)
		}
	}

	for _, val := range X {
		if val == X[0] {
			Y = append(Y, val)
		}
	}

	for _, val := range X {
		if val < X[0] {
			Y = append(Y, val)
		}
	}
	return Y
}
