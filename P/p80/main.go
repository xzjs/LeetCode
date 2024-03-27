package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) == 1 {
		return 1
	}
	i := 0
	for j, k := 0, 0; j < len(nums); j++ {
		if nums[j] != nums[i] {
			if j-i >= 2 && j < len(nums)-1 {
				k = j + 1
				for k < len(nums) && nums[k] == nums[j] {
					k++
				}
				if k < len(nums) {
					i, j = j, k
				}
			}
			nums[i] = nums[j]
		}
	}
	return i + 1
}

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	fmt.Println(removeDuplicates(nums))
}
