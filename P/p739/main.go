package main

import "fmt"

func dailyTemperatures(temperatures []int) []int {
	length := len(temperatures)
	stack := []int{}
	ans := make([]int, length)
	for i := 0; i < length; i++ {
		if len(stack) > 0 && temperatures[stack[len(stack)-1]] < temperatures[i] {
			topIndex := stack[len(stack)-1]
			ans[topIndex] = i - topIndex
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return ans
}

func main() {
	temperatures := []int{73, 74, 75, 71, 69, 72, 76, 73}
	fmt.Println(dailyTemperatures(temperatures))
}
