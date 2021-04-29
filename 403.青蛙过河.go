package leetcode

/*
 * @lc app=leetcode.cn id=403 lang=golang
 *
 * [403] 青蛙过河
 */

// @lc code=start
func canCross(stones []int) bool {
	stoneMap := make(map[int]int)
	for idx, s := range stones {
		stoneMap[s] = idx
	}
	//dp 标识第 i 个石子上一次跳 x 步是否能够到达
	dp := make([]map[int]bool, len(stones)-1)
	for i := range dp {
		dp[i] = make(map[int]bool)
	}
	var dpf func(i, lastStep int) bool
	dpf = func(i, lastStep int) (res bool) {
		if i == len(stones)-1 {
			return true
		}
		if res, ok := dp[i][lastStep]; ok {
			return res
		}
		defer func() {
			dp[i][lastStep] = res
		}()
		for curStep := lastStep - 1; curStep <= lastStep+1; curStep++ {
			if curStep <= 0 {
				continue
			}
			nextPostion := stoneMap[stones[i]+curStep]
			if nextPostion != 0 && dpf(nextPostion, curStep) {
				return true
			}
		}
		return false
	}
	return dpf(0, 0)
}

// @lc code=end
