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


// leetcode: explore/interview/card/top-interview-questions-easy/1/array/26/

func intersect(nums1 []int, nums2 []int) []int {
	//output := make([]int, 0)
	var output []int
	//var m map[int]int
	m := make(map[int]int, 0)
	for _,v1 := range nums1 {
		m[v1] += 1
	}

	for _,v := range nums2 {
		if m[v] > 0 {
			output = append(output, v)
			m[v]--
		}
	}
	return output
}

func plusOne(digits []int) []int {
	//var res []int
	l := len(digits)
	for i:=l; i>0; i-- {
		tail := digits[i-1]
		tail++
		switch tail {
		case 10:
			tail = 0
		}
		digits[i-1] = tail

		if tail > 0 {
			break
		}
		// last loop
		if i == 1 {
			digits = append([]int{1}, digits...)
		}
	}

	return digits





	// Don't work while the digits is to large
	//number := 0
	//for _,v := range digits  {
	//	number = number *  10 + v
	//}
	//number++
	//output := make([]int,0)
	//for  {
	//	if number == 0 {
	//		break
	//	}
	//	digit := number % 10
	//	output = append([]int{digit}, output...)
	//	number = (number - digit) / 10
	//}
	//return output
}

// leetcode: explore/interview/card/top-interview-questions-easy/1/array/28/
func moveZeros(nums []int) {
	l := len(nums)
	for i := l - 1; i >= 0; i-- {
		if nums[i] == 0 {
			for j := i; j < l-1; j++ {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				//nums[j] = nums[j+1]
			}
			//nums[l-1] = 0
			l--
		}
	}
}
