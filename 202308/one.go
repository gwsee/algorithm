package main

import "fmt"

func main() {
	fmt.Println(function4(3, 4))
	fmt.Println("分割线")
	fmt.Println(function4(3, 2))
	fmt.Println("分割线")
	fmt.Println(function4(3, 3))
	fmt.Println("分割线")
	fmt.Println(function2(3, 4))
}
func function1(x, n int) int {
	result := 1 // 注意 任何数的0次方等于1
	for i := 0; i < n; i++ {
		result = result * x
	}
	return result
}
func function2(x, n int) int {
	if n == 0 {
		return 1 // return 1 同样是因为0次方是等于1的
	}
	num := function2(x, n-1) * x
	fmt.Println(num)
	return num
}
func function4(x, n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	t := function4(x, n/2)
	if n%2 == 1 {
		return t * t * x
	}
	return t * t
}
