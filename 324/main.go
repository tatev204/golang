package main

import (
	"fmt"
	"project21/input"
)

func main() {
	var n int
	fmt.Print("enter n:")
	fmt.Scan(&n)

	X := input.InputArray(n)

	Y := input.LargeEqualSmallElements(X)
	fmt.Println(Y)

}
