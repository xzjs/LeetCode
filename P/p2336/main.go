package main

import "fmt"

type SmallestInfiniteSet struct {
	current int
	minHeap []int
	dict    map[int]int
}

func Constructor() SmallestInfiniteSet {
	return SmallestInfiniteSet{
		current: 1,
		minHeap: []int{},
		dict:    make(map[int]int),
	}
}

func buildHeap(a []int) {
	size := len(a)
	for i := size / 2; i >= 0; i-- {
		minHeapfiy(a, i)
	}
}

func minHeapfiy(arr []int, index int) {
	size := len(arr)
	l, r, min := index*2+1, index*2+2, index
	if l < size && arr[l] < arr[min] {
		min = l
	}
	if r < size && arr[r] < arr[min] {
		min = r
	}
	if index != min {
		arr[index], arr[min] = arr[min], arr[index]
		minHeapfiy(arr, min)
	}
}

func (this *SmallestInfiniteSet) PopSmallest() int {
	result := 0
	if len(this.minHeap) > 0 {
		result = this.minHeap[0]
		delete(this.dict, result)
		size := len(this.minHeap)
		if size == 1 {
			this.minHeap = []int{}
		} else {
			this.minHeap[0], this.minHeap[size-1] = this.minHeap[size-1], this.minHeap[0]
			this.minHeap = this.minHeap[:size-1]
			minHeapfiy(this.minHeap, 0)
		}
	} else {
		result = this.current
		this.current++
	}
	return result
}

func (this *SmallestInfiniteSet) AddBack(num int) {
	if num >= this.current {
		return
	}
	if _, ok := this.dict[num]; ok {
		return
	}
	this.dict[num] = 0
	this.minHeap = append(this.minHeap, num)
	buildHeap(this.minHeap)
}

func main() {
	obj := Constructor()
	actions := []string{"SmallestInfiniteSet", "popSmallest", "addBack", "popSmallest", "popSmallest", "popSmallest", "addBack", "addBack", "popSmallest", "popSmallest"}

	var params = [][]int{{}, {}, {1}, {}, {}, {}, {2}, {3}, {}, {}}

	for i := 0; i < len(actions); i++ {
		switch actions[i] {
		case "addBack":
			obj.AddBack(params[i][0])
		case "popSmallest":
			fmt.Println(obj.PopSmallest())
		}
	}
}
