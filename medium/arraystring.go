package medium

import (
	"sort"
)

func ThreeSum(nums []int) [][]int {

	// algorithm from leetcode's God
	sort.Ints(nums)
	l := len(nums)
	var result [][]int
	for i := 0; i < l - 1 && nums[i] <= 0; i++ {

		// 在上一个循环已经处理过相等元素问题了，故此处不再处理直接跳过
		if i > 0 && nums[i] == nums[i - 1] {
			continue
		}
		j := i + 1
		k := l - 1
		target := 0 - nums[i]
		for ; j < k; {
			sum := nums[j] + nums[k]
			if sum == target {
				result = append(result, []int{nums[i], nums[j], nums[k]})
				j++
				for j < k && nums[j-1] == nums[j] {
					j++
				}
				k--
				for j < k && nums[k+1] == nums[k] {
					k--
				}
			} else if sum > target {
				k--
				for j < k && nums[k+1] == nums[k] {
					k--
				}
			} else {
				j++
				for j < k && nums[j-1] == nums[j] {
					j++
				}
			}
		}
	}
	return result

	// Some algorithm from csdn
	//var res [][]int
	//sort.Ints(nums)
	//numsLen := len(nums)
	//if numsLen < 3 {
	//	return res
	//}
	//
	//l, r, dif := 0, 0, 0
	//
	//for i:=0; i< numsLen; i++  {
	//	if nums[i] > 0 {
	//		break
	//	}
	//
	//	if i > 0 && nums[i-1] == nums[i] {
	//		continue
	//	}
	//
	//	l = i + 1
	//	r = numsLen - 1
	//	dif = -nums[i]
	//
	//	for l < r {
	//		if nums[l] + nums[r] == dif {
	//			res = append(res, []int{nums[i], nums[l], nums[r]})
	//			for l<r && nums[l] == nums[l+1] {
	//				l += 1
	//			}
	//
	//			for l<r && nums[r] == nums[r-1]  {
	//				r -= 1
	//			}
	//
	//			l++
	//			r--
	//		} else if nums[l] + nums[r] < dif {
	//			l++
	//		} else {
	//			r--
	//		}
	//	}
	//
	//
	//}
	//return res


	// 暴躁实现，本地运行能通过，leetcode上当数据大量时编译错误，目前原因未知
	// 最后输入：[10,-2,-12,3,-15,-12,2,-11,3,-12,9,12,0,-5,-4,-2,-7,-15,7,4,-5,-14,-15,-15,-4,10,9,
	// -6,7,1,12,-6,14,-15,12,14,10,0,10,-10,3,4,-12,10,7,-9,-7,-15,-8,-15,-4,2,9,-4,3,-10,13,-3,-1,
	// 5,5,-4,-15,4,-11,4,-4,6,-11,-9,12,7,11,7,2,-5,13,10,-5,-10,3,7,0,-3,10,-12,2,9,-8,8,-9,13,12,
	// 13,-6,8,3,5,-9,7,12,10,-8,0,2,1,10,-7,-3,-10,-10,7,4,5,-11,-8,0,-2,-14,8,13,-8,-2,10,13]

	//var res [][]int
	//l := len(nums)
	//if l < 3 {
	//	return res
	//}
	//m := make(map[int]map[int]bool, l)
	//for i:=0; i<l-2; i++ {
	//	for j:=i+1;j<l-1;j++{
	//		for z:= j+1;z<l;z++ {
	//			if nums[i] + nums[j] + nums[z] == 0 {
	//				if m[nums[i]][nums[j]] || m[nums[i]][nums[z]] {
	//					continue
	//				}
	//				m[nums[i]] = map[int]bool{nums[j]:true,nums[z]:true}
	//				m[nums[j]] = map[int]bool{nums[i]:true,nums[z]:true}
	//				m[nums[z]] = map[int]bool{nums[j]:true,nums[i]:true}
	//				res = append(res, []int{nums[i],nums[j],nums[z]})
	//			}
	//		}
	//	}
	//}
	//return res
}
