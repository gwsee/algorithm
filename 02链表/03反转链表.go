package main

/**
题意：反转一个单链表。

示例: 输入: 1->2->3->4->5->NULL 输出: 5->4->3->2->1->NULL
*/

/**
如果再定义一个新的链表，实现链表元素的反转，其实这是对内存空间的浪费。

其实只需要改变链表的next指针的指向，直接将链表反转 ，而不用重新定义一个新的链表，如图所示:
*/
//双指针
func reverseList1(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

//递归
func reverseList2(head *ListNode) *ListNode {
	return helps(nil, head)
}

func helps(pre, head *ListNode) *ListNode {
	if head == nil {
		return pre
	}
	next := head.Next
	head.Next = pre
	return helps(head, next)
}
