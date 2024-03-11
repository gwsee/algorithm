package main

import "fmt"

/**
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。

你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。

示例:

给定 nums = [2, 7, 11, 15], target = 9

因为 nums[0] + nums[1] = 2 + 7 = 9

所以返回 [0, 1]


*/

// 暴力解法
func twoSum1(nums []int, target int) []int {
	for k1, _ := range nums {
		for k2 := k1 + 1; k2 < len(nums); k2++ {
			if target == nums[k1]+nums[k2] {
				return []int{k1, k2}
			}
		}
	}
	return []int{}
}

// 使用map方式解题，降低时间复杂度
func twoSum2(nums []int, target int) []int {
	m := make(map[int]int)
	for index, val := range nums {
		if preIndex, ok := m[target-val]; ok {
			return []int{preIndex, index}
		} else {
			m[val] = index
		}
	}
	return []int{}
}
func testF4() {
	target := 10
	arrs := []int{3, 5, 4, 1, 7, 2, 8}
	fmt.Println(twoSum1(arrs, target))
}
