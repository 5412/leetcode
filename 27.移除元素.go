/*
 * @lc app=leetcode.cn id=27 lang=golang
 *
 * [27] 移除元素
 */

// @lc code=start
func removeElement(nums []int, val int) int {
	var i, j = 0, len(nums)
	for i < j {
		if nums[i] == val {
			nums[i] = nums[j-1]
			j--
		} else {
			i++
		}
	}
	return j
}

// @lc code=end

