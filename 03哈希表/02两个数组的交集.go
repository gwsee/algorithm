package main

import "fmt"

/**
题意：给定两个数组，编写一个函数来计算它们的交集。

说明： 输出结果中的每个元素一定是唯一的。 我们可以不考虑输出结果的顺序。
*/
func testF2() {
	var a = []int{1, 2, 3, 4}
	var b = []int{4, 3, 4, 1}
	fmt.Println(intersection(a, b))
}

func intersection(nums1 []int, nums2 []int) []int {
	set := make(map[int]struct{}, 0) // 用map模拟set
	res := make([]int, 0)
	for _, v := range nums1 {
		if _, ok := set[v]; !ok {
			set[v] = struct{}{}
		}
	}
	for _, v := range nums2 {
		//如果存在于上一个数组中，则加入结果集，并清空该set值
		if _, ok := set[v]; ok {
			res = append(res, v)
			delete(set, v)
		}
	}
	return res
}
