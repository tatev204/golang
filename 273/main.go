package main

import (
	"fmt"
	"project5/input"
)

func main() {
	var n int
	fmt.Print("Enter number of characters: ")
	fmt.Scan(&n)

	s := input.InputArray("s", n)

	var letter rune
	fmt.Print("Enter letter to sum indexes: ")
	fmt.Scanf(" %c", &letter)

	sumIndexes := input.SumOfIndexesWithLetter(letter, s)

	fmt.Printf("The sum of indexes where '%c' appears is: %d\n", letter, sumIndexes)
}
