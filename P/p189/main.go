package main

import (
	"fmt"
)

func rotate(nums []int, k int) {
	n := len(nums)
	nums1 := make([]int, n)
	for i := 0; i < n; i++ {
		nums1[(i+k)%n] = nums[i]
	}
	copy(nums, nums1)
}

func main() {
	nums := []int{1, 2}
	k := 5
	rotate(nums, k)
	fmt.Println(nums)
	nums = []int{1, 2, 3, 4, 5, 6, 7}
	k = 3
	rotate(nums, k)
	fmt.Println(nums)
	nums = []int{-1, -100, 3, 99}
	k = 2
	rotate(nums, k)
	fmt.Println(nums)
}
