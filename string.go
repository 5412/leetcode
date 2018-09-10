package leetcode

func ReverseString(s string) string {
	l := len(s)
	res := make([]byte, 0)
	for i:=l-1;i>=0;i-- {
		res = append(res, s[i])
	}
	return string(res)
}

func ReverseInt(x int) int {
	//var res []int

	res := make([]int, 0)
	
	var isNegative bool

	if x < 0 {
		isNegative = true
		x = -x
	}
	
	for  {
		num := x % 10
		x = (x-num)/10

		res = append(res, num)
		if x == 0  {
			break
		}
	}

	var result int
	
	for i:=0; i<len(res); i++ {
		result = result * 10 + res[i]
	}

	if result > 1 << 31 - 1 {
		result = 0
	}

	if isNegative {
		result = -result
	}
	return result
}