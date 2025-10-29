package main

import (
	"fmt"
	"project7/input"
)

func main() {
	var n int
	fmt.Printf("enter n:")
	fmt.Scan(&n)

	arr := input.InputArray(n)
	d := 'd'

	newArr := input.NewArrayWithoutD(d, arr)
	fmt.Println("New array", string(newArr))

	for _, ch := range newArr {
		fmt.Printf("%c", ch)
	}
	fmt.Println()
}
