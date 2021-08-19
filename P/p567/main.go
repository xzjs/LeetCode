package main

import "fmt"

func deepCopy(des map[byte]int, src map[byte]int) {
	for k := range src {
		des[k] = src[k]
	}
}

func checkInclusion(s1 string, s2 string) bool {
	dict1, back := make(map[byte]int), make(map[byte]int)
	for i := 0; i < len(s1); i++ {
		dict1[s1[i]]++
	}
	for k := range dict1 {
		dict1[k]++
	}
	deepCopy(back, dict1)
	right := -1
	for i := 0; i < len(s2); i++ {
		if right+1 >= len(s2) {
			break
		}
		for right+1 < len(s2) {
			if dict1[s2[right+1]] == 0 { //字符不存在
				right++
				i = right
				dict1 = back
				deepCopy(dict1, back)
				break
			}
			if dict1[s2[right+1]] == 1 { //字符用完了
				dict1[s2[i]]++
				break
			}
			if dict1[s2[right+1]] > 1 {
				right++
				dict1[s2[right]]--
				if right-i+1 == len(s1) {
					return true
				}
			}
		}
	}
	return false
}

func main() {
	var s1, s2 string
	s1 = "hello"
	s2 = "ooolleoooleh"
	fmt.Println(checkInclusion(s1, s2))
	s1 = "mart"
	s2 = "karma"
	fmt.Println(checkInclusion(s1, s2))
	s1 = "ab"
	s2 = "eidbaooo"
	fmt.Println(checkInclusion(s1, s2))
	s1 = "ab"
	s2 = "eidboaoo"
	fmt.Println(checkInclusion(s1, s2))
}
