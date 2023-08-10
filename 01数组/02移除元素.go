package main

// 暴力法
func fun20(arr []int, val int) int {
	size := len(arr)
	for i := 0; i < size; i++ {
		if arr[i] == val {
			for j := i; j < size-1; j++ {
				arr[j] = arr[j+1]
			}
			size--
			i--
		}
	}
	return size
}

// 移除指针  快慢指针法
func fun21(arr []int, val int) int {
	slow := 0
	for fast := 0; fast < len(arr); fast++ {
		if arr[fast] != val {
			arr[slow] = arr[fast]
			slow++
		}
	}
	//arr = arr[:slow]
	return slow
}

// 双向指针法
func fun22(arr []int, val int) int {
	left := 0
	right := len(arr) - 1
	for left <= right {
		// 不断寻找左侧的val和右侧的非val 找到时交换位置 目的是将val全覆盖掉
		for left <= right && arr[left] != val {
			//左边第一个等于目标值的
			left++
		}
		for left <= right && arr[right] == val {
			//右边第一个不等于目标值的
			right--
		}
		//各自找到后开始覆盖 覆盖后继续寻找
		if left < right {
			//进行数据替换
			arr[left] = arr[right]
			left++
			right--
		}
	}
	return left
}
