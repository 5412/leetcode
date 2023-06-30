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

import (
	"math"
	"sort"
)

// @lc code=start
func threeSumClosest(nums []int, target int) int {
	distance := func(current int) int {
		if current-target < 0 {
			return target - current
		}
		return current - target
	}
	if len(nums) < 3 {
		return 0
	}
	now := distance(nums[0] + nums[1] + nums[2])
	sum := nums[0] + nums[1] + nums[2]
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for z := j + 1; z < len(nums); z++ {
				current := nums[i] + nums[j] + nums[z]
				dis := distance(current)
				if dis <= now {
					now = dis
					sum = current
				}
				if now == 0 {
					return sum
				}
			}
		}
	}
	return sum
}

/** leetcode 官方答案
作者：LeetCode-Solution
链接：https://leetcode.cn/problems/3sum-closest/solution/zui-jie-jin-de-san-shu-zhi-he-by-leetcode-solution/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

func threeSumClosestLeetcode(nums []int, target int) int {
	sort.Ints(nums)
	var (
		n    = len(nums)
		best = math.MaxInt32
	)

	// 根据差值的绝对值来更新答案
	update := func(cur int) {
		if abs(cur-target) < abs(best-target) {
			best = cur
		}
	}

	// 枚举 a
	for i := 0; i < n; i++ {
		// 保证和上一次枚举的元素不相等
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		// 使用双指针枚举 b 和 c
		j, k := i+1, n-1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			// 如果和为 target 直接返回答案
			if sum == target {
				return target
			}
			update(sum)
			if sum > target {
				// 如果和大于 target，移动 c 对应的指针
				k0 := k - 1
				// 移动到下一个不相等的元素
				for j < k0 && nums[k0] == nums[k] {
					k0--
				}
				k = k0
			} else {
				// 如果和小于 target，移动 b 对应的指针
				j0 := j + 1
				// 移动到下一个不相等的元素
				for j0 < k && nums[j0] == nums[j] {
					j0++
				}
				j = j0
			}
		}
	}
	return best
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}

// @lc code=end
