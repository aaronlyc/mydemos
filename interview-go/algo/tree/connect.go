package tree

import "fmt"

type Node struct {
	Val               int
	Left, Right, Next *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	traverseConnect(root.Left, root.Right)
	return root
}

func traverseConnect(node1, node2 *Node) {
	if node1 == nil || node2 == nil {
		return
	}

	node1.Next = node2
	fmt.Printf("node %d next is %d\n", node1.Val, node2.Val)
	traverseConnect(node1.Left, node1.Right)
	traverseConnect(node2.Left, node2.Right)
	traverseConnect(node1.Right, node2.Left)
}
