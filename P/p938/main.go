package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, low int, high int) int {
	stack := []*TreeNode{root}
	sum := 0
	for len(stack) > 0 {
		node := stack[0]
		stack = stack[1:]
		if node == nil {
			continue
		}
		if node.Val < low {
			stack = append(stack, node.Right)
		} else if node.Val > high {
			stack = append(stack, node.Left)
		} else {
			sum += node.Val
			stack = append(stack, node.Left, node.Right)
		}
	}
	return sum
}

func main() {
	root := &TreeNode{10, nil, nil}
	root.Left = &TreeNode{5, nil, nil}
	root.Right = &TreeNode{15, nil, nil}
	root.Left.Left = &TreeNode{3, nil, nil}
	root.Left.Right = &TreeNode{7, nil, nil}
	root.Right.Right = &TreeNode{18, nil, nil}
	fmt.Println(rangeSumBST(root, 7, 15))
	root = &TreeNode{10, nil, nil}
	root.Left = &TreeNode{5, nil, nil}
	root.Right = &TreeNode{15, nil, nil}
	root.Left.Left = &TreeNode{3, nil, nil}
	root.Left.Right = &TreeNode{7, nil, nil}
	root.Right.Left = &TreeNode{13, nil, nil}
	root.Right.Right = &TreeNode{18, nil, nil}
	root.Left.Left.Left = &TreeNode{1, nil, nil}
	root.Left.Right.Left = &TreeNode{6, nil, nil}
	fmt.Println(rangeSumBST(root, 6, 10))
}
