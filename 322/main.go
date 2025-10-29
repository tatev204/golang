package main

import (
	"fmt"
	"project20/input"
)

func main() {
	var n int
	fmt.Print("enter n:")
	fmt.Scan(&n)

	X := input.InputArray(n)

	Y := input.TwoZeroAfterZerElement(X)
	fmt.Println(Y)

}
