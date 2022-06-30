package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	for p1, p2, tail := m-1, n-1, m+n-1; p1 >= 0 || p2 >= 0; tail-- {
		var cur int
		if p1 == -1 {
			cur = nums2[p2]
			p2--
		} else if p2 == -1 {
			cur = nums1[p1]
			p1--
		} else if nums1[p1] > nums2[p2] {
			cur = nums1[p1]
			p1--
		} else {
			cur = nums2[p2]
			p2--
		}
		nums1[tail] = cur
	}
}

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3
	merge(nums1, m, nums2, n)
	fmt.Println(nums1)
	nums1 = []int{1}
	m = 1
	nums2 = []int{}
	n = 0
	merge(nums1, m, nums2, n)
	fmt.Println(nums1)
	nums1 = []int{0}
	m = 0
	nums2 = []int{1}
	n = 1
	merge(nums1, m, nums2, n)
	fmt.Println(nums1)
	nums1 = []int{4, 5, 6, 0, 0, 0}
	m = 3
	nums2 = []int{1, 2, 3}
	n = 3
	merge(nums1, m, nums2, n)
	fmt.Println(nums1)
}
