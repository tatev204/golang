package main

import (
	"fmt"
	"project13/input"
)

func main() {
	var n, c, d int
	fmt.Print("enter n:")
	fmt.Scan(&n)

	X := input.InputArray(n)

	fmt.Println("enter c:")
	fmt.Scan(&c)

	fmt.Println("enter d:")
	fmt.Scan(&d)

	Y := input.SquaresOfElementsRangecdc(X, c, d)
	fmt.Println(Y)

}
