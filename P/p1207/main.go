package main

func uniqueOccurrences(arr []int) bool {
	map1 := make(map[int]int)
	for _, v := range arr {
		map1[v]++
	}
	map2 := make(map[int]int)
	for _, v := range map1 {
		if map2[v] == 0 {
			map2[v]++
		} else {
			return false
		}
	}
	return true
}

func main() {
	arr1 := []int{1, 2, 2, 1, 1, 3}
	print(uniqueOccurrences(arr1))
}
