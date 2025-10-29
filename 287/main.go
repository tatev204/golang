package main

import (
	"fmt"
	"project12/input"
)

func main() {
	var n int
	fmt.Print("Enter n: ")
	fmt.Scan(&n)

	Y := input.InputArray(n)

	fmt.Println("New array :", Y)
}
