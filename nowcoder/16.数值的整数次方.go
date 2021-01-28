package main

import (
	"fmt"
	"math"
)

//题目描述
//给定一个 double 类型的浮点数 base 和 int 类型的整数 exponent，求 base 的 exponent 次方。

func main() {
	base := 2.0
	exp := 120
	fmt.Println(power(base, exp))
	fmt.Println(math.Pow(base, float64(exp)))
}

// 快速幂算法
// pow(x,n)=pow(x,n/2)*pow(x,n/2),x为偶数
// pow(x,n)=pow(x,n/2)*pow(x,n/2)*x，x为奇数
func power(base float64, exp int) float64 {
	if exp == 0 {
		return 1
	}
	if exp == 1 {
		return base
	}
	// 是否负指数
	var isNeg bool
	if exp < 0 {
		isNeg = true
		exp = -exp
	}
	pow := power(base*base, exp>>1)
	if exp%2 == 1 {
		pow *= base
	}
	if isNeg {
		return 1 / pow
	}
	return pow
}
