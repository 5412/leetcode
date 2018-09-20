package leetcode

func Merge(nums1 []int, m int, nums2 []int, n int)  {

	// leetcode God's algorithm
	// note: nums1 尾部多余空间，做倒序插入

	for {
		if m == 0 && n == 0 {
			break
		}
		if m > 0 && n >0 {
			if nums1[m-1] > nums2[n-1] {
				nums1[m+n-1] = nums1[m-1]
				m--
			} else {
				nums1[m+n-1] = nums2[n-1]
				n--
			}
		} else if m > 0 {
			nums1[m+n-1] = nums1[m-1]
			m--
		} else {
			nums1[m+n-1] = nums2[n-1]
			n--
		}
	}


	//i := 0
	//j := 0
	//t := 0
	//buf := make([]int, m+n)
	//for i<n && j<m {
	//	if nums2[i] < nums1[j] {
	//		buf[t] = nums2[i]
	//		i++
	//	} else {
	//		buf[t] = nums1[j]
	//		j++
	//	}
	//	t++
	//}
	//
	//for i < n {
	//	buf[t] = nums2[i]
	//	i++
	//	t++
	//}
	//
	//for j < m {
	//	buf[t] = nums1[j]
	//	j++
	//	t++
	//}
	//
	//for i,v := range(buf) {
	//	nums1[i] = v
	//}
}

// leetcode: /explore/interview/card/top-interview-questions-easy/8/sorting-and-searching/53/
// 第一个错误版本，使用C语言实现
// 基本思路，二分查找。
// note: 注意中间位置计算
// Forward declaration of isBadVersion API.


//bool isBadVersion(int version);
//
//int firstBadVersion(int n) {
//	int first = 1;
//	int end = n;
//	int mid = 0;
//	while (first <= end) {
//		mid = first + (end-first) / 2;
//		if (isBadVersion(mid)) {
//			end = mid-1;
//		} else {
//			first = mid+1;
//		}
//	}
//	return first;
//}
