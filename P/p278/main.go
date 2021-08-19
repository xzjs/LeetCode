package main

import "fmt"

func isBadVersion(v int) bool {
	return v >= 4
}

func firstBadVersion(n int) int {
	left := 1
	right := n
	for left < right {
		middle := (left + right) / 2
		if isBadVersion(middle) {
			right = middle
		} else {
			left = middle + 1
		}
	}
	return left
}

func main() {
	fmt.Println(firstBadVersion(5))
}
