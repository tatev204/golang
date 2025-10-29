package main

import (
	"fmt"
	"project11/input"
)

func main() {
	var n int
	fmt.Print("Enter n: ")
	fmt.Scan(&n)

	var p int
	fmt.Print("Enter p: ")
	fmt.Scan(&p)

	Y := input.InputArray(n, p)

	fmt.Println("New array divisible by p:", Y)
}
