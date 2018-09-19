package leetcode

import (
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

// leetcode: /explore/interview/card/top-interview-questions-easy/7/trees/50/

func LevelTreeOrder(root *TreeNode) [][]int {
	var result [][]int
	last := root
	nlast := root
	queue := make([]*TreeNode, 0)
	if root == nil {
		return result
	}
	queue = append(queue, root)
	level := make([]int, 0)

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node.Left != nil {
			queue = append(queue, node.Left)
			nlast = node.Left
		}

		if node.Right != nil {
			queue = append(queue, node.Right)
			nlast = node.Right
		}

		level = append(level, node.Val)

		if node == last {
			result = append(result, level)
			level = make([]int, 0)
			last = nlast
		}
	}
	return result
}
