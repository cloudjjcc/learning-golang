package main

import (
	"fmt"
)

//题目描述
//给定一个数组和滑动窗口的大小，找出所有滑动窗口里数值的最大值。
//例如，如果输入数组{2,3,4,2,6,2,5,1}及滑动窗口的大小3，
//那么一共存在6个滑动窗口，
//他们的最大值分别为{4,4,6,6,6,5}；
//针对数组{2,3,4,2,6,2,5,1}的滑动窗口有以下6个：
//{[2,3,4],2,6,2,5,1}， {2,[3,4,2],6,2,5,1}， {2,3,[4,2,6],2,5,1}，
//{2,3,4,[2,6,2],5,1}， {2,3,4,2,[6,2,5],1}， {2,3,4,2,6,[2,5,1]}。

func main() {
	fmt.Println(maxInWindows([]int{4, 3, 2, 1, 0, 1, 2, 3, 4}, 3))
}

func maxInWindows(arr []int, size int) []int {
	res := make([]int, 0)
	if len(arr) == 0 {
		return res
	}
	window := make([]int, 0, size)
	for i, v := range arr {
		// verify max index
		if len(window) > 0 && i-window[0] >= size {
			window = window[1:]
		}
		// discarding the small value
		for len(window) > 0 && arr[window[len(window)-1]] < v {
			window = window[:len(window)-1]
		}
		window = append(window, i)
		// collect the max value
		if i >= size-1 {
			res = append(res, arr[window[0]])
		}
	}
	return res
}
