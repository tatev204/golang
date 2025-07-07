package main

import "fmt"

func sumEvenNumbers(arr []int) int {
	sum := 0

	for _, val := range arr {
		if val%2 == 0 {
			sum += val
		}
	}
	return sum
}

func sumOddNumbers(arr []int) int {
	sum := 0

	for _, val := range arr {
		if val%2 != 0 {
			sum += val
		}
	}
	return sum

}

func InputArray(name string, n int) []int {
	arr := make([]int, n)
	fmt.Printf("Enter %d elements for %s array:\n", n, name)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	return arr

}
