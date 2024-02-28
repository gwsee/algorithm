package main

import "fmt"

func swapPairs1(head *ListNode) *ListNode {
	dummy := &ListNode{
		Next: head,
	}
	//head=list[i]
	//pre=list[i-1]
	pre := dummy
	for head != nil && head.Next != nil {
		pre.Next = head.Next
		next := head.Next.Next
		head.Next.Next = head
		head.Next = next
		//pre=list[(i+2)-1]
		pre = head
		//head=list[(i+2)]
		head = next
	}
	return dummy.Next
}

// 递归版本
func swapPairs2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	next := head.Next
	head.Next = swapPairs2(next.Next)
	next.Next = head
	return next
}

var node4 ListNode
var arr4 = []int{1, 2, 3, 41, 1, 3, 4, 1, 23, 4}

func initF4() {
	var loop *ListNode
	for _, v := range arr4 {
		if loop == nil {
			node4 = ListNode{
				Val:  v,
				Next: nil,
			}
			loop = &node4
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

func testF4() {
	printNode4(&node4)
	swapPairs2(&node4)
	fmt.Println("swap after")
	printNode4(&node4)
}

func printNode4(head *ListNode)  {
	if head == nil {
		return
	}
	fmt.Println(head.Val)
	printNode4(head.Next)
}