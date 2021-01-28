package main

import "fmt"

//题目描述
//把M个同样的苹果放在N个同样的盘子里，允许有的盘子空着不放，问共有多少种不同的分法？
//（用K表示）5，1，1和1，5，1 是同一种分法。
func main() {
	for {
		m, n := 0, 0
		b, err := fmt.Scanln(&m, &n)
		if b == 0 || err != nil {
			return
		}
		fmt.Println(divideApple(m, n))
	}
}

func divideApple(m int, n int) int {
	if m == 0 || n == 1 {
		return 1
	}
	if m < n {
		return divideApple(m, m)
	}
	return divideApple(m, n-1) + divideApple(m-n, n)
}
