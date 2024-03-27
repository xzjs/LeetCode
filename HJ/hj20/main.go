package main

import (
	"fmt"
)

func main() {
	// scanner := bufio.NewScanner(os.Stdin)
	// for scanner.Scan() {
	// s := scanner.Text()
	ss := []string{"021Abc9000", "021Abc9Abc1", "021ABC9000", "021$bc9000"}
	for _, s := range ss {
		if valid(s) {
			fmt.Println("OK")
		} else {
			fmt.Println("NG")
		}
	}
}

func valid(s string) bool {
	if len(s) < 8 {
		return false
	}
	diffMap := make(map[int]int)
	for _, v := range s {
		if v >= 'a' && v <= 'z' {
			diffMap[1]++
		} else if v >= 'A' && v <= 'Z' {
			diffMap[2]++
		} else if v >= '0' && v <= '9' {
			diffMap[3]++
		} else {
			diffMap[4]++
		}
	}
	if len(diffMap) < 3 {
		return false
	}
	window := make(map[string]int)
	for i := 0; i < len(s)-2; i++ {
		if _, ok := window[s[i:i+3]]; ok {
			return false
		} else {
			window[s[i:i+3]]++
		}
	}
	return true
}
