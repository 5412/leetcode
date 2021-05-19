package leetcode

/*
 * @lc app=leetcode.cn id=1442 lang=golang
 *
 * [1442] 形成两个异或相等数组的三元组数目
 *
 * https://leetcode-cn.com/problems/count-triplets-that-can-form-two-arrays-of-equal-xor/description/
 *
 * algorithms
 * Medium (67.22%)
 * Likes:    127
 * Dislikes: 0
 * Total Accepted:    23.6K
 * Total Submissions: 30K
 * Testcase Example:  '[2,3,1,6,7]'
 *
 * 给你一个整数数组 arr 。
 *
 * 现需要从数组中取三个下标 i、j 和 k ，其中 (0 <= i < j <= k < arr.length) 。
 *
 * a 和 b 定义如下：
 *
 *
 * a = arr[i] ^ arr[i + 1] ^ ... ^ arr[j - 1]
 * b = arr[j] ^ arr[j + 1] ^ ... ^ arr[k]
 *
 *
 * 注意：^ 表示 按位异或 操作。
 *
 * 请返回能够令 a == b 成立的三元组 (i, j , k) 的数目。
 *
 *
 *
 * 示例 1：
 *
 * 输入：arr = [2,3,1,6,7]
 * 输出：4
 * 解释：满足题意的三元组分别是 (0,1,2), (0,2,2), (2,3,4) 以及 (2,4,4)
 *
 *
 * 示例 2：
 *
 * 输入：arr = [1,1,1,1,1]
 * 输出：10
 *
 *
 * 示例 3：
 *
 * 输入：arr = [2,3]
 * 输出：0
 *
 *
 * 示例 4：
 *
 * 输入：arr = [1,3,5,7,9]
 * 输出：3
 *
 *
 * 示例 5：
 *
 * 输入：arr = [7,11,12,9,5,2,7,17,22]
 * 输出：8
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= arr.length <= 300
 * 1 <= arr[i] <= 10^8
 *
 * sum(j,k) = sum[j] ^ sum[k]
 * sum[i] ^ sum[j] = sum[j] ^ sum[k] => sum[i] = sum[k]
 * 与 j 无关, 个数为 k-i
 * 对于下标 k 所有满足的 i 对结果的贡献为
 * (k-ix) * m = mk - (i0 + i1 + ... +ix)
 * m =  i出现次数
 *
 */

// @lc code=start
func countTriplets(arr []int) (ans int) {
	cnt := map[int]int{}
	total := map[int]int{}
	s := 0
	for k, val := range arr {
		if m, has := cnt[s^val]; has {
			ans += m*k - total[s^val]
		}
		cnt[s]++
		total[s] += k
		s ^= val
	}
	return

}

// @lc code=end
