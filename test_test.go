package leetcode

import (
	"fmt"
	"testing"
)

func TestCanCross(t *testing.T) {
	fmt.Println(canCross([]int{0, 1, 2, 3, 4, 8, 9, 11}))
}

func TestIsCousins(t *testing.T) {
	root := TreeNode{
		Val: 1,
	}
	root.Left = nil
	node1 := TreeNode{Val: 2}
	root.Right = &node1
	node2 := TreeNode{Val: 3}
	node1.Right = &node2
	node3 := TreeNode{Val: 4}
	node2.Right = &node3
	node4 := TreeNode{Val: 5}
	node3.Right = &node4
	fmt.Println(isCousins(&root, 1, 3))
}
