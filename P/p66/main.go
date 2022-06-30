package main

import "fmt"

func plusOne(digits []int) []int {
	carry := true
	for i := len(digits) - 1; i >= 0; i-- {
		if carry {
			digits[i] += 1
			carry = false
			if digits[i] >= 10 {
				digits[i] -= 10
				carry = true
			}
		} else {
			break
		}
	}
	if carry {
		digits = append([]int{1}, digits...)
	}
	return digits
}

func main() {
	digits := []int{1, 2, 3}
	fmt.Println(plusOne(digits))
	digits = []int{4, 3, 2, 1}
	fmt.Println(plusOne(digits))
	digits = []int{0}
	fmt.Println(plusOne(digits))
	digits = []int{9}
	fmt.Println(plusOne(digits))
}
