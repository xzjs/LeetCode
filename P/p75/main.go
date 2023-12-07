package main

import "fmt"

func canVisitAllRooms(rooms [][]int) bool {
	keys := make(map[int]bool)
	keys[0] = true
	queue := []int{0}
	for len(queue) > 0 {
		i := queue[0]
		for _, v := range rooms[i] {
			if _, ok := keys[v]; !ok {
				keys[v] = true
				queue = append(queue, v)
			}
		}
		queue = queue[1:]
	}
	return len(keys) == len(rooms)
}

func main() {
	rooms := [][]int{{1}, {2}, {3}, {}}
	fmt.Println(canVisitAllRooms(rooms))
}
