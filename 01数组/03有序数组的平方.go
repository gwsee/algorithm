package main
/**
给你一个按 非递减顺序 排序的整数数组 nums，返回 每个数字的平方 组成的新数组，要求也按 非递减顺序 排序。

示例 1：

输入：nums = [-4,-1,0,3,10]
输出：[0,1,9,16,100]
解释：平方后，数组变为 [16,1,0,9,100]，排序后，数组变为 [0,1,9,16,100]
示例 2：

输入：nums = [-7,-3,2,3,11]
输出：[4,9,9,49,121]
 */
// 暴力解法
func fun31(arr []int) (res []int) {
	for _, v := range arr {
		res = append(res, v*v)
	}
	//sort.Ints(res)
	//或者自己写个排序方式
	for i := 0; i < len(res); i++ {
		for j := i + 1; j < len(res); j++ {
			if res[i] > res[j] {
				min := res[j]
				res[j] = res[i]
				res[i] = min
			}
		}
	}
	return
}

// 双指针法
// 因为arr是有序的 在第一个获取到的数据之后 就能够知道 他肯定是最大的 所以可以使用这种方式
func fun32(arr []int) (res []int) {
	n := len(arr)
	i, j, k := 0, n-1, n-1
	res = make([]int, n)
	for i <= j {
		lm, rm := arr[i]*arr[i], arr[j]*arr[j]
		if lm > rm {
			res[k] = lm
			i++
		} else {
			res[k] = rm
			j--
		}
		k--
	}
	return
}
