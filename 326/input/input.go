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

func ChangingPositionsElements(X []int) []int {
	n := len(X)
	for i := 0; i < n/2; i++ {
		X[i], X[n-1-i] = X[n-1-i], X[i]

	}
	Y := append([]int{}, X...)
	return Y
}
