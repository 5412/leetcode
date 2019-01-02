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

func ZigzagLevelOrder(root *TreeNode) [][]int {
	var result [][]int
	var layers []*TreeNode
	i := 1
	layers = append(layers, root)
	for len(layers) != 0 {
		temp := make([]*TreeNode, 0)
		nums := make([]int, 0)
		for _,node := range(layers) {
			nums = append(nums, node.Val)
			if node.Left != nil {
				temp = append(temp, node.Left)
			}
			if node.Right != nil {
				temp = append(temp, node.Right)
			}
		}
		if i % 2 == 0 {
			tmp := make([]int, 0)
			for j:= len(nums) -1;j>=0; j-- {
				tmp = append(tmp, nums[j])
			}
			nums = tmp
		}
		result = append(result, nums)
		layers = temp
		i++
	}
	return result
}


func ZigzagLevelOrderLeetcode(root *TreeNode) [][]int {
	if root==nil{
		return nil
	}
	stack := []*TreeNode{root}
	leftFlag := false
	ret := make([][]int,0,10)
	for len(stack)>0{
		temp := make([]int,0,len(stack))
		newStack := make([]*TreeNode,0,len(stack))
		for i:=len(stack)-1;i>=0;i--{
			temp  = append(temp,stack[i].Val)
			// 从左边开始
			if leftFlag{
				if stack[i].Right!=nil{
					newStack = append(newStack,stack[i].Right)
				}
				if stack[i].Left!=nil{
					newStack = append(newStack,stack[i].Left)
				}
			}else{
				if stack[i].Left!=nil{
					newStack = append(newStack,stack[i].Left)
				}
				if stack[i].Right!=nil{
					newStack = append(newStack,stack[i].Right)
				}
			}
		}
		leftFlag = !leftFlag
		ret = append(ret,temp)
		stack = newStack
	}
	return ret
}

