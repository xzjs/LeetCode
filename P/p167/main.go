package main

import "fmt"

func twoSum(numbers []int, target int) []int {
	mid := float64(target) / 2
	var index int
	for i := 0; i < len(numbers)-1; i++ {
		if float64(numbers[i]) == mid && float64(numbers[i+1]) == mid {
			return []int{i + 1, i + 2}
		}
		if float64(numbers[i]) < mid && float64(numbers[i+1]) > mid {
			index = i
			break
		}
	}
	for i, j := index, index+1; i >= 0 && j < len(numbers); {
		sum := numbers[i] + numbers[j]
		if sum == target {
			return []int{i + 1, j + 1}
		}
		if sum > target {
			i--
		}
		if sum < target {
			j++
		}
	}
	return []int{}
}

func main() {
	numbers := []int{-1000, -1, 0, 1}
	target := 1
	fmt.Println(twoSum(numbers, target))
}
