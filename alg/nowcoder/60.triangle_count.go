package nowcoder

import (
	"sort"
)

//给定一个包含非负整数的数组，你的任务是统计其中可以组成三角形三条边的三元组个数。
//
//示例 1:
//
//输入: [2,2,3,4]
//输出: 3
//解释:
//有效的组合是:
//2,3,4 (使用第一个 2)
//2,3,4 (使用第二个 2)
//2,2,3
//注意:
//
//数组长度不超过1000。
//数组里整数的范围为 [0, 1000]。

func triangleCount(arr []int) int {
	if len(arr) < 3 {
		return 0
	}
	// sort
	sort.Ints(arr)
	// find a,b,c
	count := 0
	for i := 0; i < len(arr); i++ {
		left, right := 0, i-1
		for left < right {
			if arr[left]+arr[right] > arr[i] {
				count += right - left
				right--
			} else {
				left++
			}

		}
	}
	return count
}
