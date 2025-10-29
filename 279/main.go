package main

import (
	"fmt"
	"project9/input"
)

func main() {
	k := 'k'
	letters := input.FilterRunesGreaterThanK(k)

	fmt.Println("Letters greater than 'k':")
	for _, ch := range letters {
		fmt.Printf("%c ", ch)
	}
	fmt.Println()
}
