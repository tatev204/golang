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
func FirstMaxElements(X []int) int {
	maxIndex := 0
	max := X[0]
	for i, val := range X {
		if val > max {
			max = val
			maxIndex = i

		}
	}
	return maxIndex
}

func LastMinElements(X []int) int {
	minIndex := 0
	min := X[0]

	for i := len(X) - 1; i >= 0; i-- {
		if X[i] < min {
			min = X[i]
			minIndex = i
		}
	}
	return minIndex

}

func YArray(X []int) []int {
	Y := make([]int, 0)

	maxIndex := FirstMaxElements(X)
	minIndex := LastMinElements(X)

	for i, val := range X {
		if i != maxIndex && i != minIndex {
			Y = append(Y, val)
		}
	}
	return Y
}
