package tree

import (
	"fmt"
	"testing"
)

func newNode(val int) *TreeNode {
	return &TreeNode{Val: val}
}

func testTreeNodes() *TreeNode {
	// 构建二叉树:
	//       1
	//      / \
	//     2   3
	//    / \   \
	//   4   5   6
	root := newNode(1)
	root.Left = newNode(2)
	root.Right = newNode(3)
	root.Left.Left = newNode(4)
	root.Left.Right = newNode(5)
	root.Right.Right = newNode(6)
	return root
}

func Test_invertTree(t *testing.T) {
	root := testTreeNodes()

	invertTree(root)
}

func Test_toList(t *testing.T) {
	root := testTreeNodes()

	toList(root)
	printTreeNodes(root)
}

func printTreeNodes(root *TreeNode) {
	if root == nil {
		fmt.Println("nil")
		return
	}

	fmt.Printf("%d\n", root.Val)
	printTreeNodes(root.Left)
	printTreeNodes(root.Right)
}
