/*
 * @lc app=leetcode.cn id=6 lang=golang
 *
 * [6] N 字形变换
 *
 * https://leetcode.cn/problems/zigzag-conversion/description/
 *
 * algorithms
 * Medium (52.04%)
 * Likes:    2014
 * Dislikes: 0
 * Total Accepted:    548.6K
 * Total Submissions: 1.1M
 * Testcase Example:  '"PAYPALISHIRING"\n3'
 *
 * 将一个给定字符串 s 根据给定的行数 numRows ，以从上往下、从左到右进行 Z 字形排列。
 *
 * 比如输入字符串为 "PAYPALISHIRING" 行数为 3 时，排列如下：
 *
 *
 * P   A   H   N
 * A P L S I I G
 * Y   I   R
 *
 * 之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："PAHNAPLSIIGYIR"。
 *
 * 请你实现这个将字符串进行指定行数变换的函数：
 *
 *
 * string convert(string s, int numRows);
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "PAYPALISHIRING", numRows = 3
 * 输出："PAHNAPLSIIGYIR"
 *
 * 示例 2：
 *
 *
 * 输入：s = "PAYPALISHIRING", numRows = 4
 * 输出："PINALSIGYAHRPI"
 * 解释：
 * P     I    N
 * A   L S  I G
 * Y A   H R
 * P     I
 *
 *
 * 示例 3：
 *
 *
 * 输入：s = "A", numRows = 1
 * 输出："A"
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * s 由英文字母（小写和大写）、',' 和 '.' 组成
 * 1
 *
 *
 */
package medium

// @lc code=start
func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	firstStep := 2 * (numRows - 1)
	res := make([]rune, 0, len(s))
	origin := []rune(s)
	for i := 0; i < numRows; i++ {
		step := firstStep - 2*i
		j := i
		if i >= len(origin) {
			break
		}
		res = append(res, origin[j])
		for j < len(s)-1 {
			if i == 0 || i == numRows-1 {
				if j+firstStep <= len(s)-1 {
					res = append(res, origin[j+firstStep])
				}
			} else {
				if j+step <= len(s)-1 {
					res = append(res, origin[j+step])
				}
				if j+firstStep <= len(s)-1 {
					res = append(res, origin[j+firstStep])
				}
			}
			j += firstStep
		}
	}
	return string(res)
}

// @lc code=end
