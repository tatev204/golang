package main

import (
	"fmt"
	"project23/input"
)

func main() {
	var n int
	fmt.Print("enter n:")
	fmt.Scan(&n)

	X := input.InputArray(n)

	Y := input.NonRepeatingElements(X)
	fmt.Println(Y)

}
