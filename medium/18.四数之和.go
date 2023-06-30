/*
 * @lc app=leetcode.cn id=18 lang=golang
 *
 * [18] 四数之和
 *
 * https://leetcode.cn/problems/4sum/description/
 *
 * algorithms
 * Medium (37.14%)
 * Likes:    1578
 * Dislikes: 0
 * Total Accepted:    445.6K
 * Total Submissions: 1.2M
 * Testcase Example:  '[1,0,-1,0,-2,2]\n0'
 *
 * 给你一个由 n 个整数组成的数组 nums ，和一个目标值 target 。请你找出并返回满足下述全部条件且不重复的四元组 [nums[a],
 * nums[b], nums[c], nums[d]] （若两个四元组元素一一对应，则认为两个四元组重复）：
 *
 *
 * 0 <= a, b, c, d < n
 * a、b、c 和 d 互不相同
 * nums[a] + nums[b] + nums[c] + nums[d] == target
 *
 *
 * 你可以按 任意顺序 返回答案 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [1,0,-1,0,-2,2], target = 0
 * 输出：[[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [2,2,2,2,2], target = 8
 * 输出：[[2,2,2,2]]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 200
 * -10^9 <= nums[i] <= 10^9
 * -10^9 <= target <= 10^9
 *
 *
 */
package medium

import (
	"fmt"
	"sort"
)

// @lc code=start
func fourSum(nums []int, target int) [][]int {
	if len(nums) < 4 {
		return nil
	}
	sort.Ints(nums)
	twoSumMap := make(map[int][][]int)
	twoSumResultMap := make(map[string]bool)
	result := make([][]int, 0)
	preMinSum := nums[0] + nums[1]

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {

			sum := nums[i] + nums[j]
			if sum+preMinSum > target {
				break
			}

			if twoSumMap[sum] == nil {
				twoSumMap[sum] = make([][]int, 0)
			}
			if !twoSumResultMap[fmt.Sprintf("%d-%d", nums[i], nums[j])] {
				twoSumResultMap[fmt.Sprintf("%d-%d", nums[i], nums[j])] = true
				twoSumMap[sum] = append(twoSumMap[sum], []int{i, j})
			}

			if pre, ok := twoSumMap[target-sum]; ok {
				for _, p := range pre {
					if p[0] == i || p[0] == j || p[1] == i || p[1] == j {
						continue
					}
					result = append(result, []int{nums[p[0]], nums[p[1]], nums[i], nums[j]})
				}
			}
		}
	}

	resultMap := make(map[string][]int, 0)

	result2 := make([][]int, 0)

	for i := range result {
		sort.Ints(result[i])
		if _, ok := resultMap[fmt.Sprintf("%d-%d-%d-%d", result[i][0], result[i][1], result[i][2], result[i][3])]; !ok {
			resultMap[fmt.Sprintf("%d-%d-%d-%d", result[i][0], result[i][1], result[i][2], result[i][3])] = result[i]
			result2 = append(result2, result[i])
		}
	}
	return result2
}

// @lc code=end
