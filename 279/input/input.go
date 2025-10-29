package input

import (
	"fmt"
	"project9/input"
)

func InputArray(n int) []rune {
	arr := make([]rune, n)
	fmt.Println("enter letters")

	for i := 0; i < n; i++ {
		fmt.Scanf("%c", &arr[i])
	}
	fmt.Printf("You entered: %s\n", string(arr))
	return arr
}

func NewArrayWithoutD(d rune, arr []rune) []rune {
	newArr := []rune{}
	for _, val := range arr {
		if val != d {
			newArr = append(newArr, val)
		}
	}
	return newArr
}
