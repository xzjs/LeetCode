package main

import "fmt"

func twoSum(nums []int, target int) []int {
	hashTable := map[int]int{}
	for i, x := range nums {
		if p, ok := hashTable[target-x]; ok {
			return []int{p, i}
		}
		hashTable[x] = i
	}
	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))
	nums = []int{3, 2, 4}
	target = 6
	fmt.Println(twoSum(nums, target))
	nums = []int{3, 3}
	target = 6
	fmt.Println(twoSum(nums, target))
}
