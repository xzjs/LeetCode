package main

func twoSum(nums []int, target int) []int {
	i, j := 0, len(nums)-1
	for i < j {
		sum := nums[i] + nums[j]
		if sum < target {
			i++
		} else if sum > target {
			j--
		} else {
			return []int{nums[i], nums[j]}
		}
	}
	return []int{}
}
