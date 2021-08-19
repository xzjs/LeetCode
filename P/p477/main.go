package main

import (
	"fmt"
)

func totalHammingDistance(nums []int) int {
	sum := 0
	n := len(nums)
	for i := 0; i < 30; i++ {
		c := 0
		for _, val := range nums {
			c += val >> i & 1
		}
		sum += c * (n - c)
	}
	return sum
}

func main() {
	nums := []int{4, 14, 2}
	fmt.Println(totalHammingDistance(nums))
}
