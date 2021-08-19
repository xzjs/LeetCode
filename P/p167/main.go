package main

import "fmt"

func twoSum(numbers []int, target int) []int {
	dict := make(map[int]int, len(numbers))
	for k, v := range numbers {
		if dict[v] != 0 {
			return []int{dict[v], k + 1}
		}
		dict[target-v] = k + 1
	}
	return []int{}
}

func main() {
	numbers := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(numbers, target))
	numbers = []int{2, 3, 4}
	target = 6
	fmt.Println(twoSum(numbers, target))
	numbers = []int{-1, 0}
	target = -1
	fmt.Println(twoSum(numbers, target))
}
