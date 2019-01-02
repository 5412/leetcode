// Copyright (c) 2019. weixiong piao, All Right Reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package structs

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func GeneralBinaryTree(nums []int) *TreeNode {
	var (
		nodes []*TreeNode
		root  *TreeNode
	)
	for _,v := range(nums) {
		node := &TreeNode{v, nil,nil}
		if len(nodes) == 0 {
			root = node
			nodes = append(nodes, node)
		} else {
			if nodes[0].Left == nil {
				nodes[0].Left = node
			} else {
				nodes[0].Right = node
				nodes = nodes[1:]
			}
			nodes = append(nodes, node)
		}
	}
	return root
}