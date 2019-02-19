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


