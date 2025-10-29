package input

import (
	"fmt"
	"math"
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

func MinElement(X []int) int {
	min := X[0]

	for i := 1; i < len(X); i++ {
		if X[i] < min {
			min = X[i]

		}
	}
	return min
}

func MaxMinAverage(arr []int) float64 {
	max := MaxElement(arr)
	min := MinElement(arr)

	average := float64(max+min) / 2.0
	return average
}

func AbsouleValueareSmallerthanAverage(arr []int) []int {
	Y := make([]int, 0)
	average := MaxMinAverage(arr)
	for _, val := range arr {
		if math.Abs(float64(val)) < average {
			Y = append(Y, val)
		}
	}
	return Y
}
