package main

import (
	"fmt"
	"strconv"
)

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	result := &ListNode{Val: 0, Next: head}
	slow, fast := result, result
	for i := 0; i <= n; i++ {
		fast = fast.Next
	}
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next
	return result.Next
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

func printListNode(head *ListNode) {
	for head != nil {
		fmt.Print(strconv.Itoa(head.Val) + "->")
		head = head.Next
	}
	fmt.Print("nil\n")
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	n := 2
	printListNode(removeNthFromEnd(initListNode(nums), n))
	nums = []int{1}
	n = 1
	printListNode(removeNthFromEnd(initListNode(nums), n))
	nums = []int{1, 2}
	n = 1
	printListNode(removeNthFromEnd(initListNode(nums), n))
}
