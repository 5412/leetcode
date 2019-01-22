// Copyright (c) 2019. weixiong piao, All Right Reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package medium

func LetterCombinations(digits string) []string {
	var res []string
	//res := make([]string, 0)

	m := map[rune][]string{
		'2': []string{"a", "b", "c"},
		'3': []string{"d", "e", "f"},
		'4': []string{"g", "h", "i"},
		'5': []string{"j", "k", "l"},
		'6': []string{"m", "n", "o"},
		'7': []string{"p", "q", "r", "s"},
		'8': []string{"t", "u", "v"},
		'9': []string{"w", "x", "y", "z"},
	}

	for _, md := range digits {
		tm := m[md]
		if res == nil {
			for _, v := range tm {
				res = append(res, v)

			}
		} else {
			//var tmp []string
			//tmp := make([]string, len(res))
			//copy(tmp, res)
			//fmt.Println(tmp)
			//res = make([]string,0)
			l := len(res)
			for _, r := range res {
				for _, v := range tm {
					res = append(res, r+v)
				}
			}
			res = res[l:]
		}
	}

	return res
}

// some leetcoder's algorithm
// 使用递归
func letterCombinations(digits string) []string {

	if len(digits) < 1 {
		return []string{}
	}

	maps := make(map[int]string)
	maps[2] = "abc"
	maps[3] = "def"
	maps[4] = "ghi"
	maps[5] = "jkl"
	maps[6] = "mno"
	maps[7] = "pqrs"
	maps[8] = "tuv"
	maps[9] = "wxyz"

	res := fun(digits, maps)

	return res
}

func fun(digits string, maps map[int]string) (result []string) {
	if len(digits) < 1 {
		return result
	}

	if len(digits) == 1 {
		ch := digits[0]
		n := int(ch - '0')

		str := maps[n]
		for i := 0; i < len(str); i++ {
			result = append(result, str[i:i+1])
		}
		return result
	}

	data := fun(digits[1:], maps)
	ch := digits[0]
	n := int(ch - '0')

	str := maps[n]
	for i := 0; i < len(str); i++ {
		for j := 0; j < len(data); j++ {
			result = append(result, str[i:i+1]+data[j])
		}
	}
	return result
}

// some algorithm from leetcode
func GenerateParenthesis(n int) []string {
	if n == 1 {
		return []string{"()"}
	}
	res := make([]string, 0)
	generate(&res, "", 0, 0,n)
	return res
}

func generate(res *[]string,s string, leftCount, rightCount, n int) {
	if leftCount > n || rightCount > n {
		return
	}
	if leftCount == n && rightCount == n {
		*res = append(*res, s)
		return
	}

	if leftCount >= rightCount {
		generate(res, s + "(", leftCount+1, rightCount, n)
		generate(res, s + ")", leftCount, rightCount+1, n)
	}

}

func generateParenthesis(n int) []string {
	res := make([]string, 0)
	gen(&res, "", n,n)
	return res
}


func gen(res *[]string, s string, left, right int) {
	if left == 0 && right == 0 {
		*res = append(*res, s)
	}

	if left > 0 {
		gen(res, s + "(", left - 1, right)
	}

	if right > 0 && left < right {
		gen(res, s + ")", left, right - 1)
	}
}

func Permute(nums []int) [][]int {
	res := make([][]int, 0)
	permuteHelper(nums, 0, &res)
	return res
}

func permuteHelper(nums []int, start int, res *[][]int) {
	if start >= len(nums) {
		temp := make([]int, len(nums))
		copy(temp, nums)
		*res = append(*res, temp)
	}

	for i:=start; i< len(nums); i++ {
		nums[start], nums[i] = nums[i], nums[start]

		permuteHelper(nums, start+1, res)
		nums[start], nums[i] = nums[i], nums[start]
	}
}

func permute(nums []int) [][]int {
	var res [][]int
	for i:=len(nums)-1; i>=0; i--{
		res = oneMore(nums[i], res)
	}
	return res
}

func oneMore(more int, exist [][]int) [][]int {
	if len(exist) == 0 {
		return [][]int{[]int{more}}
	}

	res := [][]int{}
	for _, arr := range exist {
		for i := 0; i <= len(arr); i++ {
			res = append(res, insertArrayPos(i, more, arr))
		}
	}

	return res
}

func insertArrayPos(pos, value int, arr []int) []int {
	newArr := make([]int, 0, len(arr)+1)
	newArr = append(newArr, arr[:pos]...)
	newArr = append(newArr, value)
	newArr = append(newArr, arr[pos:]...)

	return newArr
}
