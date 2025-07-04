package main

import (
	"fmt"
	"myproject/artadryal"
)

func main() {
	n, _ := artadryal.NaturalNumber()

	product := 1

	for i := 1; i <= n; i++ {
		if n%i == 3 {
			product *= i
		}
	}

	fmt.Println(product)
}
