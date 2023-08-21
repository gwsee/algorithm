package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	//testF1()
	testF2()
}

func init() {
	initF1()
	initF2()

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
	my1 := my
	my1.addAtHead(100)
	my1.addAtTail(1)
	my1.addAtIndex(2, 5)
	fmt.Println(my1.get(3))
	my1.deleteAtIndex(3)
	fmt.Println(my1.get(3))
}
