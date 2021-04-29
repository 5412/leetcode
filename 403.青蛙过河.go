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
	// dp[i][speed]：表示能否以speed的速度，到达第i个石头
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

func canCross2(stones []int) bool {
	n := len(stones)
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}
	dp[0][0] = true
	for i := 1; i < n; i++ {
		if stones[i]-stones[i-1] > i {
			return false
		}
	}
	for i := 1; i < n; i++ {
		for j := i - 1; j >= 0; j-- {
			k := stones[i] - stones[j]
			if k > j+1 {
				break
			}
			dp[i][k] = dp[j][k-1] || dp[j][k] || dp[j][k+1]
			if i == n-1 && dp[i][k] {
				return true
			}
		}
	}
	return false
}

func canCross3(stones []int) bool {
	if stones[1] > 1 {
		return false
	}
	s := make([]int, len(stones)-1)
	for i := 0; i < len(stones)-1; i++ {
		s[i] = stones[i+1] - stones[i]
	}
	last := s[len(s)-1]
	sum := 0
	for i := len(s) - 2; i >= 0; i-- {
		if i == 0 && last > 2 {
			return false
		}
		if s[i]+sum+1 >= last && s[i]+sum-1 <= last {
			last = s[i] + sum
			sum = 0
		} else if s[i]+s[i-1]+sum >= last {
			last += s[i]
		} else {
			sum += s[i]
		}
	}
	return sum == 0
}

// @lc code=end
