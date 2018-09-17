package leetcode

import (
	"golang.org/x/tools/container/intsets"
	"math"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}


// leetcode: /explore/interview/card/top-interview-questions-easy/7/trees/47/

func MaxDepth(root *TreeNode) int  {
	if root == nil {
		return 0
	}

	hr := MaxDepth(root.Right)
	hl := MaxDepth(root.Left)

	if hr > hl {
		return  1 + hr
	}
	return  1 + hl
}


// leetcode: /submissions/detail/7196167/
// some God's algorithm
func IsValidBST(root *TreeNode) bool {
	max := math.MaxInt64
	min := math.MinInt64
	return isValidBSTc(root, max, min)
}

func isValidBSTc(node *TreeNode, max int, min int) bool {
	if node == nil {
		return true
	}
	if node.Val >= max || node.Val <= min {
		return false
	}
	return isValidBSTc(node.Left, node.Val, min) && isValidBSTc(node.Right, max, node.Val)
}
