package main

import (
	"fmt"
	"project18/input"
)

func main() {
	var n int
	fmt.Print("enter n:")
	fmt.Scan(&n)

	X := input.InputArray(n)

	Y := input.ReorderBySign(X)
	fmt.Println(Y)

}
