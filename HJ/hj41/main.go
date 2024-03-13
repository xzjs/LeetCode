package main

import (
	"fmt"
)

func main1() {
	var n int
	if _, err := fmt.Scanf("%d", &n); err != nil {
		return
	}
	m, x := make([]int, n), make([]int, n)
	xTotal := 0
	for i := 0; i < n; i++ {
		if _, err := fmt.Scanf("%d", &m[i]); err != nil {
			return
		}
	}
	for i := 0; i < n; i++ {
		if _, err := fmt.Scanf("%d", &x[i]); err != nil {
			return
		}
		xTotal += x[i]
	}
	arr := make([]int, xTotal)
	for i, j := 0, 0; i < n; i++ {
		for k := 0; k < x[i]; k++ {
			arr[j] = m[i]
			j++
		}
	}
	hashMap := make(map[int]int)
	var df func(i, total int)
	df = func(i, total int) {
		if i == xTotal {
			hashMap[total]++
			return
		}
		//不选i
		df(i+1, total)
		//选
		df(i+1, total+arr[i])
	}
	df(0, 0)
	fmt.Println(len(hashMap))
}

func main() {
	a := []int{1}
	fmt.Println(a[1:])
}
