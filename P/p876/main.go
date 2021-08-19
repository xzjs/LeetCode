package main

import "fmt"

func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil {
		if fast.Next == nil {
			break
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func initListNode(nums []int) *ListNode {
	head := &ListNode{
		Val:  nums[0],
		Next: nil,
	}
	p := head
	for i := 1; i < len(nums); i++ {
		p.Next = &ListNode{Val: nums[i], Next: nil}
		p = p.Next
	}
	return head
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	listNode := initListNode(nums)
	fmt.Println(middleNode(listNode))
	nums = []int{1, 2, 3, 4, 5, 6}
	listNode = initListNode(nums)
	fmt.Println(middleNode(listNode))
}
