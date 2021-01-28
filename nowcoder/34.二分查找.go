package main

import "fmt"

//题目描述
//对于一个有序数组，我们通常采用二分查找的方式来定位某一元素，请编写二分查找的算法，在数组中查找指定元素。
//
//给定一个整数数组A及它的大小n，同时给定要查找的元素val，请返回它在数组中的位置(从0开始)，若不存在该元素，返回-1。若该元素出现多次，请返回第一次出现的位置。
//
//测试样例：
//[1,3,5,7,9],5,3
//返回：1

func main() {
	testArr := []int{1, 11}
	target := 11
	fmt.Println(binarySearch(testArr, target))
}

// O(log n)
func binarySearch(arr []int, target int) int {
	length := len(arr)
	if length == 0 || target < arr[0] || target > arr[length-1] {
		return -1
	}
	var (
		start = 0
		mid   = 0
		end   = length - 1
	)
	for start <= end {
		mid = (start + end) / 2
		if arr[mid] > target {
			end = mid - 1
		} else if arr[mid] < target {
			start = mid + 1
		} else {
			return mid
		}
	}
	return -1
}
