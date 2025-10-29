package main

import (
	"fmt"
	"project22/input"
)

func main() {
	var n int
	fmt.Print("enter n:")
	fmt.Scan(&n)

	X := input.InputArray(n)

	Y := input.ChangingPositionsElements(X)
	fmt.Println(Y)

}
