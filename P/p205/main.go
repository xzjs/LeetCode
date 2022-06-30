package main

import "fmt"

func isIsomorphic(s string, t string) bool {
	s2t := map[byte]byte{}
	t2s := map[byte]byte{}
	for i := range s {
		value, ok := s2t[s[i]]
		if ok && value != t[i] {
			return false
		} else {
			s2t[s[i]] = t[i]
		}
		value, ok = t2s[t[i]]
		if ok && value != s[i] {
			return false
		} else {
			t2s[t[i]] = s[i]
		}
	}
	return true
}

func main() {
	var s, t string
	s, t = "egg", "add"
	fmt.Println(isIsomorphic(s, t))
	s, t = "foo", "bar"
	fmt.Println(isIsomorphic(s, t))
	s, t = "paper", "title"
	fmt.Println(isIsomorphic(s, t))
	s, t = "badc", "baba"
	fmt.Println(isIsomorphic(s, t))
}
