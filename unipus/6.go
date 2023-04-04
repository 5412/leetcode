package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	intpus := strings.Split(input, " ")
	nums := make([]rune, len(intpus))
	for i, v := range intpus {
		nums[i] = rune(v[0])
	}
	scanner.Scan()
	input2 := scanner.Text()
	//n, _ := strconv.Atoi(input2)
	l2 := len(input2)
	l1 := len(nums)
	res := 0
	for i := 1; i < l2; i++ {
		res += int(math.Pow(float64(l1), float64(i)))
		fmt.Println(res)
	}

	for _, v := range input2 {
		nl := littleNum(nums, v)
		if nl == 0 {
			break
		} else {
			res += nl * int(math.Pow(float64(l1), float64(l2-1)))
			l2--
		}
		if equelNum(nums []rune, r rune)
	}
	fmt.Println(res)
}

func littleNum(nums []rune, r rune) int {
	i := 0
	for _, v := range nums {
		if v < r {
			return i
		}
		i++
	}
	return i
}

func equelNum(nums []rune, r rune) bool {
	for _, v := range nums {
		if v == r {
			return true
		}
	}
	return false
}
