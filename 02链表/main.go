package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	//testF1()
	//testF2()
	testF4()
}

func init() {
	//initF1()
	//initF2()
	initF4()
	return
}

func testF1() {
	//方法1
	by, _ := json.Marshal(fun11(&base, 1))
	fmt.Println(string(by))
	////方法2
	//by, _ := json.Marshal(fun12(&base, 1))
	//fmt.Println(string(by))
}
func testF2() {
	//单链表测试
	//my1 := my
	//my1.addAtHead(100)
	//my1.addAtTail(1)
	//my1.addAtIndex(2, 5)
	//fmt.Println(my1.get(3))
	//my1.deleteAtIndex(3)
	//fmt.Println(my1.get(3))

	//双链表测试
	my2 := initDoubleLink()
	my2.AddAtHead(55)
	my2.AddAtTail(3)
	my2.AddAtIndex(1, 22)
	fmt.Println(my2.Get(1))
	my2.DeleteAtIndex(1)
	fmt.Println(my2.Get(1))

}
