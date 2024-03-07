package main

/**

给定一个链表，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。

为了表示给定链表中的环，使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。

如果 pos 是 -1，则在该链表中没有环。

说明：不允许修改给定的链表。

输入: head =[3,2,0,-4]，pos = 1
输出: tail connects to node index 1
解释: 链表中有一个环，其尾部连接到第二个节点

输入: head =[1,2]，pos = 0
输出: tail connects to node index 0
解释: 链表中有一个环，其尾部连接到第一个节点。
*/

/**
我们使用 Floyd 的 Tortoise and Hare（龟兔赛跑）算法，
即快慢指针方法。当快指针和慢指针第一次相遇时，我们可以确定链表中存在环。
然后，将其中一个指针重置到链表的头部，并与另一个指针同步移动，直到它们再次相遇。
它们再次相遇的节点就是环的起始节点。
*/
/**
那么相遇时： slow指针走过的节点数为: x + y， fast指针走过的节点数：x + y + n (y + z)，
n为fast指针在环内走了n圈才遇到slow指针， （y+z）为 一圈内节点的个数A

因为fast指针是一步走两个节点，slow指针一步走一个节点， 所以 fast指针走过的节点数 = slow指针走过的节点数 * 2：

(x + y) * 2 = x + y + n (y + z)

两边消掉一个（x+y）: x + y = n (y + z)

因为要找环形的入口，那么要求的是x，因为x表示 头结点到 环形入口节点的的距离。

所以要求x ，将x单独放在左面：x = n (y + z) - y ,

再从n(y+z)中提出一个 （y+z）来，整理公式之后为如下公式：x = (n - 1) (y + z) + z 注意这里n一定是大于等于1的，因为 fast指针至少要多走一圈才能相遇slow指针。

这个公式说明什么呢？

先拿n为1的情况来举例，意味着fast指针在环形里转了一圈之后，就遇到了 slow指针了。

当 n为1的时候，公式就化解为 x = z，
当 n>1的时候 , 其实就是相对于,n-1 圈后 再与慢指针相遇
*/
func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			for slow != head {
				slow = slow.Next
				head = head.Next
			}
			return head
		}
	}
	return nil
}

var node7 ListNode

func initF7() {
	// 构建一个有环的链表，环的起始节点是3
	// 1 -> 2 -> 3 -> 4 -> 5
	//           |_________|
	nodea := ListNode{Val: 1}
	nodeb := &ListNode{Val: 2}
	nodec := &ListNode{Val: 3}
	noded := &ListNode{Val: 4}
	nodee := &ListNode{Val: 5}

	nodea.Next = nodeb
	nodeb.Next = nodec
	nodec.Next = noded
	noded.Next = nodee
	nodee.Next = nodec // 创建环

	node7 = nodea
}
func testF7() {
	node := detectCycle(&node7)
	printNode4(node)
}
