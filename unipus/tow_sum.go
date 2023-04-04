package main

import (
	"fmt"
)

func towSum() {
	var arrNum int
	var params []int
	var targetNum int
	var targets []int
	fmt.Scanln(&arrNum)
	params = make([]int, 0, arrNum)
	for i := 0; i < arrNum; i++ {
		var num int
		fmt.Scanf("%d", &num)
		params = append(params, num)
	}
	fmt.Scanln(&targetNum)
	targets = make([]int, 0, targetNum)
	for i := 0; i < targetNum; i++ {
		var num int
		fmt.Scanf("%d", &num)
		targets = append(targets, num)
	}
	paramsMap := make(map[int]bool)
	res := make([]string, 0)

	for _, v := range params {
		paramsMap[v] = true
	}
	for _, v := range targets {
		resovedMap := make(map[int]bool)
		for _, j := range params {
			if resovedMap[v-j] {
				continue
			}
			if paramsMap[v-j] && !resovedMap[v-j] {
				res = append(res, fmt.Sprintf("%d %d", j, v-j))
				resovedMap[v-j] = true
				resovedMap[j] = true
			}
		}
	}

	for _, v := range res {
		fmt.Println(v)
	}
}
