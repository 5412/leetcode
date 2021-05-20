package leetcode

/*
 * @lc app=leetcode.cn id=450 lang=golang
 *
 * [450] 删除二叉搜索树中的节点
 *
 * https://leetcode-cn.com/problems/delete-node-in-a-bst/description/
 *
 * algorithms
 * Medium (47.29%)
 * Likes:    448
 * Dislikes: 0
 * Total Accepted:    43.1K
 * Total Submissions: 91.2K
 * Testcase Example:  '[5,3,6,2,4,null,7]\n3'
 *
 * 给定一个二叉搜索树的根节点 root 和一个值 key，删除二叉搜索树中的 key
 * 对应的节点，并保证二叉搜索树的性质不变。返回二叉搜索树（有可能被更新）的根节点的引用。
 *
 * 一般来说，删除节点可分为两个步骤：
 *
 *
 * 首先找到需要删除的节点；
 * 如果找到了，删除它。
 *
 *
 * 说明： 要求算法时间复杂度为 O(h)，h 为树的高度。
 *
 * 示例:
 *
 *
 * root = [5,3,6,2,4,null,7]
 * key = 3
 *
 * ⁠   5
 * ⁠  / \
 * ⁠ 3   6
 * ⁠/ \   \
 * 2   4   7
 *
 * 给定需要删除的节点值是 3，所以我们首先找到 3 这个节点，然后删除它。
 *
 * 一个正确的答案是 [5,4,6,2,null,null,7], 如下图所示。
 *
 * ⁠   5
 * ⁠  / \
 * ⁠ 4   6
 * ⁠/     \
 * 2       7
 *
 * 另一个正确答案是 [5,2,6,null,4,null,7]。
 *
 * ⁠   5
 * ⁠  / \
 * ⁠ 2   6
 * ⁠  \   \
 * ⁠   4   7
 *
 *
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

/*
二叉搜索树中序遍历是一个递增数组
left node right
 删除节点替换规则
 1. 叶子节点直接删除
 2. 存在右子树, 寻找右子树最小值替换, 删除右子树 最小值节点
 3. 无右子树存在 左子树, 寻找左子树最大值替换该节点值, 删除左子树最大值节点
*/
func findMix(root *TreeNode) int {
	for root.Left != nil {
		root = root.Left
	}
	return root.Val
}

func findMax(root *TreeNode) int {
	for root.Right != nil {
		root = root.Right
	}
	return root.Val
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	} else if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	} else {
		if root.Left == nil && root.Right == nil {
			root = nil
		} else if root.Right != nil {
			min := findMix(root.Right)
			root.Val = min
			root.Right = deleteNode(root.Right, min)
		} else {
			max := findMax(root.Left)
			root.Val = max
			root.Left = deleteNode(root.Left, max)
		}
	}
	return root
}

// @lc code=end
