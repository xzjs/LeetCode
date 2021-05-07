package main

import "fmt"

func xorOperation(n int, start int) int {
	result := start
	for i := 1; i < n; i++ {
		result = result ^ (i*2 + start)
	}
	return result
}

func main() {
	fmt.Println(xorOperation(5, 0))
	fmt.Println(xorOperation(4, 3))
	fmt.Println(xorOperation(1, 7))
	fmt.Println(xorOperation(10, 5))
}
