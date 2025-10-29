package main

import (
	"fmt"
	"project16/input"
)

func main() {
	var n int
	fmt.Print("enter n:")
	fmt.Scan(&n)

	X := input.InputArray(n)

	Y := input.ZeroAfterPositiveElement(X)
	fmt.Println(Y)

}
