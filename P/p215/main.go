package main

import "fmt"

func findKthLargest(nums []int, k int) int {
	heapsize := len(nums)
	buildMaxHeap(nums, heapsize)
	for i := len(nums) - 1; i >= len(nums)-k+1; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		heapsize--
		maxHeapify(nums, 0, heapsize)
	}
	return nums[0]
}

func buildMaxHeap(a []int, heapsize int) {
	for i := heapsize / 2; i >= 0; i-- {
		maxHeapify(a, i, heapsize)
	}
}

func maxHeapify(a []int, i, heapsize int) {
	l, r, largest := i*2+1, i*2+2, i
	if l < heapsize && a[l] > a[largest] {
		largest = l
	}
	if r < heapsize && a[r] > a[largest] {
		largest = r
	}
	if largest != i {
		a[i], a[largest] = a[largest], a[i]
		maxHeapify(a, largest, heapsize)
	}
}

func main() {
	fmt.Println(findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4))
}
