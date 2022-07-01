package main

import (
	"fmt"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(data []int) *ListNode {
	if len(data) == 0 {
		return nil
	}
	head := &ListNode{
		Val:  data[0],
		Next: nil,
	}
	p := head
	for i := 1; i < len(data); i++ {
		p.Next = &ListNode{Val: data[i], Next: nil}
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

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

func main() {
	l1 := NewListNode([]int{1, 2, 3, 4, 5})
	printListNode(reverseList(l1))
	l1 = NewListNode([]int{1, 2})
	printListNode(reverseList(l1))
	l1 = NewListNode([]int{})
	printListNode(reverseList(l1))
}
