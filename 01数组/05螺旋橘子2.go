package main

import "fmt"

/*
给定一个正整数 n，生成一个包含 1 到 n^2 所有元素，且元素按顺时针顺序螺旋排列的正方形矩阵。

示例:

输入: 3 输出: [ [ 1, 2, 3 ], [ 8, 9, 4 ], [ 7, 6, 5 ] ]
*/

func fun51(n int) (res [][]int) {
	res = make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}
	loop := n / 2          // 每个圈循环几次，例如n为奇数3，那么loop = 1 只是循环一圈，矩阵中间的值需要单独处理
	startx, starty := 0, 0 // 定义每循环一个圈的起始位置
	count := 1
	offset := 1
	for loop > 0 {
		i, j := startx, starty
		// 下面开始的四个for就是模拟转了一圈
		// 模拟填充上行从左到右(左闭右开)
		//行数不变 列数在变
		for j = starty; j < n-offset; j++ {
			res[startx][j] = count
			count++
		}
		//列数不变是j 行数变
		// 模拟填充右列从上到下(左闭右开)
		for i = startx; i < n-offset; i++ {
			res[i][j] = count
			count++
		}
		//行数不变 i 列数变 j--
		// 模拟填充下行从右到左(左闭右开)
		for ; j > starty; j-- {
			res[i][j] = count
			count++
		}
		//列不变 行变
		// 模拟填充左列从下到上(左闭右开)
		for ; i > startx; i-- {
			res[i][j] = count
			count++
		}
		// 第二圈开始的时候，起始位置要各自加1， 例如：第一圈起始位置是(0, 0)，第二圈起始位置是(1, 1)
		startx++
		starty++
		offset++
		loop--
	}

	center := n / 2 // 矩阵中间的位置，例如：n为3， 中间的位置就是(1，1)，n为5，中间位置为(2, 2)
	if n%2 == 1 {
		res[center][center] = n * n
	}
	return
}

func fun51Print(res [][]int) {
	for _, v := range res {
		fmt.Println(v)
	}
}
