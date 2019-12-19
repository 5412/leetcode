package medium

/**
编写一个算法来判断一个数是不是“快乐数”。
一个“快乐数”定义为：对于一个正整数，每一次将该数替换为它每个位置上的数字的平方和，然后重复这个过程直到这个数变为 1，也可能是无限循环但始终变不到 1。如果可以变为 1，那么这个数就是快乐数。
示例:

输入: 19
输出: true
解释:
12 + 92 = 82
82 + 22 = 68
62 + 82 = 100
12 + 02 + 02 = 1
 */

func IsHappy(n int) bool {
	numsMap := make(map[int]int, 0)
	res := 0
	for res != 1  {
		if _, ok := numsMap[n]; ok {
			return false
		}
		nums := make([]int, 0)
		numsMap[n] = n
		res = 0
		for n > 0 {
			num := n%10
			nums = append(nums, num)
			res +=  num * num
			n = (n - num) / 10
		}
		n = res
	}
	return true
}

// leetcode algorithm
func isHappy(n int) bool {
	maps := map[int]int{}
	res := 0
	for res != 1 {
		for n > 0 {
			l := n % 10
			res += l * l
			n = n / 10
		}
		if res == 1 {
			return true
		}
		if _, ok := maps[res]; ok {
			return false
		}
		maps[res] = 1
		n = res
		res = 0
	}
	return true
}

/**
给定一个整数 n，返回 n! 结果尾数中零的数量。

示例 1:

输入: 3
输出: 0
解释: 3! = 6, 尾数中没有零。
示例 2:

输入: 5
输出: 1
解释: 5! = 120, 尾数中有 1 个零.
说明: 你算法的时间复杂度应为 O(log n) 。
关键就是找出能产生0的数来,可以知道,5的倍数与2的倍数相乘会产生0.
而2的倍数多于5的倍数,所以只需找出5的倍数有多少即可.
1991÷5^1=1991÷5=398.2,有398个5^1；
1991÷5^2=1991÷25=79.64,有79个5^2；
1991÷5^3=1991÷125=15.928,有15个5^3；
1991÷5^4=1991÷625=3.1856,有3个5^4.
它们的总和：398+79+15+3=495个.
也就是说,从1到1991的乘法算式里面,可以分解出来的5的质因数共有495个.每一个5与偶数相乘时都会产生一个0.所以共有495个0.
 */

func TrailingZeroes(n int) int {
	num := 0
	for n > 1  {
		num += n / 5
		n = n / 5
	}
	return num
}
