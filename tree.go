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

/* func IsSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left == nil && root.Right == nil {
		return true
	}

	return judge(root.Left, root.Right)
}

func judge(root1, root2 *TreeNode) bool {
	if root1 == nil && root2 ==nil {
		return true
	} else if root1 != nil && root2 != nil && root1.Val == root2.Val && judge(root1.Left, root2.Right) && judge(root1.Right, root2.Left) {
		return true
	} else {
		return false
	}
} */
