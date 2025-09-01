package main

import "fmt"

func main() {
	// Example: tree with depth = 3
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}

	fmt.Println(maxDepth(root)) // Output: 3
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	//fmt.Println("TreeNode:", root)
	if root == nil {
		return 0
	}
	fmt.Println(fmt.Printf("Node %v has memory addr. %p | Left Node's addr.: %p | Right Node's addr.: %p", root.Val, root, root.Left, root.Right))
	fmt.Println("calling left maxDepth for:", root.Val)
	left := maxDepth(root.Left) + 1
	fmt.Println("left:", left)
	fmt.Println("calling right maxDepth for:", root.Val)
	right := maxDepth(root.Right) + 1
	fmt.Println("right:", right)
	if left > right {
		return left
	}
	return right
}
