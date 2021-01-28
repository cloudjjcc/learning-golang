package main

import "fmt"

// 求斐波那契数列的第 n 项，n <= 39
func main() {
	fmt.Println(Fibonacci(30))
	fmt.Println(Fibonacci2(30))
}

// 递归法实现，效率极低
func Fibonacci(i int) int {
	if i <= 1 {
		return i
	}
	return Fibonacci(i-2) + Fibonacci(i-1)
}

// 改为循环，效率提高
func Fibonacci2(i int) int {
	if i <= 1 {
		return i
	}
	pre2 := 0
	pre1 := 1
	fib := 1
	for j := 1; j < i; j++ {
		fib = pre2 + pre1
		pre2 = pre1
		pre1 = fib
	}
	return fib
}
