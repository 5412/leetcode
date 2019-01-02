// Copyright (c) 2018. weixiong piao, All Right Reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package structs

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func PrintList(l *ListNode)  {
	for l.Next != nil {
		fmt.Print(l.Val, "-")
		l = l.Next
	}
	fmt.Println(l.Val)
}

func GenerateList(vals[]int) *ListNode {
	var result *ListNode
	var headList *ListNode
	for _,v := range vals {
		tempNode := &ListNode{v,nil}
		if result == nil {
			result =tempNode
			headList = result
		} else {
			result.Next = tempNode
			result = result.Next
		}
	}
	return headList
}