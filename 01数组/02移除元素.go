package main

//给你一个数组 nums 和一个值 val，你需要 原地 移除所有数值等于 val 的元素，并返回移除后数组的新长度。
//不要使用额外的数组空间，你必须仅使用 O(1) 额外空间并原地修改输入数组。
//元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。

// 暴力法
func fun20(arr []int, val int) int {
	size := len(arr)
	for i := 0; i < size; i++ {
		if arr[i] == val {
			for j := i; j < size-1; j++ {
				//将后面的数据往前移1为  注意 size 必须为size-1  ;这样j+1就不会产生溢出
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
			//其实就是重新定义不等于目标值的数据,到数组中的位置
			arr[slow] = arr[fast]
			slow++
		}
	}
	//arr = arr[:slow]
	return slow
}

// 双向指针法
func fun22(nums []int, val int) int {
	left := 0
	//原本的写法看不懂
	//right := len(arr) - 1
	//for left <= right {
	//	// 不断寻找左侧的val和右侧的非val 找到时交换位置 目的是将val全覆盖掉
	//	for left <= right && arr[left] != val {
	//		//左边第一个等于目标值的
	//		left++
	//	}
	//	for left <= right && arr[right] == val {
	//		//右边第一个不等于目标值的
	//		right--
	//	}
	//	//各自找到后开始覆盖 覆盖后继续寻找
	//	if left < right {
	//		//进行数据替换
	//		arr[left] = arr[right]
	//		left++
	//		right--
	//	}
	//}
	right := len(nums) - 1
	for left <= right {
		if nums[left] == val && nums[right] != val {
			// 交换 left 和 right 的元素
			nums[left], nums[right] = nums[right], nums[left]
		}
		if nums[left] != val {
			left++
		}
		if nums[right] == val {
			right--
		}
	}
	return left
}
