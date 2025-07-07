package main

import (
	"fmt"
	"myproject/artadryal"
)

func main() {
	var n int
	fmt.Println("enter n")
	fmt.Scan(&n)

	result := artadryal.PrimeNum(n)
	fmt.Println(result)
}
