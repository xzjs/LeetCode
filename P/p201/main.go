package main

import "fmt"

func minSubArrayLen(target int, nums []int) int {
	sum := 0
	n := len(nums)
	minLength := n + 1
	for i, j := 0, 0; j < n; j++ {
		sum += nums[j]
		for sum >= target {
			minLength = min(minLength, j-i+1)
			sum -= nums[i]
			i++
		}
	}
	if minLength > len(nums) {
		return 0
	}
	return minLength
}

func main() {
	target := 7
	nums := []int{2, 3, 1, 2, 4, 3}
	fmt.Println(minSubArrayLen(target, nums))
}
