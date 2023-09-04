package leetcode

func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

func fib1(n int) int {
	fibs := make([]int, 31)
	fibs[0] = 0
	fibs[1] = 1
	for i := 2; i <= n; i++ {
		fibs[i] = fibs[i-1] + fibs[i-2]
	}
	return fibs[n]
}

func fib2(n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 1
	}
	n1, n2, cur := 1, 1, 2
	for i := 2; i <= n; i++ {
		cur = n1 + n2
		n1, n2 = n2, cur
	}
	return cur
}
