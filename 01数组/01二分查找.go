package main

// 左闭右闭区间
// 这种写法的实质是在于 在 arr[mid] > target 的时候  high最新的赋值是可能的target数；所以 才有low<=high
func fun11(arr []int, target int) int {
	if len(arr) < 1 {
		return -1
	}
	high := len(arr) - 1
	low := 0
	for low <= high {
		mid := low + (high-low)/2
		if arr[mid] == target {
			return mid
		} else if arr[mid] > target {
			high = mid - 1
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
	for low < high {
		mid := low + (high-low)/2
		if arr[mid] == target {
			return mid
		} else if arr[mid] > target {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return -1
}
