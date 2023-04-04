package main

import (
	"fmt"
	"strconv"
)

func fun5() {
	var nums string
	fmt.Scanln(&nums)

	n, e := strconv.Atoi(nums)
	if e != nil {
		// fmt.Println(e.Error())
		fmt.Println(-1)
		return
	}
	if n%3 == 0 {
		fmt.Println(n)
		return
	}
	var s string
	for i := 1; i < len(nums); i++ {
		s = nums[i:]
		// fmt.Println("s = ", s)
		n, _ := strconv.Atoi(s)
		if n%3 == 0 {
			fmt.Println(n)
			return
		}
	}
	fmt.Println(-1)
}
