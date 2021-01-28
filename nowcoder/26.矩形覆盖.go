package main

import "fmt"

//题目描述
//我们可以用2*1的小矩形横着或者竖着去覆盖更大的矩形。
//请问用n个2*1的小矩形无重叠地覆盖一个2*n的大矩形，总共有多少种方法？
//比如n=3时，2*3的矩形块有3种覆盖方法：

func main() {
	fmt.Println(getResultRect2(40))
}

// 递归解法
// f(1)=1
// f(2)=2
// f(n)=f(n-1)+f(n-2)
func getResultRect(n int) int {
	if n < 3 {
		return n
	}
	return getResultRect(n-1) + getResultRect(n-2)
}

// 递推解法
func getResultRect2(n int) int {
	if n < 3 {
		return n
	}
	var (
		pre1   = 2
		pre2   = 1
		result = 0
	)
	for i := 3; i <= n; i++ {
		result = pre1 + pre2
		pre2 = pre1
		pre1 = result
	}
	return result
}
