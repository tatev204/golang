package main

import (
	"fmt"
)

func RemoveDuplicates(n int, nums []int) []int {
	var uniqueIndices []int
	k := 0
	for i := 0; i < n-1; i++ {
		if nums[i] != nums[i+1] {
			k++
			uniqueIndices = append(uniqueIndices, i)
		}
	}
	uniqueIndices = append(uniqueIndices, n-1)
	return uniqueIndices
}
func Numbers(n int) []int {
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}
	return nums
}

func main() {
	var n int
	fmt.Println("Enter number of elements:")
	fmt.Scan(&n)

	nums := Numbers(n)
	uniqueIndices := RemoveDuplicates(n, nums)

	fmt.Println("uniqueIndices ", uniqueIndices)

}
