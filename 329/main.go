package main

import (
	"fmt"
	"project24/input"
)

func main() {
	var n int
	fmt.Println("enter n:")
	fmt.Scan(&n)

	X := input.InputArray(n)

	Y := input.SumoftheElementsPrecedingElement(X)
	fmt.Println(Y)

}
