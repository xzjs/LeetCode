package main

import "fmt"

func buildArray(target []int, n int) []string {
	result := []string{}
	i := 1
	for _, v := range target {
		for ; i <= n; i++ {
			result = append(result, "Push")
			if v != i {
				result = append(result, "Pop")
			} else {
				i++
				break
			}
		}
	}
	return result
}

func main() {
	target := []int{1, 3}
	n := 3
	fmt.Println(buildArray(target, n))
	target = []int{1, 2, 3}
	n = 3
	fmt.Println(buildArray(target, n))
	target = []int{1, 2}
	n = 4
	fmt.Println(buildArray(target, n))
	target = []int{2, 3, 4}
	n = 4
	fmt.Println(buildArray(target, n))
}
