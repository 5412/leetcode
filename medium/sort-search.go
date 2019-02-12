// Copyright (c) 2019. weixiong piao, All Right Reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package medium

import "fmt"

func SortColor(nums []int) {
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
