package main

import (
	"errors"
	"fmt"
	"math"
	"regexp"
	"sort"
)

func First(num int64) (res int64) {
	if num < 3 {
		return num
	}
	return First(num-1) + First(num-2)
}

// Second 参考https://www.wangt.cc/2021/04/%E8%9E%BA%E6%97%8B%E6%8A%98%E7%BA%BF-%E7%AC%AC%E4%B9%9D%E5%B1%8A%E8%93%9D%E6%A1%A5%E6%9D%AFc-cb%E7%BB%84/
func Second(x, y float64) (res float64) {
	var n float64
	if math.Abs(x) <= y && y > 0 { //上
		n = y
		res = (2*n-1)*(2*n) + (x - (-n))
	} else if math.Abs(y) <= x && x > 0 { //右
		n = x
		res = (2*n)*(2*n) + (n - y)
	} else if math.Abs(x) <= math.Abs(y)+1 && y <= 0 { //下
		n = math.Abs(y)
		res = (2*n)*(2*n+1) + (n - x)
	} else if math.Abs(y) <= math.Abs(x) && x < 0 { //左
		n = math.Abs(x)
		res = (2*n-1)*(2*n-1) + (y - (-n + 1))
	}
	return
}

func Third(str string) (res string, err error) {
	regRule := "^[1-9]\\d*$"
	re := regexp.MustCompile(regRule)
	result := re.FindAllStringSubmatch(str, -1)
	if result == nil {
		err = fmt.Errorf("%v is not a valid num", str)
		return
	}
	arr := make([]string, len(str))
	for i := 0; i < len(str); i++ {
		arr[i] = string(str[i])
	}
	nums := createNums(arr)
	sort.Strings(nums)
	//fmt.Println(nums[0], nums[len(nums)-1])
	index := -1
	for k, v := range nums {
		if v > str {
			index = k
			break
		}
	}
	if index == -1 {
		err = errors.New("this num is max num")
		return
	}
	res = nums[index]
	return
}

func createNums(data []string) (res []string) {
	if len(data) == 1 {
		return data
	}
	for i := 0; i < len(data); i++ {
		var next []string
		if i == 0 {
			next = data[i+1:]
		} else if i == len(data)-1 {
			next = data[:i]
		} else {
			next = append(next, data[:i]...)
			next = append(next, data[i+1:]...)
			//比如下面的data中2的下标会被3下标值覆盖
			//fore := data[:i]
			//ends := data[i+1:]
			//next = append(fore, ends...)
			//或者改成
			//next = append(next, fore...)
			//next = append(next, ends...)
		}
		nextRes := createNums(next)
		key := data[i]
		for _, v := range nextRes {
			res = append(res, fmt.Sprintf("%v%v", key, v))
		}

	}
	return
}
