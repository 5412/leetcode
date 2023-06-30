/*
 * @lc app=leetcode.cn id=16 lang=golang
 *
 * [16] 最接近的三数之和
 *
 * https://leetcode.cn/problems/3sum-closest/description/
 *
 * algorithms
 * Medium (44.99%)
 * Likes:    1370
 * Dislikes: 0
 * Total Accepted:    459.1K
 * Total Submissions: 1M
 * Testcase Example:  '[-1,2,1,-4]\n1'
 *
 * 给你一个长度为 n 的整数数组 nums 和 一个目标值 target。请你从 nums 中选出三个整数，使它们的和与 target 最接近。
 *
 * 返回这三个数的和。
 *
 * 假定每组输入只存在恰好一个解。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [-1,2,1,-4], target = 1
 * 输出：2
 * 解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2) 。
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [0,0,0], target = 1
 * 输出：0
 *
 *
 *
 *
 * 提示：
 *
 *
 * 3 <= nums.length <= 1000
 * -1000 <= nums[i] <= 1000
 * -10^4 <= target <= 10^4
 *
 *
 */
package medium

import "testing"

func Test_threeSumClosest(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				nums:   []int{-1000, -1000, -1000},
				target: 10000,
			},
			want: -3000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeSumClosest(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("threeSumClosest() = %v, want %v", got, tt.want)
			}
		})
	}
}
