package leetcode

import "sort"

//leetcode 官方答案解析
// func permutation(s string) (ans []string) {
// 	t := []byte(s)
// 	sort.Slice(t, func(i, j int) bool { return t[i] < t[j] })
// 	n := len(t)
// 	perm := make([]byte, 0, n)
// 	vis := make([]bool, n)
// 	var backtrack func(int)
// 	backtrack = func(i int) {
// 		if i == n {
// 			ans = append(ans, string(perm))
// 			return
// 		}
// 		for j, b := range vis {
// 			if b || j > 0 && !vis[j-1] && t[j-1] == t[j] {
// 				continue
// 			}
// 			vis[j] = true
// 			perm = append(perm, t[j])
// 			backtrack(i + 1)
// 			perm = perm[:len(perm)-1]
// 			vis[j] = false
// 		}
// 	}
// 	backtrack(0)
// 	return
// }

// func permutation(s string) (ans []string) {
// 	base := []byte(s)
// 	n := len(base)
// 	res := make([]string, 0, n)
// 	resMap := make(map[string]bool)
// 	var iterator func(i int)
// 	iterator = func(i int) {
// 		if i == n {
// 			ss := string(base)
// 			if resMap[ss] {
// 				return
// 			}
// 			res = append(res, ss)
// 			resMap[ss] = true
// 			return
// 		}
// 		for x := i; x < n; x++ {
// 			base[i], base[x] = base[x], base[i]
// 			iterator(i + 1)
// 			base[i], base[x] = base[x], base[i]
// 		}
// 	}
// 	iterator(0)
// 	return res
// }

//思路：使用寻找下一个增长序列
//步骤：1.升序排序 2.依次使用寻找下一个序列查找下一个符合条件的序列 3.返回结果
//末尾开始寻找第一个非逆序元素 p1, 右数大于左数
//从后向前寻找第一个大于此元素的元素 p2, 互相交换值
//反向p1位置之后元素

func permutation(s string) []string {
	bs := []rune(s)
	l := len(bs)
	if l == 0 || l > 8 {
		return []string{}
	}

	sort.Slice(bs, func(a, b int) bool {
		return bs[a] < bs[b]
	})

	list := make([]string, 0, l)
	for {
		list = append(list, string(bs))
		if !nextPermutation(bs) {
			break
		}
	}
	return list
}

func nextPermutation(bs []rune) bool {
	l := len(bs)
	//左边起始查找位，按照之前置换原则，只有右数比左数大才能有下一个更大序列
	lIdx := l - 2
	for lIdx >= 0 && bs[lIdx] >= bs[lIdx+1] {
		lIdx--
	}
	if lIdx < 0 {
		return false
	}
	//从最右边开始找，第一个比lIdx大的数即为，最小的大值
	rIdx := l - 1
	for rIdx >= 0 && bs[lIdx] >= bs[rIdx] {
		rIdx--
	}
	bs[lIdx], bs[rIdx] = bs[rIdx], bs[lIdx]
	//按照前面的置换原则，将后面的序列，反转即为新的升序[)
	reverse(bs[lIdx+1:])
	return true
}

func reverse(list []rune) {
	for i, l := 0, len(list); i < l/2; i++ {
		list[i], list[l-1-i] = list[l-1-i], list[i]
	}
}
