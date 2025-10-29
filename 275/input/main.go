package input

import "fmt"

func InputArray(name string, n int) []rune {
	arr := make([]rune, n)
	fmt.Printf("Enter %d letters for %s:\n", n, name)
	for i := 0; i < n; i++ {
		fmt.Scanf("%c", &arr[i])
	}
	return arr
}

func CountNumberElementValueLessK(arr []rune, k rune) int {
	count := 0
	for _, val := range arr {
		if val < k {
			count++
		}
	}
	return count
}
