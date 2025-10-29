package main

import (
	"fmt"
	"project17/input"
)

func main() {
	var n int
	fmt.Print("enter n:")
	fmt.Scan(&n)

	X := input.InputArray(n)

	Y := input.AbsouleValueareSmallerthanAverage(X)
	fmt.Println(Y)

}
