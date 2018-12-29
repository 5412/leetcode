// Copyright (c) 2018. weixiong piao, All Right Reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package medium

import (
	"fmt"
	. "leetcode/structs"
)

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	var result *ListNode
	var head *ListNode

	for l1 != nil || l2 != nil {
		switch  {
		case l1 != nil && l2 != nil:
			sum := (l1.Val + l2.Val + carry) % 10
			carry = (l1.Val + l2.Val + carry) / 10
			tmpNode := ListNode{
				Val:sum,
				Next:nil,
			}
			if result == nil {
				result = &tmpNode
				head = result
			} else {
				result.Next = &tmpNode
				result = result.Next
			}
			l1 = l1.Next
			l2 = l2.Next
		case carry != 0 && l1 != nil:
			sum := (l1.Val + carry) % 10
			carry = (l1.Val + carry) / 10
			tmpNode := ListNode{
				Val:sum,
				Next:nil,
			}

			result.Next = &tmpNode
			result = result.Next
			l1 = l1.Next
		case carry != 0 && l2 != nil:
			sum := (l2.Val + carry) % 10
			carry = (l2.Val + carry) / 10
			tmpNode := ListNode{
				Val:sum,
				Next:nil,
			}
			carry = (l2.Val + carry) / 10
			result.Next = &tmpNode
			result = result.Next
			l2 = l2.Next
		case l2 != nil:
			result.Next = l2
			l2 = nil
		case l1 != nil:
			result.Next = l1
			l1 = nil
		default:
			fmt.Println("something wrong!")
		}
	}

	if carry != 0 {
		tmpNode := ListNode{
			Val:carry,
			Next:nil,
		}
		result.Next = &tmpNode
	}
	return head
}

// leetcode's God algorithm
// smart code

func AddTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	var head *ListNode
	var tail *ListNode
	var carry int
	for l1 != nil || l2 != nil {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		carry = sum / 10
		val := sum % 10

		node := &ListNode{val, nil}
		if head == nil {
			head = node
		}
		if tail == nil {
			tail = node
		} else {
			tail.Next = node
			tail = tail.Next
		}
	}

	if carry != 0 {
		node := &ListNode{carry, nil}
		if head == nil {
			 head = node
		} else {
			tail.Next = node
		}
	}
	return head
}

func OddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	walk := head
	var (
		odd *ListNode
		even *ListNode
	)
	index := 2
	odd = walk
	walk = walk.Next
	for walk != nil {
		node := walk
		walk = walk.Next
		node.Next = nil
		if (index & 1) == 1 {
			odd = insertListAtNode(odd, node)
		} else {
			even = insertListAtNode(even, node)
		}
		index++
	}
	return head
}

func insertListAtNode(n1 *ListNode, n2 *ListNode) * ListNode {
	if n1 == nil {
		n1 = n2
	} else {
		next := n1.Next
		n1.Next = n2
		n2.Next = next
	}
	return n2
}


// leetcode's God algorithm 

func OddEvenListLeet(head *ListNode) *ListNode  {
	if head == nil{
		return nil
	}
	odd,even,evenh := head,head.Next,head.Next
	for even != nil && even.Next != nil{
		odd.Next,even.Next = even.Next,even.Next.Next
		odd = odd.Next
		even = even.Next
	}
	odd.Next = evenh
	return head
}

// 或者把两个链表补到相等长度
func GetIntersectionNode(headA, headB *ListNode) *ListNode {
	scanA, scanB := headA, headB
	if scanA == nil || scanB == nil {
		return nil
	}
	for scanA != scanB  {
		if scanA == nil  {
			scanA = headB
		} else {
			scanA = scanA.Next
		}

		if scanB == nil  {
			scanB = headA
		} else {
			scanB = scanB.Next
		}
	}

	return scanA
}