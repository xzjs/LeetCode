package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	hashmap := make(map[int]int)
	for i := 0; i < len(inorder); i++ {
		hashmap[inorder[i]] = i
	}
	var build func(int, int) *TreeNode
	build = func(inorderLeft, inorderRight int) *TreeNode {
		if inorderLeft > inorderRight {
			return nil
		}
		val := postorder[len(postorder)-1]
		postorder = postorder[:len(postorder)-1]
		root := &TreeNode{Val: val}

		i := hashmap[val]
		root.Right = build(i+1, inorderRight)
		root.Left = build(inorderLeft, i-1)

		return root
	}
	return build(0, len(inorder)-1)
}

func main() {
	inorder := []int{9, 3, 15, 20, 7}
	postorder := []int{9, 15, 7, 20, 3}
	fmt.Println(buildTree(inorder, postorder))
}
