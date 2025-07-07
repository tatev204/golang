package main

import (
	"fmt"
	"myproject/artadryal"
)

func main() {
	var N int
	fmt.Println("enter N")
	fmt.Scan(&N)

	result := artadryal.Fibonacci(N)
	fmt.Println(result)
}
