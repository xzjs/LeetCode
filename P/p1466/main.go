package main

import "fmt"

func minReorder(n int, connections [][]int) int {
	changeNum := 0
	cityStack := []int{0}
	for len(cityStack) > 0 {
		city := cityStack[0]
		cityStack = cityStack[1:]
		for i := 0; i < len(connections); i++ {
			if connections[i][1] == city {
				cityStack = append(cityStack, connections[i][0])
				connections[i] = []int{0, 0}
			} else if connections[i][0] == city {
				changeNum++
				cityStack = append(cityStack, connections[i][1])
				connections[i] = []int{0, 0}
			}
		}
	}
	return changeNum
}

func main() {
	n := 6
	connections := [][]int{{0, 1}, {1, 3}, {2, 3}, {4, 0}, {4, 5}}
	fmt.Println(minReorder(n, connections))
	n = 5
	connections = [][]int{{1, 0}, {1, 2}, {3, 2}, {3, 4}}
	fmt.Println(minReorder(n, connections))
	n = 3
	connections = [][]int{{1, 0}, {2, 0}}
	fmt.Println(minReorder(n, connections))
}
