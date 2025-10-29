package main

import (
	"fmt"
	"project8/input"
)

func main() {
	var n int
	fmt.Printf("Enter n", n)
	fmt.Scan(&n)

	Y := input.InputArray(n)
	fmt.Println("New array:", Y)

}
