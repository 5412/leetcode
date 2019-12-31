package medium

import (
	"sort"
	"strconv"
)

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

/**
给定一个大小为 n 的数组，找到其中的多数元素。多数元素是指在数组中出现次数大于 ⌊ n/2 ⌋ 的元素。
你可以假设数组是非空的，并且给定的数组总是存在多数元素。
示例 1:

输入: [3,2,3]
输出: 3
示例 2:

输入: [2,2,1,1,1,2,2]
输出: 2
 */

func MajorityElement(nums []int) int {
	m := make(map[int]int)
	l := len(nums) / 2
	for _, v := range nums {
		m[v]++
		if m[v] > l {
			return v
		}
	}
	return 0
}

/**
解法三 摩尔投票法
1980 年由 Boyer 和 Moore 两个人提出来的算法，英文是 Boyer-Moore Majority Vote Algorithm。

算法思想很简单，但第一个想出来的人是真的强。

我们假设这样一个场景，在一个游戏中，分了若干个队伍，有一个队伍的人数超过了半数。所有人的战力都相同，不同队伍的两个人遇到就是同归于尽，同一个队伍的人遇到当然互不伤害。

这样经过充分时间的游戏后，最后的结果是确定的，一定是超过半数的那个队伍留在了最后。

而对于这道题，我们只需要利用上边的思想，把数组的每个数都看做队伍编号，然后模拟游戏过程即可。

group 记录当前队伍的人数，count 记录当前队伍剩余的人数。如果当前队伍剩余人数为 0，记录下次遇到的人的所在队伍号。

作者：windliang
链接：https://leetcode-cn.com/problems/majority-element/solution/xiang-xi-tong-su-de-si-lu-fen-xi-duo-jie-fa-by-4-2/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
 */
func majorityElement(nums []int) int {
	ret := nums[0]
	count := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == ret {
			count++
		} else {
			count--
			if count == 0 {
				ret = nums[i]
				count = 1
			}
		}
	}
	return ret
}

/**
给定一个用字符数组表示的 CPU 需要执行的任务列表。其中包含使用大写的 A - Z 字母表示的26 种不同种类的任务。任务可以以任意顺序执行，并且每个任务都可以在 1 个单位时间内执行完。CPU 在任何一个单位时间内都可以执行一个任务，或者在待命状态。
然而，两个相同种类的任务之间必须有长度为 n 的冷却时间，因此至少有连续 n 个单位时间内 CPU 在执行不同的任务，或者在待命状态。
你需要计算完成所有任务所需要的最短时间。
示例 1：
输入: tasks = ["A","A","A","B","B","B"], n = 2
输出: 8
执行顺序: A -> B -> (待命) -> A -> B -> (待命) -> A -> B.
注：
任务的总个数为 [1, 10000]。
n 的取值范围为 [0, 100]。
 */
func LeastInterval(tasks []byte, n int) int {
	taskMap := make(map[byte]int)
	mx := 0
	for _, t := range tasks {
		taskMap[t]++
		mx = max(mx, taskMap[t])
	}
	cn := 0
	for _, m := range taskMap {
		if m == mx {
			cn++
		}
	}

	return max(len(tasks), (mx-1) * (n + 1) + cn)
}

func leastInterval(tasks []byte, n int) int {
	taskCnt := len(tasks)
	if taskCnt < 1 {
		return 0
	}
	slots := make([]int, 26)

	for i := range tasks {
		slots[tasks[i] - 'A']++
	}

	sort.Ints(slots)
	maxCnt := slots[25] -1
	idles := maxCnt * n

	for i := 24; i >= 0 && slots[i] > 0; i-- {
		idles -= min2(slots[i], maxCnt)
	}

	if idles <= 0 {
		return taskCnt
	}
	return taskCnt + idles
}


func min2(a, b int) int {
	if a < b {
		return a
	}
	return b
}
