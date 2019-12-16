package leetcode

import "leetcode/easy"

func RotateRight(head *easy.ListNode, k int) *easy.ListNode {

	// leetcode God's algorithm
	// 使用快慢双指针找到倒数第k个节点，
	// 将其是为链表头节点,
	// 左后一个几点的next设为愿head节点，
	// 将其前一个节点断开

	if head == nil || head.Next == nil || k <= 0 {
		return head
	}

	first, last, l := head, head, 0

	// l 等于 k与链表长度取模操作的余数
	for k > 0 {
		l, k = l+1, k-1
		if last.Next == nil {
			k = k % l
			l = 0
			last = first
		} else {
			last = last.Next
		}
	}

	// 此时last为链表第k位之后那个节点
	// k刚好是l的整数倍时
	if last == head {
		return head
	}

	// 快慢指针，当last移动到末尾的时候会移动 链表长度-k 次，first则恰好移动到倒数第k位
	// 此时只需连接原链表首尾，断开倒数第K位与前节点的链接，并设置链表头节点位倒数第K位节点即可

	for last.Next != nil {
		last = last.Next
		first = first.Next
	}

	last.Next = head
	head = first.Next
	first.Next = nil

	return head

	//if k == 0 {
	//	return head
	//}
	//
	//if head == nil{
	//	return head
	//}

	// 第一次改进方案

	//len  := 1
	//temp := head
	//
	//for temp.Next != nil  {
	//	temp = temp.Next
	//	len++
	//}
	//
	//k = k%len

	// 超出了时间限制，故而进一步改进
	//var first *easy.ListNode
	//var front *easy.ListNode
	//
	//
	//for k > 0 {
	//	first = head
	//	for head.Next != nil  {
	//		front = head
	//		head = head.Next
	//	}
	//
	//	front.Next = nil
	//	head.Next = first
	//	k--
	//}
}
