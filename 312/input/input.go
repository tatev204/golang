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
func MaxElement(X []int) int {
	max := X[0]
	for i := 1; i < len(X); i++ {
		if X[i] > max {
			max = X[i]
		}
	}
	return max
}
func ValueofPositiveElementGreatest(arr []int) []int {
	Y := make([]int, 0)
	max := MaxElement(arr)
	for _, val := range arr {
		if val > 0 {
			val += max
		}
		Y = append(Y, val)
	}
	return Y
}
