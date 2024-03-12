package main

import "fmt"

/**
给定四个包含整数的数组列表 A , B , C , D ,
	计算有多少个元组 (i, j, k, l) ，使得 A[i] + B[j] + C[k] + D[l] = 0。

为了使问题简单化，所有的 A, B, C, D 具有相同的长度 N，且 0 ≤ N ≤ 500 。
	所有整数的范围在 -2^28 到 2^28 - 1 之间，最终结果不会超过 2^31 - 1 。

*/

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	m := make(map[int]int) //key:a+b的数值，value:a+b数值出现的次数
	count := 0
	// 遍历nums1和nums2数组，统计两个数组元素之和，和出现的次数，放到map中
	for _, v1 := range nums1 {
		for _, v2 := range nums2 {
			m[v1+v2]++
		}
	}
	// 遍历nums3和nums4数组，找到如果 0-(c+d) 在map中出现过的话，就把map中key对应的value也就是出现次数统计出来
	for _, v3 := range nums3 {
		for _, v4 := range nums4 {
			count += m[-v3-v4]
		}
	}
	return count
}

func testF5() {
	arrs1 := []int{3, -4, 8, 1, 7, 2, -5}
	arrs2 := []int{3, -4, 8, 1, 7, 2, -5}
	arrs3 := []int{3, -4, 8, 1, 7, 2, -5}
	arrs4 := []int{3, -4, 8, 1, 7, 2, -5}
	fmt.Println(fourSumCount(arrs1, arrs2, arrs3, arrs4))
}
