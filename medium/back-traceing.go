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


// leetcode God's algorithm
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

func subsets(nums []int) [][]int {
	res := make([][]int, 0)
	res = append(res, []int{})
	for i:=0; i<len(nums); i++ {
		subsetsHelper(nums[i], &res)
	}
	return res
}

func subsetsHelper(num int, res *[][]int)  {
	tmp := make([][]int, len(*res))
	copy(tmp, *res)
	for i:=0; i<len(tmp); i++ {
		t := make([]int, len(tmp[i]))
		copy(t, tmp[i])
		tmp[i] = append(t, num)

	}
	*res = append(*res, tmp...)
}

// leetcode God's algorithm
func Subsets(nums []int) [][]int {
	res := [][]int{[]int{}}
	slice := []int{}
	for i := 1; i <= len(nums); i++ {
		subsets2(nums, i, 0, &res, &slice)
	}
	return res
}

func subsets2(nums []int, depth, index int, res *[][]int, slice *[]int) {
	if len(*slice) == depth {
		t := make([]int, len(*slice))
		copy(t, *slice)
		*res = append(*res, t)
		return
	}

	for i := index; i < len(nums); i++ {
		t := *slice
		*slice = append(t, nums[i])
		subsets2(nums, depth, i+1, res, slice)
		if len(*slice) > 0 {
			t = *slice

			*slice = t[:len(t)-1]

		}
	}
}

func Exist(board [][]byte, word string) bool {
	visited := make(map[int]map[int]bool)
	for i:=0; i<len(board); i++ {
		visited[i] = map[int]bool{}
		for j,_ := range board[i] {
			visited[i][j] = false
		}
	}
	for i:=0; i<len(board); i++ {
		for j,_ := range board[i] {
			if board[i][j] == word[0] {
				if existHelper(board, word, i,j,0, visited) {
					return  true
				}
			}
		}
	}

	return false
}

func existHelper(board [][]byte, word string, x,y,index int, visited map[int]map[int]bool) bool  {
	if x >= len(board) || x < 0 || y >= len(board[x]) || y < 0{
		return false
	}

	if visited[x][y] == true || board[x][y] != word[index] {
		return false
	}

	if index == len(word) - 1 {
		visited[x][y] = true
		return true
	}

	visited[x][y] = true

	if 	existHelper(board, word, x + 1,y,index + 1, visited) || existHelper(board, word, x - 1,y,index + 1, visited) ||
		existHelper(board, word, x, y + 1,index + 1, visited) || existHelper(board, word, x, y - 1,index + 1, visited) {
		return true
	}

	visited[x][y] = false
	return  false
}

// leetcode's God algorithm

func exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == word[0] {
				if backtrackExist(board, i, j, word[1:]) {
					return true
				}
			}
		}
	}
	return false
}

func backtrackExist(board [][]byte, i, j int, w string) bool {
	if w == "" {
		return true
	}
	// set board[i][j] empty
	c := board[i][j]
	board[i][j] = '1'
	// check (i-1, j)
	if i > 0 && board[i-1][j] == w[0] {
		if backtrackExist(board, i-1, j, w[1:]) {
			return true
		}
	}
	// check (i+1, j)
	if i < len(board)-1 && board[i+1][j] == w[0] {
		if backtrackExist(board, i+1, j, w[1:]) {
			return true
		}
	}
	// check (i, j-1)
	if j > 0 && board[i][j-1] == w[0] {
		if backtrackExist(board, i, j-1, w[1:]) {
			return true
		}
	}
	if j < len(board[0])-1 && board[i][j+1] == w[0] {
		if backtrackExist(board, i, j+1, w[1:]) {
			return true
		}
	}
	board[i][j] = c
	return false
}
