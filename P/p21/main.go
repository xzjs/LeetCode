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

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	p := &ListNode{}
	head := p
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			p.Next = list1
			list1 = list1.Next
			p = p.Next
		} else {
			p.Next = list2
			list2 = list2.Next
			p = p.Next
		}
	}
	if list1 != nil {
		p.Next = list1
	} else {
		p.Next = list2
	}
	return head.Next
}

func main() {
	l1 := NewListNode([]int{1, 2, 4})
	l2 := NewListNode([]int{1, 3, 4})
	printListNode(mergeTwoLists(l1, l2))
	l1 = NewListNode([]int{})
	l2 = NewListNode([]int{})
	printListNode(mergeTwoLists(l1, l2))
	l1 = NewListNode([]int{})
	l2 = NewListNode([]int{0})
	printListNode(mergeTwoLists(l1, l2))
}
