package main

import (
	"fmt"
	"project1/main1"
)

func SumArray(arr []int) int {
	sum := 0
	for _, val := range arr {
		sum += val * val
	}
	return sum
}

func main() {
	X := main1.InputArray(5)
	Y := main1.InputArray(10)

	sum := SumArray(X) + SumArray(Y)
	fmt.Println("Sum:", sum)
}
