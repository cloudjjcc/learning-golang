package main

import (
	"fmt"
)

//给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。
//在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0)。
//找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
//说明：你不能倾斜容器，且 n 的值至少为 2。
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/container-with-most-water
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}

func maxArea(arr []int) int {
	if len(arr) < 2 {
		return 0
	}
	left, right := 0, len(arr)-1
	getArea := func(a, b int) int {
		weight := b - a
		if arr[a] > arr[b] {
			return arr[b] * weight
		}
		return arr[a] * weight
	}
	getMax := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	max := 0
	for left < right {
		max = getMax(getArea(left, right), max)
		if arr[left] > arr[right] {
			right--
		} else {
			left++
		}
	}
	return max
}
