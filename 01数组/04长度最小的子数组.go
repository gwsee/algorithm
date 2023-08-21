package main

/*
*
给定一个含有 n 个正整数的数组和一个正整数 s ，找出该数组中满足其和 ≥ s 的长度最小的 连续 子数组，并返回其长度。如果不存在符合条件的子数组，返回 0。

示例：

输入：s = 7, nums = [2,3,1,2,4,3]
输出：2
解释：子数组 [4,3] 是该条件下的长度最小的子数组。
提示：

1 <= target <= 10^9
1 <= nums.length <= 10^5
1 <= nums[i] <= 10^5
*/

func fun41(arr []int, target int) int {
	i := 0
	l := len(arr)   // 数组长度
	sum := 0        // 子数组之和
	result := l + 1 // 初始化返回长度为l+1，目的是为了判断“不存在符合条件的子数组，返回0”的情况
	for j := 0; j < l; j++ {
		sum += arr[j]
		for sum >= target { //条件满足其和大于等于S
			subLength := j - i + 1
			if subLength < result {
				result = subLength
			}
			sum -= arr[i]
			i++
		}
	}
	if result == l+1 {
		return 0
	} else {
		return result
	}
	return 0
}
