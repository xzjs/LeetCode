package main

import (
	"fmt"
	"sort"
)

func triangleNumber(nums []int) int {
	ans := 0
	sort.Ints(nums)
	for i, v := range nums {
		for j := i + 1; j < len(nums); j++ {
			ans += sort.SearchInts(nums[j+1:], v+nums[j])
		}
	}
	return ans
}

func main() {
	nums := []int{2, 2, 3, 4}
	fmt.Println(triangleNumber(nums))
	nums = []int{4, 2, 3, 4}
	fmt.Println(triangleNumber(nums))
	nums = []int{24, 3, 82, 22, 35, 84, 19}
	fmt.Println(triangleNumber(nums))
}
