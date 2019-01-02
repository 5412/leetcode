// Copyright (c) 2019. weixiong piao, All Right Reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package medium

import (
	. "leetcode/structs"
)

// 循环方式的中序遍历
func InorderTraversal(root *TreeNode) []int {
	var nums []int
	var middles []*TreeNode
	node := root
	for node != nil || len(middles) != 0 {
		if node != nil {
			middles = append(middles,node)
			node = node.Left
		} else  {
			node = middles[len(middles) - 1]
			tmp := make([]*TreeNode, len(middles)-1)
			copy(tmp, middles[:len(middles)-1])
			nums = append(nums, node.Val)
			middles = tmp
			node = node.Right
		}
	}
	return nums
}

// 递归方式的中序遍历
func InorderTraversalRecursion(root *TreeNode) []int {
	var nums []int
	if root != nil  {
		if root.Left != nil {
			nums = append(nums, InorderTraversal(root.Left)...)
		}
		nums = append(nums, root.Val)
		if root.Right != nil {
			nums = append(nums, InorderTraversal(root.Right)...)
		}
	}
	return nums
}


