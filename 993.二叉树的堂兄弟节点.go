package leetcode

/*
 * @lc app=leetcode.cn id=993 lang=golang
 *
 * [993] 二叉树的堂兄弟节点
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isCousins(root *TreeNode, x int, y int) bool {
	fn := func(i, level, parent int, levels, parents *[]int) {
		if i == x || i == y {
			*levels = append(*levels, level)
			*parents = append(*parents, parent)
		}
	}

	levels := make([]int, 0, 2)
	parents := make([]int, 0, 2)
	deepLoop(root, 0, 0, fn, &levels, &parents)

	return levels[0] == levels[1] && parents[0] != parents[1]
}

type helperFunc func(i, level, parent int, levels, parents *[]int)

func deepLoop(node *TreeNode, level, parent int, fn helperFunc, levels, parents *[]int) {
	if node == nil {
		return
	}
	fn(node.Val, level, parent, levels, parents)

	if node.Left != nil {
		deepLoop(node.Left, level+1, node.Val, fn, levels, parents)
	}
	if node.Right != nil {
		deepLoop(node.Right, level+1, node.Val, fn, levels, parents)
	}
}

// @lc code=end
