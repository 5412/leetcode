package medium

import "fmt"

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