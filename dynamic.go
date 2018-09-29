package leetcode

import (
	"fmt"
)

// leetcode: /explore/interview/card/top-interview-questions-easy/23/dynamic-programming/54/
// note:
// 类似斐波那契数列
// key note 当台阶n大于2时，人必定时从n-1台阶登一步到达n或者从n-2登2步到达n，就 num[n] = num[n-1] + num[n-2]
// 个人表示貌似跟动态规划不沾边
// 不过可以采用动态规划优化空间的策略，优化此算法，即只保留 n-1 与 n-2 的状态，不用使用数组

func ClimbStairs(n int) int {
	if n == 1 {
		return 1
	}

	if n == 2 {
		return 2
	}

	one := 1
	two := 2
	fn := 0
	for i:=2; i<n; i++ {
		fn = one + two
		one = two
		two = fn
	}
	return fn
	//arrs := make([]int, 2)
	//arrs[0] = 1
	//arrs[2] = 2
	//for i:=2; i<n; i++ {
	//	arrs = append(arrs, arrs[i-1] + arrs[i+1])
	//}
	//return arrs[n-1]
}
func MaxProfitOne(prices []int) int {

	// leetcode God's algorithm

	//if len(prices) <= 1 {
	//	return 0
	//}
	//gain:=0
	//buyPrice := prices[0]
	//sellPrice := prices[0]
	//for _, v := range prices {
	//	if v > sellPrice {
	//		sellPrice = v
	//	}
	//	if sellPrice - buyPrice >gain{
	//		gain = sellPrice - buyPrice
	//	}
	//	if v < buyPrice {
	//		buyPrice = v
	//		sellPrice = v
	//	}
	//}
	//return gain

	if len(prices) <= 1{
		return 0
	}
	gaps := make([]int, len(prices) - 1)
	for i:=0;i<len(gaps);i++{
		gaps[i] = prices[i+1]-prices[i]
	}
	fmt.Println(gaps)
	ans := 0
	sum := 0
	for i:=0;i<len(gaps);i++{
		sum += gaps[i]
		ans = max(ans, sum)
		sum = max(sum, 0)
	}
	return ans


	//l := len(prices)
	//maxProfit := 0
	//for j:=0; j<l; j++ {
	//	for i:=j+1; i<l; i++ {
	//		if prices[i] > prices[j] {
	//			maxProfit = max(maxProfit, prices[i]-prices[j])
	//		}
	//	}
	//}
	//
	//return maxProfit
}

func max(first,second int) int {
	if first > second {
		return  first
	}
	return second
}

func MaxSubArray(nums []int) int {

	// leetcode God's algorithm

	if len(nums)==0{
		return 0
	}

	maxSum:=nums[0]
	nowSum:=nums[0]

	for i,v:= range nums{
		if i==0{
			continue
		}

		if v>0{
			if nowSum>=0{
				nowSum+=v
			}else{
				nowSum = v
			}
			if nowSum >maxSum{
				maxSum = nowSum
			}
		}else{
			if v>nowSum{
				nowSum = v
			}else{
				nowSum+=v
			}
			if nowSum >maxSum{
				maxSum = nowSum
			}
		}
	}
	return maxSum

	//mx := math.MinInt64
	//for i:=0; i< len(nums); i++ {
	//	res := 0
	//	res += nums[i]
	//	mx = max(mx, res)
	//	for j:=i+1; j<len(nums); j++ {
	//		res += nums[j]
	//		mx = max(mx, res)
	//	}
	//}
	//return mx
}

func Robs(nums []int) int {

	// leetcode algorithm

	if len(nums)==0{
		return 0
	}

	if len(nums)==1{
		return nums[0]
	}

	m := make([]int,2)

	m[0]=nums[0]
	if nums[0]>nums[1]{
		m[1]=nums[0]
	}else {
		m[1]=nums[1]
	}

	for i:=2;i<len(nums);i++{
		if m[i-1]>(nums[i]+m[i-2]){
			m=append(m,m[i-1])
		}else {
			m=append(m,nums[i]+m[i-2])
		}
	}
	return m[len(nums)-1]

	//if len(nums) == 0 { // keynote
	//	return 0
	//}
	//
	//if len(nums) == 1 {
	//	return nums[0]
	//}
	//
	//if len(nums) == 2 {
	//	return max(nums[0], nums[1])
	//}
	//
	//sum := make([]int, 2)
	//sum[0] = nums[0]
	//sum[1] = max(nums[0], nums[1]) // keynote
	//for i:=2; i<len(nums); i++ {
	//	sum = append(sum, max(nums[i] + sum[i-2], sum[i-1])) // keynote
	//}
	//return max(sum[len(sum) - 1], sum[len(sum) - 2])
}

