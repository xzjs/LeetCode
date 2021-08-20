package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func magicTower(nums []int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum < 0 {
		return -1
	}
	helth := 1
	count := 0
	damage := &IntHeap{}
	heap.Init(damage)
	for _, v := range nums {
		helth += v
		if v < 0 {
			heap.Push(damage, v)
		}
		if helth < 1 {
			count++
			helth -= heap.Pop(damage).(int)
		}
	}
	return count
}

func main() {
	nums := []int{100, 100, 100, -250, -60, -140, -50, -50, 100, 150}
	fmt.Println(magicTower(nums))
	nums = []int{-200, -300, 400}
	fmt.Println(magicTower(nums))
	nums = []int{-1, -1, 10}
	fmt.Println(magicTower(nums))
}
