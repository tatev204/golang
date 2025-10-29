package main

import (
	"fmt"
	"project10/input"
)

func main() {

	var n int
	fmt.Print("Enter n", n)
	fmt.Scan(&n)

	Y := input.InputArray(n)

	fmt.Println("New array:", Y)
}
