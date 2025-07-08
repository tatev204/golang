package input

import "fmt"

func CountLetter(arr []rune, letter rune) int {
	count := 0
	for _, val := range arr {
		if val == letter {
			count++
		}
	}
	return count
}

func InputArray(name string, n int) []rune {
	arr := make([]rune, n)
	fmt.Printf("Enter %d characters for %s:\n", n, name)
	for i := 0; i < n; i++ {
		fmt.Scanf(" %c", &arr[i])
	}
	return arr
}
