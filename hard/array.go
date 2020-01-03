package hard

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
