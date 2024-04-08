package main

import "fmt"

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	n := len(s)
	bs := make([]byte, n)
	index := 0
	for i := 0; i < numRows; i++ {
		mid := numRows - 1
		for j := i; j < n; {
			bs[index] = s[j]
			index++
			if i == 0 || i == numRows-1 {
				j += 2 * (numRows - 1)
			} else {
				j = 2*mid - j
				mid += numRows - 1
			}
		}
	}
	return string(bs)
}

func main() {
	s := "PAYPALISHIRING"
	numRows := 3
	fmt.Println(convert(s, numRows))
}
