package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	dict := make(map[byte]int)
	distance, right := 0, -1
	for i := 0; i < len(s); i++ {
		if i != 0 {
			dict[s[i-1]] = 0
		}
		for (right+1 < len(s)) && (dict[s[right+1]] == 0) {
			right++
			dict[s[right]]++
		}
		d := right - i + 1
		if d > distance {
			distance = d
		}
	}
	return distance
}

func main() {
	s := "au"
	fmt.Println(lengthOfLongestSubstring(s))
	s = "bbbbb"
	fmt.Println(lengthOfLongestSubstring(s))
	s = "pwwkew"
	fmt.Println(lengthOfLongestSubstring(s))
	s = ""
	fmt.Println(lengthOfLongestSubstring(s))
	s = " "
	fmt.Println(lengthOfLongestSubstring(s))
}
