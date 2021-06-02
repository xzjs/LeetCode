package main

import "fmt"

func checkSubarraySum(nums []int, k int) bool {
	n := len(nums)
	if n < 2 {
		return false
	}
	mp := map[int]int{0: -1}
	remainder := 0
	for i, num := range nums {
		remainder = (remainder + num) % k
		if pre, has := mp[remainder]; has {
			if i-pre >= 2 {
				return true
			}
		} else {
			mp[remainder] = i
		}
	}
	return false
}

func main() {
	nums := []int{23, 2, 4, 6, 7}
	k := 6
	fmt.Println(checkSubarraySum(nums, k))
	nums = []int{23, 2, 6, 4, 7}
	k = 6
	fmt.Println(checkSubarraySum(nums, k))
	nums = []int{23, 2, 6, 4, 7}
	k = 13
	fmt.Println(checkSubarraySum(nums, k))
}
