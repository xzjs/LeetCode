package main

import (
	"fmt"
)

func countPairs(deliciousness []int) int {
	ans := 0
	maxVal := deliciousness[0]
	for _, val := range deliciousness[1:] {
		if val > maxVal {
			maxVal = val
		}
	}
	maxSum := maxVal * 2
	cnt := map[int]int{}
	for _, val := range deliciousness {
		for sum := 1; sum <= maxSum; sum <<= 1 {
			ans += cnt[sum-val]
		}
		cnt[val]++
	}
	return ans % (1e9 + 7)
}

func main() {
	deliciousness := []int{1, 3, 5, 7, 9}
	fmt.Println(countPairs(deliciousness))
	deliciousness = []int{1, 1, 1, 3, 3, 3, 7}
	fmt.Println(countPairs(deliciousness))
}
