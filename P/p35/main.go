package main

import "fmt"

func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums)
	for left < right {
		middle := (left + right) / 2
		if nums[middle] == target {
			return middle
		} else if nums[middle] > target {
			right = middle
		} else {
			left = middle + 1
		}
	}
	return left
}

func main() {
	nums := []int{1, 3, 5, 6}
	target := 5
	fmt.Println(searchInsert(nums, target))
	target = 2
	fmt.Println(searchInsert(nums, target))
	target = 7
	fmt.Println(searchInsert(nums, target))
	target = 0
	fmt.Println(searchInsert(nums, target))
}
