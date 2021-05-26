package leetcode

import (
	"strings"
	"unsafe"
)

/*
 * @lc app=leetcode.cn id=1190 lang=golang
 *
 * [1190] 反转每对括号间的子串
 *
 * https://leetcode-cn.com/problems/reverse-substrings-between-each-pair-of-parentheses/description/
 *
 * algorithms
 * Medium (59.26%)
 * Likes:    138
 * Dislikes: 0
 * Total Accepted:    25.7K
 * Total Submissions: 40.3K
 * Testcase Example:  '"(abcd)"'
 *
 * 给出一个字符串 s（仅含有小写英文字母和括号）。
 *
 * 请你按照从括号内到外的顺序，逐层反转每对匹配括号中的字符串，并返回最终的结果。
 *
 * 注意，您的结果中 不应 包含任何括号。
 *
 *
 *
 * 示例 1：
 *
 * 输入：s = "(abcd)"
 * 输出："dcba"
 *
 *
 * 示例 2：
 *
 * 输入：s = "(u(love)i)"
 * 输出："iloveu"
 *
 *
 * 示例 3：
 *
 * 输入：s = "(ed(et(oc))el)"
 * 输出："leetcode"
 *
 *
 * 示例 4：
 *
 * 输入：s = "a(bcdefghijkl(mno)p)q"
 * 输出："apmnolkjihgfedcbq"
 *
 *
 *
 *
 * 提示：
 *
 *
 * 0 <= s.length <= 2000
 * s 中只有小写英文字母和括号
 * 我们确保所有括号都是成对出现的
 *
 *
 */

// @lc code=start
// func reverseParentheses(s string) string {
// 	for {
// 		if !strings.Contains(s, "(") {
// 			break
// 		}
// 		s = reverseParenthesesHelper(s)
// 	}
// 	return s
// }

// func reverseParenthesesHelper(s string) string {
// 	left := -1
// 	right := strings.Index(s, ")")
// 	if right < 0 {
// 		return s
// 	}
// 	for i := right - 1; i >= 0; i-- {
// 		if s[i] == '(' {
// 			left = i
// 			break
// 		}
// 	}
// 	if left < 0 {
// 		return s
// 	}
// 	bts := []byte(s)
// 	i, j := left, right
// 	for i < j {
// 		bts[i], bts[j] = s[j], s[i]
// 		i++
// 		j--
// 	}
// 	return strings.Join([]string{string(bts[0:left]), string(bts[left+1 : right]), string(bts[right+1:])}, "")
// }

func reverseParentheses(s string) string {
	var stack []int
	bytes := *(*[]byte)(unsafe.Pointer(&s))
	for k, v := range s {
		if v == '(' {
			stack = append(stack, k)
		}
		if v == ')' {
			left := stack[len(stack)-1]
			reverseBytes(bytes, left, k)
			stack = stack[:len(stack)-1]
		}
	}
	var res strings.Builder
	for _, v := range bytes {
		if v != '(' && v != ')' {
			res.WriteByte(v)
		}
	}
	return res.String()
}

func reverseBytes(arr []byte, i int, k int) {
	for i < k {
		arr[i], arr[k] = arr[k], arr[i]
		i++
		k--
	}
}

// @lc code=end
