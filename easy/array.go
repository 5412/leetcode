package easy

func CountRangeSum(nums []int, lower int, upper int) int {
	var snum int = 0
	l := len(nums)
	for i := 0; i < l; i++ {
		sum := 0
		for j := i; j < l; j++ {
			sum += nums[j]
			if sum >= lower && sum <= upper {
				snum++
			}
		}
	}
	return snum
}

func RemoveDuplicates(nums []int) int {
	l := len(nums)
	n := 0

	for j := 0; j < l; j++ {
		if nums[n] != nums[j] {
			n++
			nums[n] = nums[j]
		}
	}

	return n + 1
}

func MaxProfit(prices []int) int {
	l := len(prices)
	profit := 0
	for i := 1; i < l; i++ {
		if prices[i] > prices[i-1] {
			profit += prices[i] - prices[i-1]
		}
	}
	return profit
}

func Rotate(nums []int, k int) {

	if k < 0 {
		return
	}

	l := len(nums)
	k = k % l

	// space O(1)
	var temp int
	for i := 0; i < k; i++ {
		temp = nums[l-1]
		for j := l - 1; j > 0; j-- {
			nums[j] = nums[j-1]
		}
		nums[0] = temp
	}

	// copy leetcode's answer

	//for i := 0; i < len(nums)/2; i++ {
	//	nums[i], nums[len(nums)-i-1] = nums[len(nums)-i-1], nums[i]
	//}
	//for i := 0; i < k/2; i++ {
	//	nums[i], nums[k-i-1] = nums[k-i-1], nums[i]
	//}
	//for i := k; i < (len(nums)+k)/2; i++ {
	//	nums[i], nums[len(nums)+k-i-1] = nums[len(nums)+k-i-1], nums[i]
	//}

	// time a little quick may cost a double space (may be under line's array doest copy)
	//last := nums[l-k:]
	//frist := nums[:l-k]
	//temp := append(last, frist...)
	//for i:=0; i<l; i++ {
	//	nums[i] = temp[i]
	//}

}

func ContainsDuplicate(nums []int) bool {
	container := make(map[int]bool, len(nums))
	for _, v := range nums {
		e := container[v]
		if e {
			return true
		}
		container[v] = true
	}
	return false
}

// leetcode subject: explore/interview/card/top-interview-questions-easy/1/array/25/
// require:
// time-complexity: O(n)
// space-complexity: O(1) , no exceptional space
func SingleNumber(nums []int) int {

	// leetcode's some god's solution
	// keynote:
	// the input only one element occurs one, each others occurs twice
	// 1 ^ 1 ^ 3 = 3
	var res int

	for _, v := range nums {
		res ^= v
	}

	return res

	// second time
	//var result,isDuplicate int
	//l := len(nums)
	//
	//for i:=0; i<l; i++ {
	//	isDuplicate = 0
	//	for j := 0; j < l; j++ {
	//		if nums[j] == nums[i] {
	//			isDuplicate += 1
	//		}
	//	}
	//
	//	if isDuplicate == 1 {
	//		result = nums[i]
	//		break
	//	}
	//}

	// first time
	//var result int
	//m := make(map[int]int, len(nums))
	//
	//for _,v := range nums {
	//	m[v] += 1
	//}
	//
	//for i,k := range m {
	//	if k == 1 {
	//		return i
	//	}
	//}
	//return result
}

// leetcode: explore/interview/card/top-interview-questions-easy/1/array/26/

func Intersect(nums1 []int, nums2 []int) []int {
	//output := make([]int, 0)
	var output []int
	//var m map[int]int
	m := make(map[int]int, 0)
	for _, v1 := range nums1 {
		m[v1] += 1
	}

	for _, v := range nums2 {
		if m[v] > 0 {
			output = append(output, v)
			m[v]--
		}
	}
	return output
}

func PlusOne(digits []int) []int {
	//var res []int
	l := len(digits)
	for i := l; i > 0; i-- {
		tail := digits[i-1]
		tail++
		switch tail {
		case 10:
			tail = 0
		}
		digits[i-1] = tail

		if tail > 0 {
			break
		}
		// last loop
		if i == 1 {
			digits = append([]int{1}, digits...)
		}
	}

	return digits

	// Don't work while the digits is to large
	//number := 0
	//for _,v := range digits  {
	//	number = number *  10 + v
	//}
	//number++
	//output := make([]int,0)
	//for  {
	//	if number == 0 {
	//		break
	//	}
	//	digit := number % 10
	//	output = append([]int{digit}, output...)
	//	number = (number - digit) / 10
	//}
	//return output
}

// leetcode: explore/interview/card/top-interview-questions-easy/1/array/28/
func MoveZeros(nums []int) {
	l := len(nums)
	for i := l - 1; i >= 0; i-- {
		if nums[i] == 0 {
			for j := i; j < l-1; j++ {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				//nums[j] = nums[j+1]
			}
			//nums[l-1] = 0
			l--
		}
	}
}

// leetcode: explore/interview/card/top-interview-questions-easy/1/array/29/
func TwoSum(nums []int, target int) []int {
	// God's algorithm
	res := make([]int, 2)
	m := make(map[int]int, len(nums))

	for k, v := range nums {
		s := target - v
		if i, ok := m[s]; ok {
			res[0] = i
			res[1] = k
		} else {
			m[v] = k
		}
	}
	return res

	//l := len(nums)
	//for i:=0; i<l ; i++ {
	//	for j:= i+1; j<l; j++ {
	//		if nums[i] + nums[j] == target {
	//			return []int{i,j}
	//		}
	//	}
	//}
	//return nil
}

func IsValidSudoku(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {

			if board[i][j] == '.' {
				continue
			}

			//  horizontal
			for z := 0; z < 9; z++ {
				if board[i][z] != '.' && z != j && board[i][z] == board[i][j] {
					return false
				}
			}

			// vertical
			for z := 0; z < 9; z++ {
				if board[z][j] != '.' && z != i && board[z][j] == board[i][j] {

					return false
				}
			}
			// surface
			x := i / 3
			y := j / 3
			beginX := x * 3
			beginY := y * 3
			for z := 0; z < 3; z++ {
				for h := 0; h < 3; h++ {

					if board[beginX+z][beginY+h] == '.' {
						continue
					}

					if  beginX+z != i && beginY+h != j && board[beginX+z][beginY+h] == board[i][j] {
						return false
					}
				}
			}
		}

	}
	return true
}


// leetcode: /explore/interview/card/top-interview-questions-easy/1/array/31/
func MatrixRotate(matrix [][]int) {
	n := len(matrix)
	l := len(matrix)

	// n*n matrix
	for i:=0; i<n; i++ {
		for j:=i; j<n-1; j++ {
			x := i
			y := j
			for z:=0; z<3; z++  {
				x2 := l - 1 - y
				y2 := x
				matrix[x][y], matrix[x2][y2] = matrix[x2][y2],matrix[x][y]
				x = x2
				y = y2
			}
		}
		n--
	}
}

