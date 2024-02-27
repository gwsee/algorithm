package main

/*
在链表类中实现这些功能：

get(index)：获取链表中第 index 个节点的值。如果索引无效，则返回-1。
addAtHead(val)：在链表的第一个元素之前添加一个值为 val 的节点。插入后，新节点将成为链表的第一个节点。
addAtTail(val)：将值为 val 的节点追加到链表的最后一个元素。
addAtIndex(index,val)：在链表中的第 index 个节点之前添加值为 val  的节点。
	如果 index 等于链表的长度，则该节点将附加到链表的末尾。
	如果 index 大于链表长度，则不会插入节点。如果index小于0，则在头部插入节点。
deleteAtIndex(index)：如果索引 index 有效，则删除链表中的第 index 个节点。

*/

var my MyLink

type MyLink struct {
	Link *ListNode
	Size int
}

func initF2() {
	//my = MyLink{
	//	Link: &base,
	//	Size: len(arr) + 1,
	//}
}

// 单链表实现
func (l *MyLink) get(index int) int {
	if l == nil || index < 0 || index > l.Size {
		return -1
	}
	loop := l.Link
	for i := 1; i < index; i++ { // 遍历到索引所在的节点
		loop = loop.Next
	}
	return loop.Val
	//其实这里可以直接循环 loop的
	//step := 1
	//for loop != nil {
	//	if index == step {
	//		return loop.Val
	//	} else {
	//		loop = loop.Next
	//		step++
	//	}
	//}
	//return -1
}
func (l *MyLink) addAtHead(val int) {
	link := ListNode{
		Val:  val,
		Next: l.Link,
	}
	l.Link = &link
	l.Size++
}
func (l *MyLink) addAtTail(val int) {
	newNode := &ListNode{Val: val} // 创建新节点
	cur := l.Link                  // 设置当前节点为虚拟头节点
	for cur.Next != nil {          // 遍历到最后一个节点
		cur = cur.Next
	}
	cur.Next = newNode // 在尾部添加新节点
	l.Size++           // 链表大小增加1
}
func (l *MyLink) addAtIndex(index, val int) {
	if index < 0 { // 如果索引小于0，设置为0
		index = 0
	} else if index > l.Size { // 如果索引大于链表长度，直接返回
		return
	}

	newNode := &ListNode{Val: val} // 创建新节点
	cur := l.Link                  // 设置当前节点为虚拟头节点
	for i := 1; i < index; i++ {   // 遍历到指定索引的前一个节点
		cur = cur.Next
	}
	newNode.Next = cur.Next // 新节点指向原索引节点
	cur.Next = newNode      // 原索引的前一个节点指向新节点
	l.Size++                // 链表大小增加1
}
func (l *MyLink) deleteAtIndex(index int) {
	if index < 0 || index > l.Size { // 如果索引无效则直接返回
		return
	}
	cur := l.Link                // 设置当前节点为虚拟头节点
	for i := 1; i < index; i++ { // 遍历到要删除节点的前一个节点
		cur = cur.Next
	}
	if cur.Next != nil {
		cur.Next = cur.Next.Next // 当前节点直接指向下下个节点，即删除了下一个节点
	}
	l.Size-- // 注意删除节点后应将链表大小减一
}

// MyDoubleLink 循环双链表实现
type MyDoubleLink struct {
	Link *Node
}
type Node struct {
	Val  int
	Next *Node
	Pre  *Node
}

func initDoubleLink() MyDoubleLink {
	rear := &Node{
		Val:  -1,
		Next: nil,
		Pre:  nil,
	}
	rear.Next = rear
	rear.Pre = rear
	return MyDoubleLink{rear}
}

func (l *MyDoubleLink) Get(index int) int {
	node := l.Link.Next
	for node != l.Link && index > 0 {
		index--
		node = node.Next
	}
	if 0 != index {
		return -1
	}
	return node.Val
}

func (l *MyDoubleLink) AddAtHead(val int) {
	link := l.Link
	node := &Node{
		Val: val,
		//head.Next指向原头节点
		Next: link.Next,
		//head.Pre 指向哑节点
		Pre: link,
	}

	//更新原头节点
	link.Next.Pre = node
	//更新哑节点
	link.Next = node
	//以上两步不能反
}

func (l *MyDoubleLink) AddAtTail(val int) {
	link := l.Link
	rear := &Node{
		Val: val,
		//rear.Next = dummy(哑节点)
		Next: link,
		//rear.Pre = ori_rear
		Pre: link.Pre,
	}

	//ori_rear.Next = rear
	link.Pre.Next = rear
	//update dummy
	link.Pre = rear
	//以上两步不能反
}

// 以下标0开始计算的；0是首位 1是第二位
func (l *MyDoubleLink) AddAtIndex(index int, val int) {
	head := l.Link.Next
	//head = MyLinkedList[index]
	for head != l.Link && index > 0 {
		head = head.Next
		index--
	}
	if index > 0 {
		return
	}
	node := &Node{
		Val: val,
		//node.Next = MyLinkedList[index]
		Next: head,
		//node.Pre = MyLinkedList[index-1]
		Pre: head.Pre,
	}
	//MyLinkedList[index-1].Next = node
	head.Pre.Next = node
	//MyLinkedList[index].Pre = node
	head.Pre = node
	//以上两步不能反
}

func (l *MyDoubleLink) DeleteAtIndex(index int) {
	//链表为空
	if l.Link.Next == l.Link {
		return
	}
	head := l.Link.Next
	//head = MyLinkedList[index]
	for head.Next != l.Link && index > 0 {
		head = head.Next
		index--
	}
	//验证index有效
	if index == 0 {
		//MyLinkedList[index].Pre = index[index-2]
		head.Next.Pre = head.Pre
		//MyLinedList[index-2].Next = index[index]
		head.Pre.Next = head.Next
		//以上两步顺序无所谓
	}
}
