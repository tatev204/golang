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
func SquaresOfElementsRangecdc(arr []int, c int, d int) []int {
	Y := make([]int, 0)

	for _, val := range arr {
		square := val * val
		if square > c && square < d {
			Y = append(Y, val)
		}
	}
	return Y
}
