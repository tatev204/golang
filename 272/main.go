package main

import (
	"fmt"
	"project4/input"
)

func main() {
	var n int
	fmt.Print("Enter number of characters: ")
	fmt.Scan(&n)

	A := input.InputArray("A", n)

	var letter rune
	fmt.Print("Enter letter to count: ")
	fmt.Scanf(" %c", &letter)

	count := input.CountLetter(A, letter)

	fmt.Printf("The letter '%c' appears %d times\n", letter, count)
}
