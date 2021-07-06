package main

import (
	"fmt"
	"sort"
)

type void struct{}

var member void

func nthSuperUglyNumber(n int, primes []int) int {
	uglyNumbers := []int{1}
	mp := map[int]void{1: member}
	for i := 0; i < n; i++ {
		for _, val := range primes {
			temp := uglyNumbers[i] * val
			if _, has := mp[temp]; !has {
				uglyNumbers = append(uglyNumbers, temp)
				mp[temp] = member
			}
		}
		sort.Ints(uglyNumbers)
	}
	return uglyNumbers[n-1]
}

func main() {
	primes := []int{2, 7, 13, 19}
	n := 12
	fmt.Println(nthSuperUglyNumber(n, primes))
}
