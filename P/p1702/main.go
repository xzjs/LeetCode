package main

import "fmt"

func maximumBinaryString(binary string) string {
	if len(binary) < 2 {
		return binary
	}
	bs := []byte(binary)
	n := len(bs)
	for i := 0; i < n-1; i++ {
		if bs[i] == '0' {
			if bs[i+1] == '0' {
				bs[i] = '1'
			} else if bs[i+1] == '1' && i+2 < n && bs[i+2] == '0' {
				bs[i] = '1'
				bs[i+1] = '0'
			}
		}
	}
	return string(bs)
}

func main() {
	binary := "000110"
	fmt.Println(maximumBinaryString(binary))
}
