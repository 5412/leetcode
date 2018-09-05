package main

func countRangeSum(nums []int, lower int, upper int) int {
	var snum int = 0
	l := len(nums)
	for i:=0;i<l;i++{
		sum := 0
		for j:=i;j<l;j++{
			sum += nums[j]
			if (sum >= lower && sum <= upper) {
				snum++
			}
		}
	}
	return snum
}

func removeDuplicates(nums []int)  int  {
	l := len(nums)
	n := 0

	for j:=0; j<l; j++ {
		if nums[n] != nums[j] {
			n++
			nums[n] = nums[j]
		}
	}

	return n + 1
}

func maxProfit(prices []int) int {
	l := len(prices)
	profit := 0
	for i:=1; i<l; i++ {
		if prices[i] > prices[i-1] {
			profit += prices[i] - prices[i-1]
		}
	}
	return profit
}

func rotate(nums []int, k int)  {

	if k < 0 {
		return
	}

	l := len(nums)
	k = k % l

	// space O(1)
	var temp int
	for i:=0; i<k; i++ {
		temp = nums[l-1]
		for j := l-1; j > 0; j-- {
			nums[j] = nums[j-1]
		}
		nums[0] = temp
	}

	// copy leetcode's answer

	//for i := 0; i < len(nums)/2; i++ {
	//	nums[i], nums[len(nums)-i-1] = nums[len(nums)-i-1], nums[i]
	//}
	//for i := 0; i < k/2; i++ {
	//	nums[i], nums[k-i-1] = nums[k-i-1], nums[i]
	//}
	//for i := k; i < (len(nums)+k)/2; i++ {
	//	nums[i], nums[len(nums)+k-i-1] = nums[len(nums)+k-i-1], nums[i]
	//}


	// time a little quick may cost a double space (may be under line's array doest copy)
	//last := nums[l-k:]
	//frist := nums[:l-k]
	//temp := append(last, frist...)
	//for i:=0; i<l; i++ {
	//	nums[i] = temp[i]
	//}

}

func containsDuplicate(nums []int) bool {
	container := make(map[int]bool, len(nums))
	for _,v := range nums {
		e := container[v]
		if e {
			return true
		}
		container[v] = true
	}
	return false
}

// leetcode subject: explore/interview/card/top-interview-questions-easy/1/array/25/
// require:
// time-complexity: O(n)
// space-complexity: O(1) , no exceptional space
func singleNumber(nums []int) int {

	// leetcode's some god's solution
	// keynote:
	// the input only one element occurs one, each others occurs twice
	// 1 ^ 1 ^ 3 = 3
	var res int

	for _,v := range nums {
		res ^= v
	}

	return res

	// second time
	//var result,isDuplicate int
	//l := len(nums)
	//
	//for i:=0; i<l; i++ {
	//	isDuplicate = 0
	//	for j := 0; j < l; j++ {
	//		if nums[j] == nums[i] {
	//			isDuplicate += 1
	//		}
	//	}
	//
	//	if isDuplicate == 1 {
	//		result = nums[i]
	//		break
	//	}
	//}

	// first time
	//var result int
	//m := make(map[int]int, len(nums))
	//
	//for _,v := range nums {
	//	m[v] += 1
	//}
	//
	//for i,k := range m {
	//	if k == 1 {
	//		return i
	//	}
	//}
	//return result
}
