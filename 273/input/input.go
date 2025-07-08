package input

import (
	"fmt"
)

func InputArray(name string, n int) []rune {
	arr := make([]rune, n)
	fmt.Printf("Enter %d characters for %s:\n", n, name)
	for i := 0; i < n; i++ {
		fmt.Scanf(" %c", &arr[i])
	}
	return arr
}

func SumOfIndexesWithLetter(letter rune, arr []rune) int {
	sum := 0
	for i, val := range arr {
		if val == letter {
			sum += i
		}
	}
	return sum
}
