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

func SumoftheElementsPrecedingElement(X []int) []int {
	Y := make([]int, 0)

	for i, _ := range X {
		if i == 0 {
			Y = append(Y, X[i])
		} else {
			X[i] = X[i] + X[i-1]
			Y = append(Y, X[i])
		}
	}
	return Y
}
