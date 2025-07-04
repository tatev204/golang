package artadryal

import (
	"fmt"
)

func NaturalNumber() (int, bool) {
	var n int
	fmt.Print("enter n: ")
	fmt.Scan(&n)
	if n > 0 {
		return n, true
	}
	return n, false
}
