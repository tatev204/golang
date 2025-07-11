package main

import (
	"fmt"
	"project19/input"
)

func main() {
	var n int
	fmt.Print("enter n:")
	fmt.Scan(&n)

	X := input.InputArray(n)

	Y := input.YArray(X)
	fmt.Println(Y)

}
