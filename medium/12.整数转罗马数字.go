/*
 * @lc app=leetcode.cn id=12 lang=golang
 *
 * [12] 整数转罗马数字
 *
 * https://leetcode.cn/problems/integer-to-roman/description/
 *
 * algorithms
 * Medium (66.16%)
 * Likes:    1096
 * Dislikes: 0
 * Total Accepted:    376.6K
 * Total Submissions: 569.4K
 * Testcase Example:  '3'
 *
 * 罗马数字包含以下七种字符： I， V， X， L，C，D 和 M。
 *
 *
 * 字符          数值
 * I             1
 * V             5
 * X             10
 * L             50
 * C             100
 * D             500
 * M             1000
 *
 * 例如， 罗马数字 2 写做 II ，即为两个并列的 1。12 写做 XII ，即为 X + II 。 27 写做  XXVII, 即为 XX + V +
 * II 。
 *
 * 通常情况下，罗马数字中小的数字在大的数字的右边。但也存在特例，例如 4 不写做 IIII，而是 IV。数字 1 在数字 5 的左边，所表示的数等于大数
 * 5 减小数 1 得到的数值 4 。同样地，数字 9 表示为 IX。这个特殊的规则只适用于以下六种情况：
 *
 *
 * I 可以放在 V (5) 和 X (10) 的左边，来表示 4 和 9。
 * X 可以放在 L (50) 和 C (100) 的左边，来表示 40 和 90。
 * C 可以放在 D (500) 和 M (1000) 的左边，来表示 400 和 900。
 *
 *
 * 给你一个整数，将其转为罗马数字。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: num = 3
 * 输出: "III"
 *
 * 示例 2:
 *
 *
 * 输入: num = 4
 * 输出: "IV"
 *
 * 示例 3:
 *
 *
 * 输入: num = 9
 * 输出: "IX"
 *
 * 示例 4:
 *
 *
 * 输入: num = 58
 * 输出: "LVIII"
 * 解释: L = 50, V = 5, III = 3.
 *
 *
 * 示例 5:
 *
 *
 * 输入: num = 1994
 * 输出: "MCMXCIV"
 * 解释: M = 1000, CM = 900, XC = 90, IV = 4.
 *
 *
 *
 * 提示：
 *
 *
 * 1
 *
 *
 */

package medium

// @lc code=start
func intToRoman2(num int) string {
	//进位与数字关系
	numsMap := map[int]map[int]string{
		1: {
			1: "I",
			5: "V",
		},
		10: {
			1: "X",
			5: "L",
		},
		100: {
			1: "C",
			5: "D",
		},
		1000: {
			1: "M",
		},
	}

	fn := func(radix map[int]string, num int, nextRadix string) string {
		res := ""

		if num <= 3 { //3
			for i := 0; i < num; i++ {
				res += radix[1]
			}
		} else if num == 4 { //4
			res += radix[1] + radix[5]
		} else if num == 5 { //5
			res += radix[5]
		} else if num <= 8 { //8
			res += radix[5]
			num -= 5
			for i := 0; i < num; i++ {
				res += radix[1]
			}
		} else { //9
			res += radix[1]
			res += nextRadix
		}
		return res
	}
	factor := 10
	res := ""
	for num > 0 {
		remainder := num % 10
		nextRadix := ""
		if factor < 10000 {
			nextRadix = numsMap[factor][1]
		}
		if remainder > 0 {
			ranNum := fn(numsMap[factor/10], remainder, nextRadix)
			res = ranNum + res
		}
		num = num / 10
		factor *= 10
	}

	return res
}

var valueSymbols = []struct {
	value  int
	symbol string
}{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func intToRoman(num int) string {
	roman := []byte{}
	for _, vs := range valueSymbols {
		for num >= vs.value {
			num -= vs.value
			roman = append(roman, vs.symbol...)
		}
		if num == 0 {
			break
		}
	}
	return string(roman)
}

// @lc code=end
