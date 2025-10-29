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

func TwoZeroAfterZerElement(arr []int) []int {
	Y := make([]int, 0)
	for _, val := range arr {
		Y = append(Y, val)
		if val == 0 {
			Y = append(Y, 0, 0)
		}
	}
	return Y
}
