package medium

import (
	"fmt"
	"math"
)

/**
给定一个非负整数数组，你最初位于数组的第一个位置。
数组中的每个元素代表你在该位置可以跳跃的最大长度。
判断你是否能够到达最后一个位置。
示例 1:

输入: [2,3,1,1,4]
输出: true
解释: 我们可以先跳 1 步，从位置 0 到达 位置 1, 然后再从位置 1 跳 3 步到达最后一个位置。
示例 2:

输入: [3,2,1,0,4]
输出: false
解释: 无论怎样，你总会到达索引为 3 的位置。但该位置的最大跳跃长度是 0 ， 所以你永远不可能到达最后一个位置。
 */
func CanJump(nums []int) bool {
	l := len(nums)
	maxJ := optCanJump(nums, l - 1)
	fmt.Println(maxJ)
	return maxJ >= l - 1
}

func optCanJump(nums []int, i int) int {
	if i == 0 {
		return nums[i]
	}
	beforeJ := optCanJump(nums, i - 1)
	if beforeJ < i {
		return beforeJ
	}
	return max(beforeJ, i + nums[i] )
}

// 倒序，如果能到达最后一个，肯定能到达前一个
func canJump(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}

	lens := len(nums)
	lp := lens - 1
	for i := lens - 1;i >= 0;i-- {
		if i + nums[i] >= lp {
			lp = i
		}
	}

	return lp == 0
}

/**
一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。
机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。
问总共有多少条不同的路径？
例如，上图是一个7 x 3 的网格。有多少可能的路径？
说明：m 和 n 的值均不超过 100。
示例 1:
输入: m = 3, n = 2
输出: 3
解释:
从左上角开始，总共有 3 条路径可以到达右下角。
1. 向右 -> 向右 -> 向下
2. 向右 -> 向下 -> 向右
3. 向下 -> 向右 -> 向右
必须向右2次向下一次
math C(3,1) = 3
 */

func UniquePaths(m int, n int) int {
	//return optUniquePaths(m,n,n-1,m-1)
	res := make([]int, n)
	res[0] = 1
	for i:=0; i<m; i++ {
		for j:=1;j<n; j++ {
			res[j] += res[j-1]
		}
	}
	return res[n-1]
}

func optUniquePaths(m, n, x, y int) int {
	if x < 0 || y < 0 {
		return 0
	}
	if x == 0 && y == 0 {
		return 1
	}
	return optUniquePaths(m,n,x-1, y) + optUniquePaths(m,n,x,y-1)
}

/**
给定不同面额的硬币 coins 和一个总金额 amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。如果没有任何一种硬币组合能组成总金额，返回 -1。
dp[amount] = min(dp[amount-coins[j] + 1)
示例 1:

输入: coins = [1, 2, 5], amount = 11
输出: 3
解释: 11 = 5 + 5 + 1
示例 2:

输入: coins = [2], amount = 3
输出: -1
说明:
你可以认为每种硬币的数量是无限的。
 */
func CoinChange(coins []int, amount int) int {
	res := make(map[int]int)
	for i:=0; i<amount+1; i++ {
		res[i] = math.MaxInt64 - 1
	}
	res[0] = 0
	for _, coin := range coins {
		for i:=coin; i<amount + 1; i++ {
			if i - coin < 0 {
				continue
			}
			res[i] = selectMin(res[i], res[i-coin] + 1)
		}
	}
	if res[amount] < math.MaxInt64 - 2 {
		return res[amount]
	}
	return  0
}

func selectMin(x, y int) int {
	if x < y {
		return x
	}
	return y
}

//本题归结为完全背包问题
//直接套用完全背包的套路就行
func coinChange(coins []int, amount int) int {
	dp:=make([]int,amount+1)
	//初始值
	for i:=0;i<=amount;i++{
		dp[i]=amount+1
	}
	dp[0]=0

	for i:=1;i<=len(coins);i++{
		for j:=0;j<=amount;j++{
			if j >= coins[i-1]{
				dp[j]=selectMin(dp[j],dp[j-coins[i-1]]+1)
			}
		}
	}
	if dp[amount] > amount{
		return -1
	}else{
		return dp[amount]
	}
}

/**
  Longest Increasing Subsequence
给定一个无序的整数数组，找到其中最长上升子序列的长度。
示例:

输入: [10,9,2,5,3,7,101,18]
输出: 4
解释: 最长的上升子序列是 [2,3,7,101]，它的长度是 4。
说明:
dp[i] = dp[i-1] + 1 nums[i] > nums[i-1]
dp[i] = dp[i-1] nums[i] <= nums[i-1]
可能会有多种最长上升子序列的组合，你只需要输出对应的长度即可。
你算法的时间复杂度应该为 O(n2) 。
进阶: 你能将算法的时间复杂度降低到 O(n log n) 吗?
 */
func LengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	for i:=0; i<len(nums); i++ {
		dp[i] = 1
	}
	for j:=1; j<len(nums); j++ {
		for i:=0; i<j; i++ {
			if nums[i] < nums[j] {
				dp[j] = max(dp[i] + 1, dp[j])
			}
		}
	}
	res := 1
	for _,v := range dp {
		res = max(res, v)
	}
	return res
}

func lengthOfLIS(nums []int) int {
	var minTailLis = make([]int, 0, len(nums))

	for i := 0; i < len(nums); i++ {
		minTailLis = fillNum(minTailLis, nums[i])
	}

	return len(minTailLis)
}

func fillNum(minTailLis []int, num int) []int {
	if len(minTailLis) == 0 {
		minTailLis = append(minTailLis, num)
	}
	if num > minTailLis[len(minTailLis)-1] {
		minTailLis = append(minTailLis, num)
		return minTailLis
	}
	index := binarySearchFirstMore(minTailLis, 0, len(minTailLis)-1, num)
	if index != -1 {
		minTailLis[index] = num
	}
	//fmt.Println(minTailLis)
	return minTailLis
}

func binarySearchFirstMore(minTailLis []int, l, r, num int) int {

	if l > r {
		return -1
	}

	midIndex := (l + r) / 2
	midValue := minTailLis[midIndex]
	if midValue <= num {
		return binarySearchFirstMore(minTailLis, midIndex+1, r, num)
	}
	if midIndex == 0 || minTailLis[midIndex-1] < num {
		return midIndex
	}
	if minTailLis[midIndex-1] == num {
		return -1
	}
	return binarySearchFirstMore(minTailLis, l, midIndex-1, num)
}
