package main

import "fmt"

func moveZeroes(nums []int) {
	for k, v := range nums {
		if v == 0 {
			for i := k + 1; i < len(nums); i++ {
				if nums[i] != 0 {
					nums[k], nums[i] = nums[i], nums[k]
					break
				}
			}
		}
	}

}

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	fmt.Println(nums)
}
