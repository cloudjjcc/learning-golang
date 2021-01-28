package main

import "fmt"

//把一根绳子剪成多段，并且使得每段的长度乘积最大。
//n = 2
//return 1 (2 = 1 + 1)
//n = 10
//return 36 (10 = 3 + 3 + 4)
func main() {
	fmt.Println(cutRope2(8))
}

// f(x)=max(i*f(x-i))
func cutRope(i int) int {
	if i < 2 {
		return 0
	}
	if i == 2 {
		return 1
	}
	if i == 3 {
		return 2
	}
	return cutRopeCore(i)
}
func cutRopeCore(n int) int {
	if n < 4 {
		return n
	}
	maxValue := 0
	for j := 1; j < n; j++ {
		tmp := j * cutRopeCore(n-j)
		if tmp > maxValue {
			maxValue = tmp
		}
	}
	return maxValue
}

// 动态规划求解
func cutRope2(i int) int {
	if i < 2 {
		return 0
	}
	if i == 2 {
		return 1
	}
	if i == 3 {
		return 2
	}
	tmps := make([]int, i+1)
	tmps[0] = 0
	tmps[1] = 1
	tmps[2] = 2
	tmps[3] = 3
	//
	for j := 4; j <= i; j++ {
		maxValue := 0
		for k := 1; k < j; k++ {
			tmp := k * tmps[j-k]
			if maxValue < tmp {
				maxValue = tmp
			}
		}
		tmps[j] = maxValue
	}
	return tmps[i]
}
