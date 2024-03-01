package main
/**
给你两个单链表的头节点 headA 和 headB ，
	请你找出并返回两个单链表相交的起始节点。如果两个链表没有交点，返回 null 。
 */

var node61 ListNode
var node62 ListNode
var arr61 = []int{32,341,221,21}
var arr62 = []int{2,1,2}
func initF6()  {
	initF61()
	initF62()
	var sames = []int{5,2,3}
	var nodeSame ListNode
	var loopSame *ListNode
	for _,v:=range sames{
		if loopSame == nil {
			nodeSame = ListNode{
				Val:  v,
				Next: nil,
			}
			loopSame = &nodeSame
			continue
		}
		temp:=ListNode{
			Val:  v,
			Next: nil,
		}
		loopSame.Next = &temp
		loopSame = &temp
	}
	node61.Next = &nodeSame
	node62.Next = &nodeSame
}
func initF61() {
	var loop *ListNode
	for _, v := range arr61 {
		if loop == nil {
			node61 = ListNode{
				Val:  v,
				Next: nil,
			}
			loop = &node61
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
func initF62() {
	var loop *ListNode
	for _, v := range arr62 {
		if loop == nil {
			node62 = ListNode{
				Val:  v,
				Next: nil,
			}
			loop = &node62
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

func testF6()  {
	same:=getIntersectionNode2(&node61,&node62)
	printNode4(same)
}

func getIntersectionNode1(headA, headB *ListNode) *ListNode {
	curA := headA
	curB := headB
	lenA, lenB := 0, 0
	// 求A，B的长度
	for curA != nil {
		curA = curA.Next
		lenA++
	}
	for curB != nil {
		curB = curB.Next
		lenB++
	}
	var step int
	var fast, slow *ListNode
	// 请求长度差，并且让更长的链表先走相差的长度
	if lenA > lenB {
		step = lenA - lenB
		fast, slow = headA, headB
	} else {
		step = lenB - lenA
		fast, slow = headB, headA
	}
	for i:=0; i < step; i++ {
		fast = fast.Next
	}
	// 遍历两个链表遇到相同则跳出遍历
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}
	return fast
}


func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	l1,l2 := headA, headB
	for l1 != l2 {
		if l1 != nil {
			l1 = l1.Next
		} else {
			l1 = headB
		}

		if l2 != nil {
			l2 = l2.Next
		} else {
			l2 = headA
		}
	}

	return l1
}