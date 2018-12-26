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

func SetZeros(matrix [][]int) {

	rows := len(matrix)
	if rows == 0 {
		return
	}

	colls := len(matrix[0])
	if colls == 0 {
		return
	}

	firstRow := false
	firstColl := false

	// 第一步如果出现0 则将横轴第一位设为0，纵轴第一位设为0，如此设计可定不会改变循环中[i,j]之后的元素
	for i:=0; i<rows; i++ {
		for j:=0; j<colls; j++ {
			if matrix[i][j] == 0 {
				if i == 0 {
					firstRow = true
				}
				if j == 0 {
					firstColl = true
				}
				matrix[0][j] = 0
				matrix[i][0] = 0
			}
		}
	}

	// 分别从横向（纵向）第二行（列）循环，如果首元素为1则全行（列）都设为0
	// 从第二列循环的原因就是不污染第一行第一列的值
	for i:=1; i<rows; i++ {
		if matrix[i][0] == 0 {
			for j:=1; j<colls; j++ {
				matrix[i][j] = 0
			}
		}
	}

	for j:=1; j<colls; j++ {
		if matrix[0][j] == 0 {
			for i:=1; i<rows; i++ {
				matrix[i][j] = 0
			}
		}
	}

	//此时处理第一行第一列的情况
	if firstRow {
		for i:=0; i<colls; i++ {
			matrix[0][i] = 0
		}
	}

	if firstColl {
		for i:=0; i<rows; i++ {
			matrix[i][0] = 0
		}
	}

	// 非原地算法实现，原地算法的意思就是固定空间
	//var (
	//	x []int
	//	y []int
	//)
	//
	//
	//m := len(matrix)
	//n := len(matrix[0])
	//
	//if m*n == 0 {
	//	return
	//}
	//
	//
	//for i := 0; i<m; i++ {
	//	for j:=0; j<n; j++ {
	//		if matrix[i][j] == 0 {
	//			x = append(x, i)
	//			y = append(y, j)
	//		}
	//	}
	//}
	//fmt.Println(x, y)
	//
	//if len(x) > 0 {
	//	for i:=0; i<len(x); i++ {
	//		for j:=0; j<n; j++ {
	//			matrix[x[i]][j] = 0
	//		}
	//	}
	//}
	//
	//if len(y) > 0 {
	//	for i:=0; i<len(y); i++ {
	//		for j:=0; j<m; j++ {
	//			matrix[j][y[i]] = 0
	//		}
	//	}
	//}
}

func GroupAnagrams(str []string) [][]string {
	l := len(str)

	if l == 0 {
		return nil
	}

	if l == 1 {
		return [][]string{str}
	}

	m := make(map[string][]string, 0)

	for i:=0; i<l; i++ {
		bs := []byte(str[i])
		sort.Slice(bs, func(i, j int) bool {
			return bs[i] < bs[j]
		})
		keyString := string(bs)
		m[keyString] = append(m[keyString], str[i])
	}
	result := make([][]string, 0)
	for _, v := range m {
		result = append(result, v)
	}

	return result
}

// leetcode上大神的字母异位词分组
// 考虑 所有输入均为小写字母，故可使用countSort返回排序好的字符串
// 不考虑输出顺序， m := make(map[string]int) 此map的key存的是单词排序好的字符串，value存的值为，应该在返回结果中的位置
// 大神的能省CPU也能省空间
func groupAnagrams(strs []string) [][]string {
	m := make(map[string]int)
	r := make([][]string, 0)
	i := 0
	for _, s := range strs {
		ss := countSort(s)
		if id, ok := m[ss]; ok {
			r[id] = append(r[id], s)
		} else {
			m[ss] = i
			r = append(r, []string{s})
			i++
		}
	}
	return r
}

func countSort(s string) string {
	var cnt [26]int
	for i := 0; i < len(s); i++ {
		cnt[s[i]-byte('a')]++
	}
	ret := make([]byte, len(s))
	j := 0
	for i := 0; i < 26; i++ {
		for k := 0; k < cnt[i]; k++ {
			ret[j] = byte(i) + byte('a')
			j++
		}
	}
	return string(ret)
}

func LengthOfLongestSubstring(s string) int {
	m := make(map[rune]bool,0)
	runes := make([]rune,0)
	l := 0
	for _,v := range s {
		if m[v] {
			if l< len(runes) {
				l = len(runes)
			}
			i := 0
			for i<len(runes) {
				if runes[i] == v {
					m[v] = false
					i++
					break
				}
				m[runes[i]] = false
				i++

			}
			runes = runes[i:]
		}
		m[v] = true
		runes = append(runes, v)
	}
	if l< len(runes) {
		l = len(runes)
	}
	return l
}


// leetcode's God algorithm

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	mp := make(map[byte]int)
	ans := 0
	for i, j := 0, 0; j < n; j++ {
		if mp[s[j]] != 0 {
			if mp[s[j]] > i {
				// mp[a] 存储的为a元素下一个元素的位置
				// i为无重复字符的起始点
				// 如a存在 且起始点i < 之前a元素的下一个元素点则替换i，以保证从i到j无重复记录
				// 1，2 = 2-1+1个元素
				i = mp[s[j]]
			}
		}
		if t := j - i + 1; t > ans {
			ans = t
		}
		mp[s[j]] = j + 1
	}
	return ans
}