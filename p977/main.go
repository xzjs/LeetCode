package main

import "fmt"
import "sort"

func sortedSquares(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		nums[i] *= nums[i]
	}
	sort.Ints(nums)
	return nums
}

func main() {
	nums := []int{-4, -1, 0, 3, 10}
	fmt.Println(sortedSquares(nums))
	nums = []int{-7, -3, 2, 3, 11}
	fmt.Println(sortedSquares(nums))
}
