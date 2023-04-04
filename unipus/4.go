package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func fun4() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	intpus := strings.Split(input, " ")
	nums := make([]int, len(intpus))
	for i, v := range intpus {
		nums[i], _ = strconv.Atoi(v)
	}
	res := make([]int, len(nums))
	if len(nums) == 1 {
		fmt.Println(nums[0])
		return
	}
	if len(nums) <= 2 {
		fmt.Println(max(nums[0], nums[1]))
		return
	}
	res[0] = nums[0]
	res[1] = max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		res[i] = max(res[i-2]+nums[i], res[i-1])
	}
	fmt.Println(res[len(nums)-1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
