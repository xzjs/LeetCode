package main

import "fmt"

func isSubsequence(s string, t string) bool {
	p1, p2 := 0, 0
	for p1 < len(s) && p2 < len(t) {
		if s[p1] == t[p2] {
			p1++
		}
		p2++
	}
	return p1 == len(s)
}

func main() {
	var s, t string
	s, t = "abc", "ahbgdc"
	fmt.Println(isSubsequence(s, t))
	s, t = "axc", "ahbgdc"
	fmt.Println(isSubsequence(s, t))
}
