package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func totalCost(costs []int, k int, candidates int) int64 {
	sum := 0
	leftIndex, rightIndex := candidates, len(costs)-1
	leftArr, rightArr := []int{}, []int{}
	leftArr = append(leftArr, costs[:candidates]...)
	buildHeap(leftArr)
	for i := 0; i < candidates && rightIndex >= leftIndex; i++ {
		rightArr = append(rightArr, costs[rightIndex])
		rightIndex--
	}
	buildHeap(rightArr)
	for i := 0; i < k; i++ {
		if len(leftArr) > 0 && (len(rightArr) == 0 || leftArr[0] <= rightArr[0]) {
			size := len(leftArr)
			sum += leftArr[0]
			if leftIndex <= rightIndex {
				leftArr[0] = costs[leftIndex]
				leftIndex++
				heapfiy(leftArr, 0)
			} else {
				leftArr[0], leftArr[size-1] = leftArr[size-1], leftArr[0]
				leftArr = leftArr[:size-1]
				heapfiy(leftArr, 0)
			}
		} else {
			size := len(rightArr)
			sum += rightArr[0]
			if rightIndex >= leftIndex {
				rightArr[0] = costs[rightIndex]
				rightIndex--
				heapfiy(rightArr, 0)
			} else {
				rightArr[0], rightArr[size-1] = rightArr[size-1], rightArr[0]
				rightArr = rightArr[:size-1]
				heapfiy(rightArr, 0)
			}
		}
	}
	return int64(sum)
}

func totalCost1(costs []int, k, candidates int) int64 {
	ans := 0
	arr := []int{}
	if n := len(costs); candidates*2 < n {
		pre := hp{costs[:candidates]}
		heap.Init(&pre) // 原地建堆
		suf := hp{costs[n-candidates:]}
		heap.Init(&suf)
		for i, j := candidates, n-1-candidates; k > 0 && i <= j; k-- {
			if pre.IntSlice[0] <= suf.IntSlice[0] {
				ans += pre.IntSlice[0]
				arr = append(arr, pre.IntSlice[0])
				pre.IntSlice[0] = costs[i]
				heap.Fix(&pre, 0)
				i++
			} else {
				ans += suf.IntSlice[0]
				arr = append(arr, suf.IntSlice[0])
				suf.IntSlice[0] = costs[j]
				heap.Fix(&suf, 0)
				j--
			}
		}
		costs = append(pre.IntSlice, suf.IntSlice...)
	}
	sort.Ints(costs)
	for _, c := range costs[:k] { // 也可以用快速选择算法求前 k 小
		ans += c
		arr = append(arr, c)
	}
	fmt.Println(arr)
	return int64(ans)
}

type hp struct{ sort.IntSlice }

func (hp) Push(interface{})     {} // 没有用到，留空即可
func (hp) Pop() (_ interface{}) { return }

func buildHeap(a []int) {
	size := len(a)
	for i := size / 2; i >= 0; i-- {
		heapfiy(a, i)
	}
}

func heapfiy(a []int, index int) {
	size := len(a)
	l, r, min := index*2+1, index*2+2, index
	if l < size && a[l] < a[min] {
		min = l
	}
	if r < size && a[r] < a[min] {
		min = r
	}
	if min != index {
		a[min], a[index] = a[index], a[min]
		heapfiy(a, min)
	}
}

func main() {
	costs := []int{515, 705, 303, 791, 304, 382, 756, 957, 402, 399, 743, 919, 568, 141, 894, 488, 14, 452, 459, 930, 981, 662, 464, 663, 576, 302, 720, 855, 838, 51, 174, 97, 375, 813, 537, 750, 191, 991, 915, 972, 908, 370, 758, 864, 209, 478, 442, 685, 552, 717, 860, 996, 171, 168, 560, 595, 460, 285, 18, 96, 970, 746, 512, 420, 844, 183, 607, 267, 40, 491, 232, 278, 751, 277, 19, 419, 384, 85, 563, 556, 643, 896, 333, 468}
	k := 57
	candidates := 15
	fmt.Println(totalCost(costs, k, candidates))
	fmt.Println(totalCost1(costs, k, candidates))
}
