package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Print("enter n: ")
	fmt.Scan(&n)

	X := InputArray("X", n)
	Y := InputArray("Y", n)

	sumEvenX := sumEvenNumbers(X)
	sumEvenY := sumOddNumbers(Y)

	totalSum := sumEvenX + sumEvenY

	fmt.Println("only even in X:", sumEvenX)
	fmt.Println("only odd in Y:", sumEvenY)
	fmt.Println("total amount:", totalSum)
}
