package main

import "fmt"

func reverseString(s []byte) {
	left, right := 0, len(s)-1
	for left <= right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}

func main() {
	chars := []byte{'h', 'e', 'l', 'l', 'o'}
	reverseString(chars)
	fmt.Println(chars)
}
