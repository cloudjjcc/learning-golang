package main

import (
	"fmt"
	"sort"
)

//题目描述
//明明想在学校中请一些同学一起做一项问卷调查，为了实验的客观性，
//他先用计算机生成了N个1到1000之间的随机整数（N≤1000），对于其中重复的数字，
//只保留一个，把其余相同的数去掉，不同的数对应着不同的学生的学号。
//然后再把这些数从小到大排序，按照排好的顺序去找同学做调查。
//请你协助明明完成“去重”与“排序”的工作(同一个测试用例里可能会有多组数据，希望大家能正确处理)。
//
//注：测试用例保证输入参数的正确性，答题者无需验证。测试用例不止一组。

func main() {
	n := 0
	tmp := 0
	for {
		b, err := fmt.Scanln(&n)
		if b == 0 || err != nil {
			return
		}
		arr := make([]int, n)
		for i := 0; i < n; i++ {
			b, err = fmt.Scanln(&tmp)
			if b == 0 || err != nil {
				return
			}
			arr[i] = tmp
		}
		// 排序
		sort.Ints(arr)
		for i, v := range arr {
			// 去重
			if i == 0 || v != arr[i-1] {
				fmt.Println(v)
			}
		}
	}
}
