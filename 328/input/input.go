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

func NonRepeatingElements(X []int) []int {
	Y := make([]int, 0)

	for i := 0; i < len(X); i++ {
		count := 0
		for j := 0; j < len(X); j++ {
			if X[i] == X[j] {
				count++
			}
		}
		if count == 1 {
			Y = append(Y, X[i])
		}
	}

	return Y
}
