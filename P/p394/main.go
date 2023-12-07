package main

import "fmt"

func decodeString(s string) string {
	result := []rune{}
	numbers := []rune{}
	chars := []rune{}
	for _, v := range s {
		if v != ']' {
			result = append(result, v)
		} else {
			for result[len(result)-1] != '[' {
				chars = append(chars, result[len(result)-1])
				result = result[:len(result)-1]
			}
			result = result[:len(result)-1]
			for len(result) > 0 && result[len(result)-1] >= '0' && result[len(result)-1] <= '9' {
				numbers = append(numbers, result[len(result)-1])
				result = result[:len(result)-1]
			}
			for n := toInt(numbers); n > 0; n-- {
				for i := len(chars) - 1; i >= 0; i-- {
					result = append(result, chars[i])
				}
			}
			numbers = []rune{}
			chars = []rune{}
		}
	}
	return string(result)
}

func toInt(arr []rune) int {
	result := 0
	for i := len(arr) - 1; i >= 0; i-- {
		result = result*10 + int(arr[i]-'0')
	}
	return result
}

func main() {
	s := "100[leetcode]"
	fmt.Println(decodeString(s))
}
