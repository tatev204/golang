package main

import (
	"fmt"
	"myproject/artadryal"
)

func main() {
	var n int
	fmt.Println("enter n")
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("enter positive number")
		return
	}

	result := artadryal.Min(n)
	if result != 0 {
		fmt.Println(n, result)
	} else {
		fmt.Println("enter again")
	}
}
