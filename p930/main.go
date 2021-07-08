package main

import "fmt"

func numSubarraysWithSum(nums []int, goal int) int {
	ans := 0
	cnt := map[int]int{}
	sum := 0
	for _, num := range nums {
		cnt[sum]++
		sum += num
		ans += cnt[sum-goal]
	}
	return ans
}

func main() {
	nums := []int{1, 0, 1, 0, 1}
	goal := 2
	fmt.Println(numSubarraysWithSum(nums, goal))
	nums = []int{0, 0, 0, 0, 0}
	goal = 0
	fmt.Println(numSubarraysWithSum(nums, goal))
}
