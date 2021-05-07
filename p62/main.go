package main

import (
	"fmt"
	"math/big"
)

func uniquePaths(m int, n int) int {
	return int(new(big.Int).Binomial(int64(m+n-2), int64(n-1)).Int64())
}

func main() {
	fmt.Println(uniquePaths(3, 7))
	fmt.Println(uniquePaths(3, 2))
	fmt.Println(uniquePaths(7, 3))
	fmt.Println(uniquePaths(3, 3))
}
