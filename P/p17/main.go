package main

import "fmt"

var data = map[string][]string{
	"2": {"a", "b", "c"},
	"3": {"d", "e", "f"},
	"4": {"g", "h", "i"},
	"5": {"j", "k", "l"},
	"6": {"m", "n", "o"},
	"7": {"p", "q", "r", "s"},
	"8": {"t", "u", "v"},
	"9": {"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	result := []string{}
	if len(digits) == 0 {
		return result
	}
	if len(digits) == 1 {
		return data[digits]
	}
	head, tail := digits[:1], digits[1:]
	for _, v1 := range data[head] {
		temp := letterCombinations(tail)
		for _, v2 := range temp {
			result = append(result, v1+v2)
		}
	}
	return result
}

func main() {
	fmt.Println(letterCombinations("23"))
}
