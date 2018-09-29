package easy

import (
	"math"
	"strconv"
)

func FizzBuzz(n int) []string {

	result := make([]string, 0, n)
	for i:=1; i<=n;i++ {
		if i%15 == 0 {
			result = append(result, "FizzBuzz")
		} else if i%5 == 0 {
			result = append(result, "Buzz")
		} else if i%3 == 0 {
			result = append(result, "Fizz")
		} else {
			result = append(result, strconv.Itoa(i))
		}
	}
	return result

	//strs := make([]string,n)
	//for i:=1; i<n+1; i++ {
	//	temp := ""
	//	switch  {
	//	case i % 15 == 0:
	//		temp = "FizzBuzz"
	//	case i % 3 == 0 :
	//		temp = "Fizz"
	//	case i % 5 == 0:
	//		temp = "Buzz"
	//	default:
	//		temp = strconv.Itoa(i)
	//	}
	//	strs[i-1] = temp
	//}
	//return strs
}

//func CountPrimes(n int) int {
//	num := 0
//	for i:=1; i<n; i++ {
//		if isPrime(i) {
//			num++
//		}
//	}
//	return num
//}
//
//func isPrime(n int) bool {
//	if n < 2 {
//		return false
//	}
//	for i:=2; i*i<n; i++ {
//		if n % i == 0 {
//			return false
//		}
//	}
//	return true
//}

// leetcode God's algorithm

func CountPrimes(n int) int {
	notPrime := make([]bool, n)
	count := 0

	for i := 2; i < n; i++ {
		if !notPrime[i] {
			count++
			if temp := i * i; temp < n {
				for j := temp; j < n; j += i {
					notPrime[j] = true
				}
			}
		}
	}
	return count
}

func IsPowerOfThree(n int) bool {
	// God's algorithm
	// 换底公式： (log a b) = (log c a) / (log c b)

	// 计算int范围内最大的3的幂数
	if n <= 0 {
		return false
	}

	if n == 1 || n == 3 {
		return true
	}

	// note: 最大的Int值竟然计算的值不准，目前还不晓得原因, 经测试位数大于54后出现精度不准问题
	// max3 := int(math.Log10(0x7fffffff) / math.Log10(3))
	max3 := int(math.Log10(1 << 50) / math.Log10(3))

	maxNum := math.Pow(3, float64(max3))

	if int(maxNum) % n == 0 {
		return true
	}
	return false

	//if n == 3 || n == 1 {
	//	return true
	//}
	//
	//pow := 3
	//
	//for {
	//	pow *= 3
	//	if pow == n {
	//		return true
	//	}
	//	if pow > n {
	//		break
	//	}
	//}
	//
	//return false
}

func RomanToInt(s string) int {
	result := 0
	m := map[rune]int{'I':1, 'V':5, 'X':10, 'L':50, 'C':100, 'D':500, 'M':1000}
	runes := []rune(s)
	for i:=0; i<len(runes);i++ {
		switch {
		case i < len(runes) - 1 && m[runes[i]] < m[runes[i+1]]:
			result += m[runes[i+1]] - m[runes[i]]
			i++
		default:
			result += m[runes[i]]
		}
	}
	return result
}
