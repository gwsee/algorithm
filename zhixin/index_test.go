package main

import (
	"fmt"
	"testing"
)

func TestFirst(t *testing.T) {
	fmt.Println(First(4))
}
func TestSecond(t *testing.T) {
	fmt.Println(Second(-10, -10))
}
func TestThird(t *testing.T) {
	res, err := Third("1258462481")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(res)
	}
}
