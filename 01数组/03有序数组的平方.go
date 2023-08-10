package main

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
