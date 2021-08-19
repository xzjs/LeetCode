package main

import (
	"fmt"
	"sort"
)

func shipWithinDays(weights []int, D int) int {
	right := 0
	left := 0
	for _, value := range weights {
		right += value
		if value > left {
			left = value
		}
	}
	return left + sort.Search(right-left, func(x int) bool {
		temp := 0 //一天的装载量
		day := 0  //天数
		x += left
		for _, value := range weights {
			if temp+value > x {
				day++
				if day >= D {
					return false
				}
				temp = 0
			}
			temp += value
		}
		return true
	})
}

func main() {
	weights := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	D := 5
	fmt.Println(shipWithinDays(weights, D))
	weights = []int{3, 2, 2, 4, 1, 4}
	D = 3
	fmt.Println(shipWithinDays(weights, D))
	weights = []int{1, 2, 3, 1, 1}
	D = 4
	fmt.Println(shipWithinDays(weights, D))

}
