package main

//给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target  ，
//写一个函数搜索 nums 中的 target，如果目标值存在返回下标，否则返回 -1。

/**
这两种的区别是在于

在对比后续轮次的时候的较大值(high),是否可能是目标(target)

如果可能是的写法 就用 high = mid - 1 ; 然后 用low<=high
如果不可能是的写法    high = mid;      然后 用low<high
*/

// 左闭右闭区间
// 这种写法的实质是在于 在 arr[mid] > target 的时候  high最新的赋值是可能的target数；所以 才有low<=high
func fun11(arr []int, target int) int {
	if len(arr) < 1 {
		return -1
	}
	high := len(arr) - 1
	low := 0
	for low <= high { //不同点1
		mid := low + (high-low)/2
		if arr[mid] == target {
			return mid
		} else if arr[mid] > target {
			high = mid - 1 //不同点2
		} else {
			low = mid + 1
		}
	}
	return -1
}

// 左闭右开区间
// 这种写法的实质是在于 在 arr[mid] > target 的时候  high不可能是target数；所以 才有low<high
func fun12(arr []int, target int) int {
	if len(arr) < 1 {
		return -1
	}
	high := len(arr) - 1
	low := 0
	for low < high { //不同点1
		mid := low + (high-low)/2
		if arr[mid] == target {
			return mid
		} else if arr[mid] > target {
			high = mid //不同点2
		} else {
			low = mid + 1
		}
	}
	return -1
}
