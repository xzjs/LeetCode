package main

import "fmt"

func firstDayBeenInAllRooms(nextVisit []int) int {
	mod := int64(1e9 + 7)
	n := len(nextVisit)
	dp := make([]int64, n)
	dp[0] = 0
	for i := 1; i < n; i++ {
		to := nextVisit[i-1]
		dp[i] = dp[i-1]*2 - dp[to] + 2
		dp[i] %= mod
	}
	return int(dp[n-1])
}

func main() {
	nextVisit := []int{0, 1, 2, 0}
	fmt.Println(firstDayBeenInAllRooms(nextVisit))
}
