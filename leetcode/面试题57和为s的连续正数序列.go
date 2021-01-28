package main

import "fmt"

//输入一个正整数 target ，输出所有和为 target 的连续正整数序列（至少含有两个数）。
//序列内的数字由小到大排列，不同序列按照首个数字从小到大排列。
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/he-wei-sde-lian-xu-zheng-shu-xu-lie-lcof
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	fmt.Println(findContinuousSequence(100))
}

func findContinuousSequence(target int) [][]int {
	p1 := 1
	p2 := 2
	sum := p1 + p2
	res := make([][]int, 0)
	addSequence := func(start, end int) {
		tmp := make([]int, 0)
		for i := start; i <= end; i++ {
			tmp = append(tmp, i)
		}
		res = append(res, tmp)
	}
	for p1 < p2 {
		if sum > target {
			sum -= p1
			p1++
		} else if sum < target {
			p2++
			sum += p2
		} else {
			addSequence(p1, p2)
			p2++
			sum += p2
			sum -= p1
			p1++
		}
	}
	return res
}
