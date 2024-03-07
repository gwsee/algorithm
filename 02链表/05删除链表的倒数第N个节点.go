package main

import "fmt"

/**
给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。

进阶：你能尝试使用一趟扫描实现吗？
*/
var node5 ListNode
var arr5 = []int{1, 2, 3, 4, 5, 6, 7}

func initF5() {
	var loop *ListNode
	for _, v := range arr5 {
		if loop == nil {
			node5 = ListNode{
				Val:  v,
				Next: nil,
			}
			loop = &node5
			continue
		}
		temp := ListNode{
			Val:  v,
			Next: nil,
		}
		loop.Next = &temp
		loop = &temp
	}
}

func testF5() {
	printNode4(&node5)
	removeNthFromEnd(&node5, 2)
	fmt.Println("remove after")
	printNode4(&node5)
}
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyHead := &ListNode{}
	dummyHead.Next = head //使用虚拟节点 避免考虑头节点的问题
	cur := head
	prev := dummyHead
	i := 1
	for cur != nil {
		cur = cur.Next
		if i > n {
			//使用的快慢指针，并先执行n+1步骤
			//这样
			//快指针跑到尾部的时候
			//慢指针还在n-1的时候 即 prev
			//最后将prev.Next = prev.Next.Next  就实现了删除
			prev = prev.Next
		}
		i++
	}
	prev.Next = prev.Next.Next
	return dummyHead.Next
}
