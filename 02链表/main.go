package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	testF1()
}

func testF1() {
	var base = ListNode{
		Val:  1,
		Next: nil,
	}
	arr := []int{1, 2, 3, 41, 1, 3, 4, 1, 23, 4}
	loop := &base
	for _, v := range arr {
		temp := ListNode{
			Val:  v,
			Next: nil,
		}
		loop.Next = &temp
		loop = &temp
	}
	//方法1
	by, _ := json.Marshal(fun11(&base, 1))
	fmt.Println(string(by))
	////方法2
	//by, _ := json.Marshal(fun12(&base, 1))
	//fmt.Println(string(by))
}
