package main

import (
	"fmt"
	"myproject/artadryal"
)

func main() {
	result := artadryal.Min()
	if result != 0 {
		fmt.Println(result)
	} else {
		fmt.Println("not found")
	}
}
