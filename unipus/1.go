package main

import "fmt"

func func1() {
	str := ""
	fmt.Scanln(&str)
	left := make([]rune, 0)
	for _, v := range str {
		if v == '(' {
			left = append(left, v)
		} else {
			if len(left) == 0 {
				fmt.Println(0)
				return
			} else {
				if len(left) == 0 {
					fmt.Println(0)
					return
				}
				left = left[:len(left)-1]
			}
		}
	}
	if len(left) > 0 {
		fmt.Println(0)
		return
	}
	fmt.Println(1)
}
