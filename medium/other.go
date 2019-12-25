package medium

import "strconv"

/**
不使用运算符 + 和 - ​​​​​​​，计算两整数 ​​​​​​​a 、b ​​​​​​​之和。
示例 1:
输入: a = 1, b = 2
输出: 3
示例 2:
输入: a = -2, b = 3
输出: 1
 */

func getSum(a int, b int) int {
	result := a ^ b
	forward := a & b << 1
	for forward != 0 {
		return getSum(result, forward)
	}
	return result
}

/**
根据逆波兰表示法，求表达式的值。
有效的运算符包括 +, -, *, / 。每个运算对象可以是整数，也可以是另一个逆波兰表达式。
说明：
整数除法只保留整数部分。
给定逆波兰表达式总是有效的。换句话说，表达式总会得出有效数值且不存在除数为 0 的情况。
示例 1：
输入: ["2", "1", "+", "3", "*"]
输出: 9
解释: ((2 + 1) * 3) = 9
示例 2：

输入: ["4", "13", "5", "/", "+"]
输出: 6
解释: (4 + (13 / 5)) = 6
示例 3：
输入: ["10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"]
输出: 22
解释:
  ((10 * (6 / ((9 + 3) * -11))) + 17) + 5
= ((10 * (6 / (12 * -11))) + 17) + 5
= ((10 * (6 / -132)) + 17) + 5
= ((10 * 0) + 17) + 5
= (0 + 17) + 5
= 17 + 5
= 22
 */

func EvalRPN(tokens []string) int {
	type operation func(x, y int) int
	numStack := make([]int, 0)
	operator := map[string]operation{
		"+" : func(x, y int) int {
			return x + y
		},
		"-" : func(x, y int) int {
			return  x - y
		},
		"*" : func(x, y int) int {
			return x * y
		},
		"/" : func(x, y int) int {
			return x / y
		},
	}
	for _, token := range tokens {
		if o,  ok := operator[token]; ok {
			nums := numStack[len(numStack)-2:]
			numStack = numStack[:len(numStack)-2]
			numStack = append(numStack, o(nums[0], nums[1]))
		} else {
			num, _ := strconv.Atoi(token)
			numStack = append(numStack, num)
		}
	}
	return numStack[0]
}

func evalRPN(tokens []string) int {
	stack := []int{}
	for i := 0; i < len(tokens); i++ {
		if tokens[i] == "+" || tokens[i] == "-" || tokens[i] == "*" || tokens[i] == "/" {
			num1 := stack[len(stack)-1]
			num2 := stack[len(stack)-2]
			stack = stack[0 : len(stack)-2]
			num3 := 0
			switch tokens[i] {
			case "+":
				num3 = num1 + num2
			case "-":
				num3 = num2 - num1
			case "*":
				num3 = num1 * num2
			case "/":
				num3 = num2 / num1
			}
			stack = append(stack, num3)
			continue
		}
		num, _ := strconv.Atoi(tokens[i])
		stack = append(stack, num)
	}
	return stack[0]
}