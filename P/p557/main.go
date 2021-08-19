package main

import "fmt"

func reverseWords(s string) string {
	left := 0
	result := []byte(s)
	result = append(result, ' ')
	for k, v := range result {
		if v == ' ' {
			right := k - 1
			for left <= right {
				result[left], result[right] = result[right], result[left]
				left++
				right--
			}
			left = k + 1
		}
	}
	return string(result[:len(s)])
}

func main() {
	s := "Let's take LeetCode contest"
	fmt.Println(reverseWords(s))
}
