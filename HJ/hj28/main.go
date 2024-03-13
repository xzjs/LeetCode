package main

import (
	"fmt"
)

func isPrime(num int) bool { //判断一个数是否为素数
	if num == 1 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func match(odd int, evens []int, visited map[int]int, suited map[int]int) bool {
	for _, even := range evens {
		if isPrime(odd+even) && visited[even] == 0 {
			visited[even] = 1
			if suited[even] == 0 || match(suited[even], evens, visited, suited) {
				suited[even] = odd
				return true
			}
		}
	}
	return false
}

func main() {
	// var n int
	var odds, evens []int

	// if _,err:=fmt.Scan(&n);err!=nil{
	// 	return
	// }
	// for i := 0; i < n; i++ {
	// 	var m int
	// 	if _, err := fmt.Scan(&m); err != nil {
	// 		return
	// 	}
	// 	if m%2 == 0 {
	// 		evens = append(evens, m)
	// 	} else {
	// 		odds = append(odds, m)
	// 	}
	// }

	odds = []int{5, 13}
	evens = []int{2, 6}

	suited := make(map[int]int) //连线
	res := 0
	for i := 0; i < len(odds); i++ {
		visited := make(map[int]int)
		if match(odds[i], evens, visited, suited) {
			res++
		}
	}
	fmt.Println(res)
}
