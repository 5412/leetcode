// Copyright (c) 2019. weixiong piao, All Right Reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package medium

import (
	"fmt"
	"sort"
)

func sortColor(nums []int) {
	if len(nums) < 2 {
		return
	}
	if nums[0] > nums[1] {
		nums[0], nums[1] = nums[1], nums[0]
	}
	min := nums[0]
	max := nums[1]
	minIndex := 1
	maxIndex := 1


	for i:=2;i < len(nums); i++ {
		fmt.Println(i, minIndex, maxIndex, nums[i],  min, max, nums)
		if nums[i] < min {
			min = nums[i]
			nums[0], nums[i] = nums[i], nums[0]
			if nums[maxIndex]  > nums[i] {
				nums[maxIndex],  nums[i]  = nums[i], nums[maxIndex]
				maxIndex++
			}
			minIndex = 1
		} else if nums[i] !=  nums[i-1] && nums[i] == min {
			nums[minIndex], nums[i] =  nums[i], nums[minIndex]
			if nums[i] < nums[maxIndex] {
				nums[maxIndex], nums[i] = nums[i], nums[maxIndex]
			}
			minIndex++
			maxIndex++
		} else  if  nums[i]  ==  min  {
			minIndex++
		}else if nums[i] > max {
				max = nums[i]
				if nums[minIndex] ==  min {
					minIndex  = i
				}
				maxIndex = i
		} else if nums[i] < max && nums[i] > min {
			nums[maxIndex],  nums[i] = nums[i], nums[maxIndex]
			maxIndex++
		} else  if nums[i]  > nums[i-1] {
			if nums[i] >=  nums[minIndex] {
				minIndex = i
			}

			if nums[i] > nums[maxIndex] {
				i = maxIndex
			}
		}
	}
}

func SortColor(nums []int) {
	var (
		i,j int
	)

	for f:=0; f < len(nums); f++ {
		switch nums[f] {
		case 0:
			nums[f] =  2
			nums[j] = 1
			nums[i] = 0
			i++
			j++
			//
		case 1:
			nums[f] =  2
			nums[j] =  1
			j++
		case 2:
			//nums[j]  = 2
		}
	}
}

// leetcode God's algorithm
func sortColors(a []int) {
	i, j, k := 0, 0, len(a)-1

	// for 循环中， nums[i:j] 中始终全是 1
	// 循环结束后，
	// nums[:i] 中全是 0
	// nums[j:] 中全是 2
	for j <= k {
		switch a[j] {
		case 0:
			a[i], a[j] = a[j], a[i]
			i++
			j++
		case 1:
			j++
		case 2:
			a[j], a[k] = a[k], a[j]
			k--
		}
	}

}

func TopKFrequent(nums []int, k int) []int {
	result := make([]int, 0, k)
	m := make(map[int]int)
	mk := make(map[int][]int)
	for _,j := range nums {
		m[j] = m[j] + 1
		//fmt.Println(j, m[j])
		topKHelper(&result,  j, m[j], m)
	}

	for k,v := range m {
		if mk[v] == nil {
			mk[v] = []int{k}
		} else {
			mk[v] = append(mk[v], k)
		}
	}
	tmp := make([]int, 0, len(mk))

	for k,_ := range mk  {
		tmp =  append(tmp, k)
	}

	sort.Ints(tmp)

	for i:=0; i<len(tmp); i++ {
		result = append(result, mk[tmp[i]]...)
	}

	return result[len(result) - k :]
}

func topKHelper(result *[]int, index, num int, m map[int]int)  {
	if len(*result) != cap(*result)  &&  num ==  1 {
		//fmt.Println(index)
		*result = append(*result, index)
	} else {
		tmp := *result
		exist := false
		for _,j := range tmp {
			if j == index {
				exist = true
			}
		}
		if !exist {
			for i,j := range tmp {
				if j != index && m[j] < num {
					tmp[i] = index
					break
				}
			}
		}
	}
}


// leetcode's God algorithm
func topKFrequent(nums []int, k int) []int {
	res := make([]int, 0, k)
	// 统计每个数字出现的次数
	rec := make(map[int]int, len(nums))
	for _, n := range nums {
		rec[n]++
	}
	// 对出现次数进行排序
	counts := make([]int, 0, len(rec))
	for _, c := range rec {
		counts = append(counts, c)
	}
	sort.Ints(counts)
	// min 是 前 k 个高频数字的底线
	min := counts[len(counts)-k]

	// 收集所有　不低于　底线的数字
	for n, c := range rec {
		if c >= min {
			res = append(res, n)
		}
	}

	return res
}

/**
给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。

你的算法时间复杂度必须是 O(log n) 级别。

如果数组中不存在目标值，返回 [-1, -1]。

示例 1:

输入: nums = [5,7,7,8,8,10], target = 8
输出: [3,4]
示例 2:

输入: nums = [5,7,7,8,8,10], target = 6
输出: [-1,-1]
 */
func SearchRange(nums []int, target int) []int {

	index := searchRangeHelper(nums, 0, len(nums)-1, target)
	begin := -1
	end := -1
	if index > -1 {
		begin = index
		end = index
		for index > -1 && nums[index] == target {
			begin = index
			index--
		}
		index = end
		for index < len(nums) && nums[index] == target {
			end = index
			index++
		}
	}
	return []int{begin, end}
}

// 二分查找目标值 找到返回下标，找不到返回 -1
func searchRangeHelper(nums []int,begin, end, target int) int {
	if len(nums) == 0 {
		return -1
	}
	if nums[begin] > target {
		return -1
	}

	if nums[end] < target {
		return -1
	}

	if begin == end && nums[begin] != target{
		return -1
	}
	middle := (end + begin) / 2
	if nums[middle] > target {
		return searchRangeHelper(nums, begin, middle - 1, target)
	} else if nums[middle] < target {
		return searchRangeHelper(nums, middle + 1, end, target)
	} else {
		return middle
	}
}

func searchRangeHelperFind(nums []int,begin, end, target int) int {
	index := -1
	for begin <= end {
		middle := ( begin + end ) / 2
		if nums[middle] == target {
			index = middle
			break
		}
		if nums[middle] < target {
			begin = middle + 1
		}
		if nums[middle] > target {
			end = middle - 1
		}
	}
	return  index
}

/**
给出一个区间的集合，请合并所有重叠的区间。

示例 1:

输入: [[1,3],[2,6],[8,10],[15,18]]
输出: [[1,6],[8,10],[15,18]]
解释: 区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
示例 2:

输入: [[1,4],[4,5]]
输出: [[1,5]]
解释: 区间 [1,4] 和 [4,5] 可被视为重叠区间。
 */

func Merge(inputs [][]int) [][]int {
	if len(inputs) == 0 {
		return nil
	}
	sort.Slice(inputs, func(i, j int) bool {
		return inputs[i][0] < inputs[j][0]
	})
	begin := inputs[0]
	result := make([][]int, 0)
	for i:=1; i<len(inputs); i++ {
		if inputs[i][0] <= begin[1] {
			if begin[1] < inputs[i][1] {
				begin[1] = inputs[i][1]
			}
		} else {
			result = append(result, begin)
			begin = inputs[i]
		}
	}
	result = append(result, begin)
	return result
}

// leetcode's God algorithm 原地计算
type TwoDArray [][]int

func (s TwoDArray) Len() int {
	return len(s)
}

func (s TwoDArray) Swap(i, j int) {
	s[i][0], s[j][0] = s[j][0], s[i][0]
	s[i][1], s[j][1] = s[j][1], s[i][1]
}

func (s TwoDArray) Less(i, j int) bool {
	return s[i][0] < s[j][0]
}
func merge(intervals [][]int) [][]int {
	_len := len(intervals)
	_index := 0
	if _len == 0 {
		return intervals
	}
	sort.Sort(TwoDArray(intervals))
	for i := 1; i < _len; i++ {
		if intervals[i][0] <= intervals[i - 1][1] {
			intervals[i][0] = intervals[i - 1][0]
			if intervals[i][1] < intervals[i - 1][1] {
				intervals[i][1] = intervals[i - 1][1]
			}
		} else {
			intervals[_index] = intervals[i - 1]
			_index++
		}
	}
	intervals[_index] = intervals[_len - 1]
	intervals = intervals[:_index + 1]
	return intervals
}

/**
假设按照升序排序的数组在预先未知的某个点上进行了旋转。

( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [5,6,7,0,1,2,4] )。

搜索一个给定的目标值，如果数组中存在这个目标值，则返回它的索引，否则返回 -1 。

你可以假设数组中不存在重复的元素。

你的算法时间复杂度必须是 O(log n) 级别。
 */
func SearchAndReverse(inputs []int, target int) int {
	left := 0
	right := len(inputs) - 1
	fmt.Println(inputs)
	for left <= right {
		mid := left + (right - left) / 2
		fmt.Println(left, right, mid)
		if inputs[mid] == target {
			return mid
		}
		if inputs[mid] >= inputs[left] { // mid 左边有序
			if inputs[mid] > target && inputs[left] <= target {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else { // mid 右边有序
			if inputs[mid] < target && inputs[right] >= target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	min := min(nums)

	if nums[min] == target {
		return min
	}

	// 数组没有旋转
	if min == 0 {
		return binarySearch(nums, 0, len(nums)-1, target)
	}

	// 在后一半升序数组中找
	if target < nums[0] {
		return binarySearch(nums, min, len(nums)-1, target)
	} else { // 在前一半升序数组中找
		return binarySearch(nums, 0, min, target)
	}
}

// 寻找旋转点，即数组的最小值
func min(nums []int) int {
	p1 := 0
	p2 := len(nums) - 1
	// 将mid设置为p1,一旦发现第一个数小于最后一个数，说明数组旋转了0位
	// 即没旋转，可以直接返回答案nums[p1],即第一个数
	mid := p1
	for nums[p1] >= nums[p2] {
		// 说明p1指向了第一个递增数组的尾部
		// p2指向了第二个递增数组的头部，即最小值
		if p2-p1 == 1 {
			mid = p2
			break
		}

		mid = (p1 + p2) / 2

		// 如果p1,mid,p2三者指向的数字相同，则没法判断最小值位于什么位置，只能遍历查找
		if nums[p1] == nums[mid] && nums[mid] == nums[p2] {
			res := p1
			for i := p1; i <= p2; i++ {
				if nums[res] > nums[i] {
					res = i
				}
			}
			return res
		}

		if nums[mid] >= nums[p1] { // 说明mid位于第一个递增数组
			p1 = mid
		} else if nums[mid] <= nums[p2] { // 说明mid位于第二个递增数组
			p2 = mid
		}

	}
	return mid
}

// 二分查找
func binarySearch(nums []int, i, j, t int) int {
	for i <= j {
		mid := (i + j) / 2
		if nums[mid] == t {
			return mid
		} else if nums[mid] < t {
			i = mid + 1
		} else {
			j = mid - 1
		}
	}
	return -1
}

/**
编写一个高效的算法来搜索 m x n 矩阵 matrix 中的一个目标值 target。该矩阵具有以下特性：

每行的元素从左到右升序排列。
每列的元素从上到下升序排列。
示例:

现有矩阵 matrix 如下：

[
  [1,   4,  7, 11, 15],
  [2,   5,  8, 12, 19],
  [3,   6,  9, 16, 22],
  [10, 13, 14, 17, 24],
  [18, 21, 23, 26, 30]
]
给定 target = 5，返回 true。

给定 target = 20，返回 false。
 */

// 规律 左下角(x,y) 为竖向最大值，横项最小值，
// 如果其比目标值小，说明竖向肯定不包含目标值，
// 如果其比目标值大说明横向这一排肯定不包含目标值
func SearchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}

	row := len(matrix) - 1
	col := 0

	for row >= 0 && col <len(matrix[0]) {
		v := matrix[row][col]
		if v == target {
			return true
		} else if v > target {
			row--
		} else {
			col++
		}
	}

	return false
}

func searchMatrix(matrix [][]int, target int) bool {
	column := len( matrix )
	if 0 == column { return false }
	length := len( matrix[0] )
	if 0 == length || matrix[0][0] > target || matrix[column-1][length-1] < target { return false }

	low := 0
	for i := 0; i < column; i++{
		low = i
		if target == matrix[i][0] {
			return true
		}else if target < matrix[i][0] {
			break
		}
	}

	high := 0
	for i := 0; i <= low; i++{
		high = i
		if target == matrix[i][length-1]{
			return true
		}else if target < matrix[i][length-1] {
			break
		}
	}

	i := high
	j := 0
	//ri := 0
	//rj := 0
	for ; i <= low; i++{
		//ri = i
		for j = 0; j < length; j++{
			//rj = j
			if target == matrix[i][j] {
				return true
			}
		}
	}

	return false
}