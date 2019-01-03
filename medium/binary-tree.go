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

func BuildTree(preorder []int, inorder []int) *TreeNode {
	//fmt.Println(preorder, inorder)
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	node := &TreeNode{preorder[0], nil, nil}
	for i:=0; i<len(inorder); i++ {
		if preorder[0] == inorder[i] {
			if i-1 >= 0 {
				node.Left = BuildTree(preorder[1:i+1], inorder[0:i])
			} else {
				node.Left = nil
			}

			if i+1 < len(inorder) {
				node.Right = BuildTree(preorder[i+1:], inorder[i+1:])
			} else {
				node.Right = nil
			}

		}
	}
	return node
}

// leetcode 每个节点的右向指针
/**
 * Definition for binary tree with next pointer.
 * struct TreeLinkNode {
 *  int val;
 *  struct TreeLinkNode *left, *right, *next;
 * };
 *
 */
//void connect(struct TreeLinkNode *root) {
//
//if (root == NULL) {
//return;
//}
//connectHelper(root, true);
//}
//
//void connectHelper(struct TreeLinkNode *root, bool next) {
//
//if (root == NULL) {
//return;
//}
//struct TreeLinkNode *nextNode;
//
//if (root->left != NULL) {
//root->left->next = root->right;
//nextNode = root->left;
//if (root->next != NULL) {
//root->right->next = root->next->left;
//connectHelper(root->next, false);
//}
//} else {
//return;
//}
//
//if (nextNode != NULL && next) {
//connectHelper(nextNode, true);
//}
//}

//void connect(struct TreeLinkNode *root) {
//if(root==NULL) return;
//
//struct TreeLinkNode *b[10240]={0};
//int start=0;
//int end=1;
//int from=1;
//
//b[0]=root;
//while(start<end)
//{
//int i=0;
//for(i=start;i<end-1;i++)
//{
//b[i]->next=b[i+1];
//if(b[i]->left!=NULL)b[from++]=b[i]->left;
//if(b[i]->right!=NULL)b[from++]=b[i]->right;
//}
//if(i==end-1)
//{
//b[i]->next=NULL;
//if(b[i]->left!=NULL)b[from++]=b[i]->left;
//if(b[i]->right!=NULL)b[from++]=b[i]->right;
//}
//
//start=end;
//end=from;
//}
//}

// 二叉搜索树第K小的元素
// c代码实现,
//1、计算左子树元素个数left。
//
//2、 left+1 = K，则根节点即为第K个元素
//
//      left >=k, 则第K个元素在左子树中，
//
//     left +1 <k, 则转换为在右子树中，寻找第K-left-1元素。

// 方法二：因中序遍历为一个有序的数组，所以可以在中序遍历的过程中进行比较
//int kthSmallest(struct TreeNode* root, int k) {
//if (root == NULL)
//return 0;
//int leftSize = calcTreeSize(root->left);
//if (k == leftSize+1){
//return root->val;
//}else if (leftSize >= k){
//return kthSmallest(root->left,k);
//}else{
//return kthSmallest(root->right, k-leftSize-1);
//}
//
//}
//
//int calcTreeSize(struct TreeNode* root){
//if (root == NULL)
//return 0;
//return 1+calcTreeSize(root->left) + calcTreeSize(root->right);
//}
