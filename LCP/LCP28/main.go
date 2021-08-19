package main

import (
	"fmt"
	"sort"
)

func purchasePlans(nums []int, target int) int {
	sort.Ints(nums)
	result, i, j := 0, 0, len(nums)-1
	for i < j {
		sum := nums[i] + nums[j]
		if sum > target {
			j--
		} else {
			result += (j - i)
			i++
		}
	}
	return result % 1000000007
}

func main() {
	nums := []int{2, 5, 3, 5}
	target := 6
	fmt.Println(purchasePlans(nums, target))
	nums = []int{2, 2, 1, 9}
	target = 6
	fmt.Println(purchasePlans(nums, target))
}
