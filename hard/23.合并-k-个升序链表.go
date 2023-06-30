/*
 * @lc app=leetcode.cn id=23 lang=golang
 *
 * [23] 合并 K 个升序链表
 *
 * https://leetcode.cn/problems/merge-k-sorted-lists/description/
 *
 * algorithms
 * Hard (57.68%)
 * Likes:    2416
 * Dislikes: 0
 * Total Accepted:    635.4K
 * Total Submissions: 1.1M
 * Testcase Example:  '[[1,4,5],[1,3,4],[2,6]]'
 *
 * 给你一个链表数组，每个链表都已经按升序排列。
 *
 * 请你将所有链表合并到一个升序链表中，返回合并后的链表。
 *
 *
 *
 * 示例 1：
 *
 * 输入：lists = [[1,4,5],[1,3,4],[2,6]]
 * 输出：[1,1,2,3,4,4,5,6]
 * 解释：链表数组如下：
 * [
 * ⁠ 1->4->5,
 * ⁠ 1->3->4,
 * ⁠ 2->6
 * ]
 * 将它们合并到一个有序链表中得到。
 * 1->1->2->3->4->4->5->6
 *
 *
 * 示例 2：
 *
 * 输入：lists = []
 * 输出：[]
 *
 *
 * 示例 3：
 *
 * 输入：lists = [[]]
 * 输出：[]
 *
 *
 *
 *
 * 提示：
 *
 *
 * k == lists.length
 * 0 <= k <= 10^4
 * 0 <= lists[i].length <= 500
 * -10^4 <= lists[i][j] <= 10^4
 * lists[i] 按 升序 排列
 * lists[i].length 的总和不超过 10^4
 *
 *
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

package hard

import (
	"math"
	. "leetcode/structs"
)


func mergeKLists(lists []*ListNode) *ListNode {

	var head *ListNode
	var last *ListNode
	minFunc := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	for {
		emptySum := 0
		var current *ListNode
		var minList *ListNode
		var index int
		min := math.MaxInt
		min2 := math.MaxInt
		for i := 0; i < len(lists); i++ {
			if lists[i] != nil {
				if lists[i].Val <= min {
					min2 = min
					min = lists[i].Val
					minList = lists[i]
					index = i
				} else {
					min2 = minFunc(lists[i].Val, min2)
				}
			} else {
				emptySum++
			}
		}
		if emptySum == len(lists) {
			break
		}
		current = minList
		if head == nil {
			head = current
		} else {
			last.Next = current
			last = current
		}
		for current != nil && current.Val <= min2 {
			last = current
			current = current.Next
		}
		lists[index] = current
		last.Next = nil
	}
	return head
}

// @lc code=end
