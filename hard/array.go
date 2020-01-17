package hard

import (
	"fmt"
	"sort"
)

// ProductExceptSelf 给定长度为 n 的整数数组 nums，其中 n > 1，返回输出数组 output ，其中 output[i] 等于 nums 中除 nums[i] 之外其余各元素的乘积。
// 示例:
// 输入: [1,2,3,4]
// 输出: [24,12,8,6]
// 说明: 请不要使用除法，且在 O(n) 时间复杂度内完成此题。
// 进阶：
// 你可以在常数空间复杂度内完成这个题目吗？（ 出于对空间复杂度分析的目的，输出数组不被视为额外空间。）
func ProductExceptSelf(nums []int) []int {
	out := make([]int, len(nums))
	out[0] = 1
	if len(nums) == 2 {
		return []int{nums[1], nums[0]}
	}
	for i := 1; i < len(nums); i++ {
		out[i] = out[0] * nums[0]
		out[0] *= nums[i]
	}
	outAfter := make([]int, len(nums))
	end := len(nums) - 1
	outAfter[end] = 1
	for j := len(nums) - 2; j > 0; j-- {
		outAfter[j] = outAfter[end] * nums[end]
		outAfter[end] *= nums[j]
	}
	for i := 1; i < end; i++ {
		out[i] = out[i] * outAfter[i]
	}
	out[end] = outAfter[end]
	return out
}

func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums))
	k := 1
	for i := 0; i < len(nums); i++ {
		result[i] = k
		k *= nums[i]
	}
	k = 1
	for j := len(nums) - 1; j >= 0; j-- {
		result[j] *= k
		k *= nums[j]
	}
	return result
}

// 给定一个包含 m x n 个元素的矩阵（m 行, n 列），请按照顺时针螺旋顺序，返回矩阵中的所有元素。
// 示例 1:
// 输入:
// [
//  [ 1, 2, 3 ],
//  [ 4, 5, 6 ],
//  [ 7, 8, 9 ]
// ]
// 输出: [1,2,3,6,9,8,7,4,5]
// 示例 2:
// 输入:
// [
//   [1, 2, 3, 4],
//   [5, 6, 7, 8],
//   [9,10,11,12]
// ]
// 输出: [1,2,3,4,8,12,11,10,9,5,6,7]
func spiralOrder(matrix [][]int) []int {
	m := len(matrix)
	n := len(matrix[0])
	out := make([]int, 0, m*n)
	x := 0
	i := 0
	j := 0
	minI := 0
	minJ := 0
	maxI := m - 1
	maxJ := n - 1
	for minI <= maxI && minJ <= maxJ {
		switch x % 4 {
		case 0:
			for j <= maxJ {
				out = append(out, matrix[i][j])
				j++
			}
			x++
			i++
			j--
			minI++
		case 1:
			for i <= maxI {
				out = append(out, matrix[i][j])
				i++
			}
			i--
			j--
			x++
			maxJ--
		case 2:
			for j >= minJ {
				out = append(out, matrix[i][j])
				j--
			}
			j++
			x++
			i--
			maxI--
		case 3:
			for i >= minI {
				out = append(out, matrix[i][j])
				i--
			}
			i++
			x++
			minJ++
			j++
		}
	}
	return out
}

//FourSumCount 给定四个包含整数的数组列表 A , B , C , D ,计算有多少个元组 (i, j, k, l) ，使得 A[i] + B[j] + C[k] + D[l] = 0。
// 为了使问题简单化，所有的 A, B, C, D 具有相同的长度 N，且 0 ≤ N ≤ 500 。所有整数的范围在 -228 到 228 - 1 之间，最终结果不会超过 231 - 1 。
// 例如:
// 输入:
// A = [ 1, 2]
// B = [-2,-1]
// C = [-1, 2]
// D = [ 0, 2]
// 输出:
// 2
// 解释:
// 两个元组如下:
// 1. (0, 0, 0, 1) -> A[0] + B[0] + C[0] + D[1] = 1 + (-2) + (-1) + 2 = 0
// 2. (1, 1, 0, 0) -> A[1] + B[1] + C[0] + D[0] = 2 + (-1) + (-1) + 0 = 0
func FourSumCount(A []int, B []int, C []int, D []int) int {
	mapAB := make(map[int]int, len(A)*len(B))
	mapCD := make(map[int]int, len(C)*len(D))
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(B); j++ {
			v := A[i] + B[j]
			_, ok := mapAB[v]
			if ok {
				mapAB[v]++
			} else {
				mapAB[v] = 1
			}
		}
	}
	for i := 0; i < len(C); i++ {
		for j := 0; j < len(D); j++ {
			v := C[i] + D[j]
			_, ok := mapCD[v]
			if ok {
				mapCD[v]++
			} else {
				mapCD[v] = 1
			}
		}
	}
	res := 0
	for i, v := range mapAB {
		target := 0 - i
		t, ok := mapCD[target]
		if ok {
			res += v * t
		}
	}
	return res
}

func fourSumCount(A []int, B []int, C []int, D []int) int {
	cdsums := make(map[int]int, len(C)*len(D))

	sort.Ints(C)
	sort.Ints(D)

	// O(N^2)
	for i := 0; i < len(C); i++ {
		i1 := i + 1
		for ; i1 < len(C) && C[i1] == C[i]; i1++ {
		}
		for j := 0; j < len(D); j++ {
			j1 := j + 1
			for ; j1 < len(D) && D[j1] == D[j]; j1++ {
			}

			sum := C[i] + D[j]
			cdsums[sum] += (i1 - i) * (j1 - j)
			j = j1 - 1
		}

		i = i1 - 1
	}

	// O(N^2)
	count := 0
	for i := range A {
		for j := range B {
			count += cdsums[0-A[i]-B[j]]
		}
	}

	return count
}

//MaxArea 给定 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0)。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
// 说明：你不能倾斜容器，且 n 的值至少为 2。
// 图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。
// 示例:
// 输入: [1,8,6,2,5,4,8,3,7]
// 输出: 49
func MaxArea(height []int) int {
	res, i, j := 0, 0, len(height)-1
	for i < j {
		res = max(res, (j-i)*min(height[i], height[j]))
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxArea(height []int) int {
	var (
		i, j             = 0, len(height) - 1
		h, area, maxArea int
	)
	for {
		if i == j {
			break
		}
		if height[i] > height[j] {
			h = height[j]
			j--
		} else {
			h = height[i]
			i++
		}
		area = h * (j - i + 1)
		if area > maxArea {
			maxArea = area
		}
	}
	return maxArea
}

// 根据百度百科，生命游戏，简称为生命，是英国数学家约翰·何顿·康威在1970年发明的细胞自动机。
// 给定一个包含 m × n 个格子的面板，每一个格子都可以看成是一个细胞。每个细胞具有一个初始状态 live（1）即为活细胞， 或 dead（0）即为死细胞。每个细胞与其八个相邻位置（水平，垂直，对角线）的细胞都遵循以下四条生存定律：
// 如果活细胞周围八个位置的活细胞数少于两个，则该位置活细胞死亡；
// 如果活细胞周围八个位置有两个或三个活细胞，则该位置活细胞仍然存活；
// 如果活细胞周围八个位置有超过三个活细胞，则该位置活细胞死亡；
// 如果死细胞周围正好有三个活细胞，则该位置死细胞复活；
// 根据当前状态，写一个函数来计算面板上细胞的下一个（一次更新后的）状态。下一个状态是通过将上述规则同时应用于当前状态下的每个细胞所形成的，其中细胞的出生和死亡是同时发生的。
// 示例:
// 输入:
// [
//   [0,1,0],
//   [0,0,1],
//   [1,1,1],
//   [0,0,0]
// ]
// 输出:
// [
//   [0,0,0],
//   [1,0,1],
//   [0,1,1],
//   [0,1,0]
// ]
// 进阶:
// 你可以使用原地算法解决本题吗？请注意，面板上所有格子需要同时被更新：你不能先更新某些格子，然后使用它们的更新后的值再更新其他格子。
// 本题中，我们使用二维数组来表示面板。原则上，面板是无限的，但当活细胞侵占了面板边界时会造成问题。你将如何解决这些问题？
// 0和1的二进制码分别是0000和0001，表示细胞处于dead状态和live状态。如果试图将下一个状态也记录在这里面，就像
// 0000表示当前状态是dead并且下一个状态是dead，值为0
// 0010表示当前状态是dead并且下一个状态是live，值为2
// 0001表示当前状态是live并且下一个状态是dead，值为1
// 0011表示当前状态是live并且下一个状态是live，值为3
// 这样，在第二遍遍历时只需要将每个细胞的值右移一位，就从当前状态变为下一个状态了。
func gameOfLife(board [][]int) {
	m := len(board)
	if m == 0 {
		return
	}
	n := len(board[0])
	tempBoard := make([][]int, m)
	for i := 0; i < m; i++ {
		tmp := make([]int, n)
		tempBoard[i] = tmp
		for j := 0; j < n; j++ {
			tempBoard[i][j] = gameOflifeHelper(board, i, j, m, n)
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			switch {
			case board[i][j] == 1:
				if tempBoard[i][j] < 2 || tempBoard[i][j] > 3 {
					board[i][j] = 0
				}
			case board[i][j] == 0:
				if tempBoard[i][j] == 3 {
					board[i][j] = 1
				}
			}
		}
	}
}

func gameOflifeHelper(board [][]int, i, j, m, n int) int {
	res := 0
	if i < 0 || i >= m {
		return res
	}

	if j < 0 || j >= n {
		return res
	}
	if i-1 >= 0 && j-1 >= 0 {
		res += board[i-1][j-1]
	}
	if i-1 >= 0 {
		res += board[i-1][j]
	}
	if i-1 >= 0 && j+1 < n {
		res += board[i-1][j+1]
	}

	if j-1 >= 0 {
		res += board[i][j-1]
	}
	if j+1 < n {
		res += board[i][j+1]
	}

	if i+1 < m && j-1 >= 0 {
		res += board[i+1][j-1]
	}
	if i+1 < m {
		res += board[i+1][j]
	}
	if i+1 < m && j+1 < n {
		res += board[i+1][j+1]
	}

	return res
}

// FirstMissingPositive  第一个缺失的正数
// 给定一个未排序的整数数组，找出其中没有出现的最小的正整数。
// 示例 1:
// 输入: [1,2,0]
// 输出: 3
// 示例 2:
// 输入: [3,4,-1,1]
// 输出: 2
// 示例 3:
// 输入: [7,8,9,11,12]
// 输出: 1
// 说明:
// 你的算法的时间复杂度应为O(n)，并且只能使用常数级别的空间。
func FirstMissingPositive(nums []int) int {
	for i, v := range nums {
		if v <= 0 || v >= len(nums) {
			continue
		}
		for nums[v-1] != v {
			nums[v-1], nums[i] = nums[i], nums[v-1]
			v = nums[i]
			if v <= 0 || v >= len(nums) {
				break
			}
		}
	}

	for i, v := range nums {
		if v != i+1 {
			return i + 1
		}
	}
	return len(nums) + 1
}

func abs(n int) int {
	if n <= 0 {
		return 0 - n
	}
	return n
}

//FirstMissingPositive1 第一个确实的正整数 leetcode版
func FirstMissingPositive1(nums []int) int {
	found1 := false
	length := len(nums)
	for i := 0; i < length; i++ {
		if nums[i] == 1 {
			found1 = true
		} else if nums[i] > length || nums[i] <= 0 {
			nums[i] = 1
		}
	}
	if !found1 {
		return 1
	}
	for i := 0; i < length; i++ {
		num := abs(nums[i])
		if nums[num-1] > 0 {
			nums[num-1] = -nums[num-1]
		}
	}
	for i := 0; i < length; i++ {
		if nums[i] > 0 {
			return i + 1
		}
	}
	return length + 1
}

// LongestConsecutive 最长连续序列
// 给定一个未排序的整数数组，找出最长连续序列的长度。
// 要求算法的时间复杂度为 O(n)。
// 示例:
// 输入: [100, 4, 200, 1, 3, 2]
// 输出: 4
// 解释: 最长连续序列是 [1, 2, 3, 4]。它的长度为 4。
// 针对O(n)时间复杂度的实现思路：遍历nums[]数组，利用Map存储元素nums[i]的值以及其所在连续序列的长度，此时基本只有两种情况：
// 数组中出现过元素nums[i]-1或nums[i]+1，意味着当前元素可以归入左或右序列，那么此时假如左右序列的长度分别为left、right，那么显然加入nums[i]后，这整段序列的长度为 1+left+right，而由于这一整段序列中，只可能在左右两端扩展，所以只需要更新左右两端的value值即可。
// 数组中未出现过元素nums[i]-1或nums[i]+1，意味着当前元素所在的连续序列就是自身（只有自己一个元素）。
func LongestConsecutive(nums []int) int {
	res := 0
	longMap := make(map[int]int, len(nums)/2)
	for _, v := range nums {
		if _, ok := longMap[v]; ok {
			continue
		}
		left := longMap[v-1]
		right := longMap[v+1]
		longMap[v] = 1 + right + left
		res = max(res, longMap[v])
		if left != 0 {
			longMap[v-left] = longMap[v]
		}
		if right != 0 {
			longMap[v+right] = longMap[v]
		}
	}
	fmt.Println(longMap)
	return res
}
