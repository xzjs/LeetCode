package main

import "fmt"

func search(nums []int, target int) int {
	head := 0
	tail := len(nums)
	for head < tail {
		middle := (head + tail) / 2
		if nums[middle] == target {
			return middle
		} else if nums[middle] > target {
			tail = middle
		} else {
			head = middle + 1
		}
	}
	return -1
}

func main() {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 9
	fmt.Println(search(nums, target))
	nums = []int{-1, 0, 3, 5, 9, 12}
	target = 2
	fmt.Println(search(nums, target))
}
