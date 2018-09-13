package leetcode

type ListNode struct {
	Val int
	Next *ListNode
}


// leetcode: explore/interview/card/top-interview-questions-easy/6/linked-list/43/
func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	// God's algorithm
	// 快慢双指针，快指针比慢指针先走n步，然后快慢双指针同时走，当快指针达到尾部时，慢指针达到第l-n节
	// 点，倒数第n节点，即l-n节点的下一节点
	// 需要处理删除头节点和尾节点的情况

	fast := head
	slow := head
	l := 0
	for fast.Next != nil  {
		if n <= 0 {
			slow = slow.Next
		}
		n--
		l++
		fast = fast.Next
	}
	if n == 1 { // 删除头节点
		return head.Next
	} else if n == -l { // 删除尾节点
		//slow.Next = nil // 经测试此处是个伪命题
	} else {
		slow.Next = slow.Next.Next
	}
	return head

	//current := head
	//pre := head
	//for {
	//	tmp := current
	//	for i:=0;i<n;i++ {
	//		tmp = tmp.Next
	//	}
	//
	//	if tmp == nil {
	//		break
	//	} else {
	//		pre = current
	//		current = current.Next
	//	}
	//}
	//if (current.Next == nil) {
	//	pre.Next = nil
	//	if pre == head {
	//		return pre.Next
	//	}
	//} else {
	//	current.Val = current.Next.Val
	//	current.Next = current.Next.Next
	//}
	////first.Val = first.Next.Val
	////first.Next = first.Next.Next
	//
	//return head
}

// leetcode: explore/interview/card/top-interview-questions-easy/6/linked-list/43/

func ReverseList(head *ListNode) *ListNode {

	// leetcode God's algorithm

	if head == nil {
		return nil
	}

	cur := head

	for cur.Next != nil {
		next := cur.Next // next 为当前节点的下一个节点
		cur.Next = next.Next // 设置当前节点的下一个节点为 next 节点的下一个节点， 这样next节点就飞起来了
		next.Next = head // 将next 节点的写一个节点设为head节点，即将next节点放到head前面
		head = next // 重置 head
	}

	return head

	//var h *ListNode
	//var tmp *ListNode
	//
	//if head == nil {
	//	return nil
	//}
	//
	//if head.Next == nil {
	//	return head
	//}
	//
	//for head.Next != nil {
	//	tmp = head
	//	head = head.Next
	//	tmp.Next = h
	//	h = tmp
	//}cd
	//tmp = head
	//tmp.Next = h
	//h = tmp
	//return h
}




