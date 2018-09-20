package leetcode

import (
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
