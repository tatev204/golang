package main

import (
	"fmt"
	"project15/input"
)

func main() {
	var n int
	fmt.Print("enter n:")
	fmt.Scan(&n)

	X := input.InputArray(n)

	Y := input.ValueofPositiveElementGreatest(X)
	fmt.Println(Y)

}
