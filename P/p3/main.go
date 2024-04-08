package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	n := len(s)
	maxLength := 1
	i, j := 0, 0
	for ; i < n; i++ {
		dict := make(map[byte]bool)
		for j = i; j < n; j++ {
			if _, ok := dict[s[j]]; ok {
				break
			}
			maxLength = max(maxLength, j-i+1)
			dict[s[j]] = true
		}
	}
	return maxLength
}

func main() {
	s := "abba"
	fmt.Println(lengthOfLongestSubstring(s))
}
