package leetcode

import (
	"math"
)

/*
 * @lc app=leetcode.cn id=664 lang=golang
 *
 * [664] 奇怪的打印机
 *
 * https://leetcode-cn.com/problems/strange-printer/description/
 *
 * algorithms
 * Hard (49.17%)
 * Likes:    157
 * Dislikes: 0
 * Total Accepted:    11K
 * Total Submissions: 18.1K
 * Testcase Example:  '"aaabbb"'
 *
 * 有台奇怪的打印机有以下两个特殊要求：
 *
 *
 * 打印机每次只能打印由 同一个字符 组成的序列。
 * 每次可以在任意起始和结束位置打印新字符，并且会覆盖掉原来已有的字符。
 *
 *
 * 给你一个字符串 s ，你的任务是计算这个打印机打印它需要的最少打印次数。
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "aaabbb"
 * 输出：2
 * 解释：首先打印 "aaa" 然后打印 "bbb"。
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = "aba"
 * 输出：2
 * 解释：首先打印 "aaa" 然后在第二个位置打印 "b" 覆盖掉原来的字符 'a'。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * s 由小写英文字母组成
 *
 *
 */

/*
 dp[i][i] = 1;
 dp[i][j] = dp[i][j-1]; s[i] = s[j]
 dp[i][j] = min(dp[i,k] + dp[k,j]); s[i] != s[j]; i<k<j;
*/

// @lc code=start
func strangePrinter(s string) int {
	l := len(s)
	dp := make([][]int, l)
	for i := 0; i < l; i++ {
		dp[i] = make([]int, l)
	}
	for i := l - 1; i >= 0; i-- {
		dp[i][i] = 1
		for j := i + 1; j < l; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i][j-1]
			} else {
				minx := math.MaxInt64
				for k := i; k < j; k++ {
					minx = min(minx, dp[i][k]+dp[k+1][j])
				}
				dp[i][j] = minx
			}
		}
	}
	return dp[0][l-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
