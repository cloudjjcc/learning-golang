package main

import "fmt"

//矩形以列表 [x1, y1, x2, y2] 的形式表示，其中 (x1, y1) 为左下角的坐标，(x2, y2) 是右上角的坐标。
//如果相交的面积为正，则称两矩形重叠。需要明确的是，只在角或边接触的两个矩形不构成重叠。
//给出两个矩形，判断它们是否重叠并返回结果。

func main() {
	fmt.Println(isRectangleOverlap([]int{0, 0, 2, 2}, []int{1, 1, 3, 3}))
}

// 方法一：检查位置
func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	return !(rec1[1] >= rec2[3] || //rec1在rec2上方
		rec1[3] <= rec2[1] || //rec1在rec2下方
		rec1[2] <= rec2[0] || //rec1在rec2左侧
		rec1[0] >= rec2[2]) //rec1在rec2右侧
}
