package main

import (
	"fmt"
	"project6/input"
)

func main() {
	var n int

	fmt.Print("enter number of letters ")
	fmt.Scan(&n)

	arr := input.InputArray("Array", n)

	var k rune = 'k'
	fmt.Println(" letter k:")
	fmt.Printf("%c\n", k)

	count := input.CountNumberElementValueLessK(arr, k)
	fmt.Printf("'%c' has numeric value %d, and %d letters are less than it\n", k, k, count)

}
