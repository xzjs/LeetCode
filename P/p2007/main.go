package main

import (
	"fmt"
	"sort"
)

func findOriginalArray(changed []int) []int {
	n := len(changed)
	if n%2 == 0 {
		res := []int{}
		sort.Ints(changed)
		dict := make(map[int]int)
		for _, v := range changed {
			if _, ok := dict[v]; ok {
				dict[v]--
				if dict[v] == 0 {
					delete(dict, v)
				}
			} else {
				res = append(res, v)
				dict[v*2]++
			}
		}
		if len(dict) == 0 {
			return res
		}
	}
	return []int{}
}

func main() {
	changed := []int{2, 1, 2, 4, 2, 4}
	fmt.Println(findOriginalArray(changed))
}
