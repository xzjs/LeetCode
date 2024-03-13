package main

import (
	"fmt"
)

func main() {
	hashMap := make(map[byte]int)
	var s, t string
	if _, err := fmt.Scanf("%s %s", &s, &t); err != nil {
		return
	}
	// s, t = "XDOYEZODEYXNZ", "XYZ"
	for _, v := range t {
		hashMap[byte(v)]++
	}
	window := make(map[byte]int)
	minLength := len(s)
	start, end, l, r, match := 0, 0, 0, 0, 0
	for ; r < len(s); r++ {
		if _, ok := hashMap[s[r]]; ok {
			window[s[r]]++
			if window[s[r]] == hashMap[s[r]] {
				match++
			}
			for len(hashMap) == match {
				length := r - l + 1
				if length < minLength {
					minLength = length
					start, end = l, r
				}
				if _, ok := hashMap[s[l]]; ok {
					window[s[l]]--
					if window[s[l]] < hashMap[s[l]] {
						match--
					}
				}
				l++
			}
		}
	}
	fmt.Println(s[start : end+1])
}
