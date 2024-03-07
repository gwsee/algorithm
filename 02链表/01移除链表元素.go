package main

import (
	"encoding/json"
	"fmt"
)

/*
题意：删除链表中等于给定值 val 的所有节点。

示例 1： 输入：head = [1,2,6,3,4,5,6], val = 6 输出：[1,2,3,4,5]

示例 2： 输入：head = [], val = 1 输出：[]

示例 3： 输入：head = [7,7,7,7], val = 7 输出：[]
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

var base ListNode
var arr = []int{1, 2, 3, 41, 1, 3, 4, 1, 23, 4}

func initF1() {
	base = ListNode{
		Val:  1,
		Next: nil,
	}
	loop := &base
	for _, v := range arr {
		temp := ListNode{
			Val:  v,
			Next: nil,
		}
		loop.Next = &temp
		loop = &temp
	}
}

// 直接使用原来的链表来进行移除节点操作
func fun11(head *ListNode, val int) *ListNode {
	// 删除头结点
	for head != nil && head.Val == val {
		head = head.Next
	}
	cur := head
	for cur != nil && cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}

// 设置一个虚拟头结点在进行移除节点操作
func fun12(head *ListNode, val int) *ListNode {
	dummyHead := &ListNode{}
	dummyHead.Next = head
	cur := dummyHead
	for cur != nil && cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return dummyHead.Next
}
func testF1() {
	//方法1
	by, _ := json.Marshal(fun11(&base, 1))
	fmt.Println(string(by))
	////方法2
	//by, _ := json.Marshal(fun12(&base, 1))
	//fmt.Println(string(by))
}
