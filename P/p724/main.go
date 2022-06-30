package main

import "fmt"

func pivotIndex(nums []int) int {
	left := 0
	right := 0
	for i := 0; i < len(nums); i++ {
		right += nums[i]
	}
	for i := 0; i < len(nums); i++ {
		right -= nums[i]
		if left == right {
			return i
		}
		left += nums[i]
	}
	return -1
}

func main() {
	nums := []int{1, 7, 3, 6, 5, 6}
	fmt.Println(pivotIndex(nums))
	nums = []int{1, 2, 3}
	fmt.Println(pivotIndex(nums))
	nums = []int{2, 1, -1}
	fmt.Println(pivotIndex(nums))
}
