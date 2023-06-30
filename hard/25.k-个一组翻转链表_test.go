/*
 * @lc app=leetcode.cn id=25 lang=golang
 *
 * [25] K 个一组翻转链表
 *
 * https://leetcode.cn/problems/reverse-nodes-in-k-group/description/
 *
 * algorithms
 * Hard (67.70%)
 * Likes:    2008
 * Dislikes: 0
 * Total Accepted:    457.4K
 * Total Submissions: 675.7K
 * Testcase Example:  '[1,2,3,4,5]\n2'
 *
 * 给你链表的头节点 head ，每 k 个节点一组进行翻转，请你返回修改后的链表。
 *
 * k 是一个正整数，它的值小于或等于链表的长度。如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。
 *
 * 你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：head = [1,2,3,4,5], k = 2
 * 输出：[2,1,4,3,5]
 *
 *
 * 示例 2：
 *
 *
 *
 *
 * 输入：head = [1,2,3,4,5], k = 3
 * 输出：[3,2,1,4,5]
 *
 *
 *
 * 提示：
 *
 *
 * 链表中的节点数目为 n
 * 1 <= k <= n <= 5000
 * 0 <= Node.val <= 1000
 *
 *
 *
 *
 * 进阶：你可以设计一个只用 O(1) 额外内存空间的算法解决此问题吗？
 *
 *
 *
 *
 */

package hard

import (
	. "leetcode/structs"
	"reflect"
	"testing"
)

func Test_reverseKGroup(t *testing.T) {
	type args struct {
		head []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "",
			args: args{
				head: []int{1, 2, 3, 4, 5},
				k:    2,
			},
			want: []int{2, 1, 4, 3, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverseKGroup(GenerateList(tt.args.head), tt.args.k)
			if !reflect.DeepEqual(GenerateSlice(got), tt.want) {
				PrintList(got)
			}
		})
	}
}
