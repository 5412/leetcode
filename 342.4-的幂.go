package leetcode

/*
 * @lc app=leetcode.cn id=342 lang=golang
 *
 * [342] 4的幂
 *
 * https://leetcode-cn.com/problems/power-of-four/description/
 *
 * algorithms
 * Easy (50.41%)
 * Likes:    219
 * Dislikes: 0
 * Total Accepted:    63.8K
 * Total Submissions: 123.4K
 * Testcase Example:  '16'
 *
 * 给定一个整数，写一个函数来判断它是否是 4 的幂次方。如果是，返回 true ；否则，返回 false 。
 *
 * 整数 n 是 4 的幂次方需满足：存在整数 x 使得 n == 4^x
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：n = 16
 * 输出：true
 *
 *
 * 示例 2：
 *
 *
 * 输入：n = 5
 * 输出：false
 *
 *
 * 示例 3：
 *
 *
 * 输入：n = 1
 * 输出：true
 *
 *
 *
 *
 * 提示：
 *
 *
 * -2^31
 *
 *
 *
 *
 * 进阶：
 *
 *
 * 你能不使用循环或者递归来完成本题吗？
 *
 *
 */

// @lc code=start
func isPowerOfFour(n int) bool {
	return n > 0 && n&(n-1) == 0 && n&0xaaaaaaaa == 0
}

// @lc code=end
