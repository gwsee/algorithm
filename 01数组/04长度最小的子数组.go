package main

func fun41(arr []int, target int) int {
	i := 0
	l := len(arr)   // 数组长度
	sum := 0        // 子数组之和
	result := l + 1 // 初始化返回长度为l+1，目的是为了判断“不存在符合条件的子数组，返回0”的情况
	for j := 0; j < l; j++ {
		sum += arr[j]
		for sum >= target {
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
