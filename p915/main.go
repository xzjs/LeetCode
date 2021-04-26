package main

import (
	"fmt"
)

func partitionDisjoint(A []int) int {
	length := len(A)
	maxLeft := make([]int, length)
	minRight := make([]int, length)

	m := A[0]
	for i, v := range A {
		if m < v {
			m = v
		}
		maxLeft[i] = m
	}

	m = A[length-1]
	for i := length - 1; i >= 0; i-- {
		if m > A[i] {
			m = A[i]
		}
		minRight[i] = m
	}

	for i := 1; i < length; i++ {
		if maxLeft[i-1] < minRight[i] {
			return i
		}
	}
	return 0
}

func main() {
	A := []int{5, 0, 3, 8, 6}
	fmt.Println(partitionDisjoint(A))
	A = []int{1, 1, 1, 0, 6, 12}
	fmt.Println(partitionDisjoint(A))
}
