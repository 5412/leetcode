package hard

import "sort"

/**
给定长度为 n 的整数数组 nums，其中 n > 1，返回输出数组 output ，其中 output[i] 等于 nums 中除 nums[i] 之外其余各元素的乘积。
示例:
输入: [1,2,3,4]
输出: [24,12,8,6]
说明: 请不要使用除法，且在 O(n) 时间复杂度内完成此题。
进阶：
你可以在常数空间复杂度内完成这个题目吗？（ 出于对空间复杂度分析的目的，输出数组不被视为额外空间。）
 */
func ProductExceptSelf(nums []int) []int {
	out := make([]int, len(nums))
	out[0] = 1
	if len(nums) == 2 {
		return []int{nums[1], nums[0]}
	}
	for i:=1; i<len(nums);i++ {
		out[i] = out[0] * nums[0]
		out[0] *= nums[i]
	}
	outAfter := make([]int, len(nums))
	end := len(nums) - 1
	outAfter[end] = 1
	for j:=len(nums) - 2; j>0;j-- {
		outAfter[j] = outAfter[end]  * nums[end]
		outAfter[end] *= nums[j]
	}
	for i:=1; i<end;i++ {
		out[i] = out[i] * outAfter[i]
	}
	out[end] = outAfter[end]
	return out
}

func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums))
	k := 1
	for i:=0; i<len(nums); i++ {
		result[i] = k
		k *= nums[i]
	}
	k = 1
	for j:=len(nums)-1; j>=0; j-- {
		result[j] *= k
		k *= nums[j]
	}
	return result
}

/**
给定一个包含 m x n 个元素的矩阵（m 行, n 列），请按照顺时针螺旋顺序，返回矩阵中的所有元素。
示例 1:
输入:
[
 [ 1, 2, 3 ],
 [ 4, 5, 6 ],
 [ 7, 8, 9 ]
]
输出: [1,2,3,6,9,8,7,4,5]
示例 2:
输入:
[
  [1, 2, 3, 4],
  [5, 6, 7, 8],
  [9,10,11,12]
]
输出: [1,2,3,4,8,12,11,10,9,5,6,7]
*/
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

/**
给定四个包含整数的数组列表 A , B , C , D ,计算有多少个元组 (i, j, k, l) ，使得 A[i] + B[j] + C[k] + D[l] = 0。
为了使问题简单化，所有的 A, B, C, D 具有相同的长度 N，且 0 ≤ N ≤ 500 。所有整数的范围在 -228 到 228 - 1 之间，最终结果不会超过 231 - 1 。
例如:
输入:
A = [ 1, 2]
B = [-2,-1]
C = [-1, 2]
D = [ 0, 2]
输出:
2
解释:
两个元组如下:
1. (0, 0, 0, 1) -> A[0] + B[0] + C[0] + D[1] = 1 + (-2) + (-1) + 2 = 0
2. (1, 1, 0, 0) -> A[1] + B[1] + C[0] + D[0] = 2 + (-1) + (-1) + 0 = 0
*/
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
	for i := 0; i < len(C); i += 1 {
		i1 := i + 1
		for ; i1 < len(C) && C[i1] == C[i]; i1 += 1 {
		}
		for j := 0; j < len(D); j += 1 {
			j1 := j + 1
			for ; j1 < len(D) && D[j1] == D[j]; j1 += 1 {
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

/**
给定 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0)。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
说明：你不能倾斜容器，且 n 的值至少为 2。
图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。

示例:
输入: [1,8,6,2,5,4,8,3,7]
输出: 49
*/

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
