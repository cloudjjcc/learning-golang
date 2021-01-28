package main

import "fmt"

//题目描述
//功能:输入一个正整数，按照从小到大的顺序输出它的所有质因子（如180的质因子为2 2 3 3 5 ）
//最后一个数后面也要有空格

func main() {
	getPrimer(24)
}

func getPrimer(n int) {
	for i := 2; i <= n; i++ {
		for n%i == 0 {
			fmt.Printf("%d ", i)
			n /= i
		}
	}
}
