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
	for res != 1 {
		if _, ok := numsMap[n]; ok {
			return false
		}
		nums := make([]int, 0)
		numsMap[n] = n
		res = 0
		for n > 0 {
			num := n % 10
			nums = append(nums, num)
			res += num * num
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
	for n > 1 {
		num += n / 5
		n = n / 5
	}
	return num
}

/**
给定一个Excel表格中的列名称，返回其相应的列序号。

例如，

    A -> 1
    B -> 2
    C -> 3
    ...
    Z -> 26
    AA -> 27
    AB -> 28
    ...
示例 1:

输入: "A"
输出: 1
示例 2:

输入: "AB"
输出: 28
示例 3:

输入: "ZY"
输出: 701
*/

func TitleToNumber(s string) int {
	// A = 'A' - 'A' + 1
	// B = 'B' - 'A' + 1
	l := len(s)
	res := 0
	for i := 0; i < l; i++ {
		nums := s[i] - 'A' + 1
		res = res*26 + int(nums)
	}
	return res
}

/**
实现 pow(x, n) ，即计算 x 的 n 次幂函数。
示例 1:
输入: 2.00000, 10
输出: 1024.00000
示例 2:
输入: 2.10000, 3
输出: 9.26100
示例 3:
输入: 2.00000, -2
输出: 0.25000
解释: 2-2 = 1/22 = 1/4 = 0.25
说明:
-100.0 < x < 100.0
n 是 32 位有符号整数，其数值范围是 [−231, 231 − 1] 。
*/
func MyPow(x float64, n int) float64 {
	if n < 0 {
		n = -n
		x = 1 / x
	}
	var pow, sq float64 = 1, x
	for n > 0 {
		if n&1 != 0 {
			pow *= sq
		}
		sq *= sq
		n >>= 1
	}
	return pow
}

func myPow(x float64, n int) float64 {
	if n < 0 {
		x = 1 / x
		n = -n
	}
	return pow(x, n)
}
func pow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	half := pow(x, n/2)
	res := float64(1)
	if n%2 == 0 {
		res = half * half
	} else {
		res = half * half * x
	}
	return res
}

/**
实现 int sqrt(int x) 函数。
计算并返回 x 的平方根，其中 x 是非负整数。
由于返回类型是整数，结果只保留整数的部分，小数部分将被舍去。
示例 1:
输入: 4
输出: 2
示例 2:
输入: 8
输出: 2
说明: 8 的平方根是 2.82842...,
     由于返回类型是整数，小数部分将被舍去。
*/

func MySqrt(x int) int {
	q := 1
	check := 1
	for check > 0 {
		q = (q + x/q) / 2
		check = q*q - x
	}
	return q
}

func mySqrt(x int) int {
	if x == 0 || x == 1 {
		return x
	}

	l, r := 1, x/2
	for l < r {
		m := l + (r-l+1)/2
		if m > x/m {
			r = m - 1
		} else {
			l = m
		}
	}
	return l
}

/**
给定两个整数，被除数 dividend 和除数 divisor。将两数相除，要求不使用乘法、除法和 mod 运算符。
返回被除数 dividend 除以除数 divisor 得到的商。
示例 1:
输入: dividend = 10, divisor = 3
输出: 3
示例 2:
输入: dividend = 7, divisor = -3
输出: -2

说明:
被除数和除数均为 32 位有符号整数。
除数不为 0。
假设我们的环境只能存储 32 位有符号整数，其数值范围是 [−2-31,  2-31 − 1]。本题中，如果除法结果溢出，则返回 231 − 1。
*/

func Divide(dividend int, divisor int) int {
	res, c := 0, 1
	var flag bool
	switch {
	case dividend > 0 && divisor < 0:
		{
			divisor = -divisor
			flag = false
		}
	case dividend < 0 && divisor > 0:
		{

			dividend = -dividend
			flag = false

		}
	case dividend < 0 && divisor < 0:
		{

			dividend = -dividend
			divisor = -divisor
			flag = true

		}
	default:
		flag = true
	}
	for dividend >= divisor<<1 {
		divisor <<= 1
		c <<= 1
	}
	for c > 0 && dividend > 0 {
		if dividend >= divisor {
			dividend -= divisor
			res += c
		}
		c >>= 1
		divisor >>= 1
	}
	if !flag {
		res = -res
	}
	if !(res >= int(-0x80000000) && res <= int(0x7FFFFFFF)) {
		return int(0x7FFFFFFF)
	}
	return res
}
