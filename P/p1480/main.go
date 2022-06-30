package main

import "fmt"

func runningSum(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		if i != 0 {
			nums[i] += nums[i-1]
		}
	}
	return nums
}

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(runningSum(nums))
	nums = []int{1, 1, 1, 1, 1}
	fmt.Println(runningSum(nums))
	nums = []int{3, 1, 2, 10, 1}
	fmt.Println(runningSum(nums))
}
