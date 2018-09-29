package easy

import (
	"fmt"
)

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

func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	// leetcode: God's algorithm

	head := &ListNode{}
	ptr := head
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			ptr.Next = l1
			l1 = l1.Next
		} else {
			ptr.Next = l2
			l2 = l2.Next
		}
		ptr = ptr.Next
	}

	if l1 != nil {
		ptr.Next = l1
	} else {
		ptr.Next = l2
	}

	return head.Next


	//step1 := l1
	//step2 := l2
	//
	//if l1 == nil {
	//	return l2
	//}
	//
	//if l2 == nil {
	//	return l1
	//}
	//
	//if l1.Val > l2.Val {
	//	l1, l2 = l2, l1
	//}
	//
	//h := l1
	//
	//
	//for l1 != nil && l2 !=nil {
	//
	//	for l1 != nil && l1.Val <= l2.Val {
	//		step1 = l1
	//		l1 = l1.Next
	//	}
	//
	//	step1.Next = l2
	//
	//	if l1 == nil {
	//		break
	//	}
	//
	//	for l2 != nil && l1 != nil && l2.Val < l1.Val {
	//		step2 = l2
	//		l2 = l2.Next
	//	}
	//
	//	step2.Next = l1
	//
	//}
	//return h
}

func IsPalindromeList(head *ListNode) bool {

	if head == nil || head.Next == nil {
		return true
	}

	fast := head
	slow := head

	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	slow.Next = ReverseList(slow.Next)
	slow = slow.Next

	for slow != nil  {
		if head.Val != slow.Val {
			return false
		}
		head = head.Next
		slow = slow.Next
	}

	return true
}

func PrintList(l *ListNode)  {
	for l.Next != nil {
		fmt.Print(l.Val, "-")
		l = l.Next
	}
	fmt.Println(l.Val)
}

/*
链表练习完成：有两道用C实现的代码：
bool hasCycle(struct ListNode *head) {
    if (head == NULL) {
        return false;
    }
    if (head->next == NULL) {
        return false;
    }
    struct ListNode* current;
    struct ListNode* tmp;

    current = head;

    while (current->next != NULL) {
        tmp = head;
        while(tmp != current && tmp != current->next) {
            tmp = tmp->next;
        }
        if (tmp == current->next) {
                return true;
        }
        current = current->next;
    }
    return false;
}

// Some God's algorithm:
bool hasCycle(struct ListNode *head) {
    struct ListNode *newSign, *oldSign, *k;
    if(head == NULL) return false;
    if(head -> next == head) return true;
    newSign = head -> next;  //将链表翻转
    oldSign = head;
    head -> next = NULL;
    while(newSign!= NULL){
        k = newSign;
        newSign = newSign-> next;
        k-> next = oldSign;
        oldSign = k;
    }
    return k == head;
}

bool hasCycle(struct ListNode *head) {
    struct ListNode *slow=head;
    struct ListNode *fast=head;

        while(fast&&fast->next){

            slow=slow->next;

            fast=fast->next->next;

            if(slow==fast)return true;

        }

        return false;
}

void deleteNode(struct ListNode* node) {
    node->val = node->next->val;
    node->next = node->next->next;
}
 */




