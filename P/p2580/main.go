package main

import (
	"fmt"
	"math"
	"sort"
)

func countWays(ranges [][]int) int {
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i][0] < ranges[j][0]
	})
	res := [][]int{ranges[0]}
	for i := 1; i < len(ranges); i++ {
		n := len(res)
		if ranges[i][1] > res[n-1][1] {
			if ranges[i][0] > res[n-1][1] {
				res = append(res, ranges[i])
			} else {
				res[n-1][1] = ranges[i][1]
			}
		}
	}
	return int(int64(math.Pow(2, float64(len(res)))) % int64(1e9+7))
}

func main() {
	ranges := [][]int{{6, 10}, {5, 15}}
	fmt.Println(countWays(ranges))
}
