/*
 * @lc app=leetcode.cn id=4 lang=golang
 *
 * [4] 寻找两个正序数组的中位数
 *
 * https://leetcode.cn/problems/median-of-two-sorted-arrays/description/
 *
 * algorithms
 * Hard (41.66%)
 * Likes:    6434
 * Dislikes: 0
 * Total Accepted:    925.7K
 * Total Submissions: 2.2M
 * Testcase Example:  '[1,3]\n[2]'
 *
 * 给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。
 *
 * 算法的时间复杂度应该为 O(log (m+n)) 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums1 = [1,3], nums2 = [2]
 * 输出：2.00000
 * 解释：合并数组 = [1,2,3] ，中位数 2
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums1 = [1,2], nums2 = [3,4]
 * 输出：2.50000
 * 解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5
 *
 *
 *
 *
 *
 *
 * 提示：
 *
 *
 * nums1.length == m
 * nums2.length == n
 * 0 <= m <= 1000
 * 0 <= n <= 1000
 * 1 <= m + n <= 2000
 * -10^6 <= nums1[i], nums2[i] <= 10^6
 *
 *
 */
package hard

// @lc code=start
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l1, l2 := len(nums1), len(nums2)
	total := +l1 + l2
	middleIndex1, middleIndex2 := (total-1)/2, total/2
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	quickFind := func(k int) int {
		index1, index2 := 0, 0
		for {
			if index1 == len(nums1) {
				return nums2[index2+k-1]
			}
			if index2 == len(nums2) {
				return nums1[index1+k-1]
			}
			if k == 1 {
				return min(nums1[index1], nums2[index2])
			}
			half := k / 2
			newIndex1 := min(index1+half, len(nums1)) - 1
			newIndex2 := min(index2+half, len(nums2)) - 1
			pivot1, pivot2 := nums1[newIndex1], nums2[newIndex2]
			if pivot1 <= pivot2 {
				k -= (newIndex1 - index1 + 1)
				index1 = newIndex1 + 1
			} else {
				k -= (newIndex2 - index2 + 1)
				index2 = newIndex2 + 1
			}
		}
	}

	return float64(quickFind(middleIndex1+1)+quickFind(middleIndex2+1)) / 2
}

// @lc code=end
