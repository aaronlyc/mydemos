package tree

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 主函数
func invertTree(root *TreeNode) *TreeNode {
	// 遍历二叉树，交换每个节点的子节点
	traverse(root)
	return root
}

func traverse(root *TreeNode) {
	if root == nil {
		return
	}

	root.Left, root.Right = root.Right, root.Left
	fmt.Printf("\n starting---\nroot: %d\n", root.Val)
	if root.Left != nil {
		fmt.Printf("left: %d\n", root.Left.Val)
	}
	if root.Right != nil {
		fmt.Printf("right: %d\n", root.Right.Val)
	}
	traverse(root.Left)
	traverse(root.Right)
}

func toList(root *TreeNode) {
	if root == nil {
		return
	}

	toList(root.Left)
	toList(root.Right)

	left := root.Left
	right := root.Right

	root.Left = nil
	root.Right = left

	dump := root
	for dump.Right != nil {
		dump = dump.Right
	}
	dump.Right = right
}
