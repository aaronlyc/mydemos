package tree

import (
	"testing"
)

func testNodes() *Node {
	// 构建二叉树:
	//       1
	//      / \
	//     2   3
	//    / \  / \
	//   4   5 6  7
	root := &Node{Val: 1}
	root.Left = &Node{Val: 2}
	root.Right = &Node{Val: 3}
	root.Left.Left = &Node{Val: 4}
	root.Left.Right = &Node{Val: 5}
	root.Right.Left = &Node{Val: 6}
	root.Right.Right = &Node{Val: 7}
	return root
}

func Test_connect(t *testing.T) {
	root := testNodes()

	connect(root)
}
