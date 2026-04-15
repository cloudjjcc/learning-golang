package main

import "fmt"

// 数组A,`2*n`个元素,n个奇数、n个偶数,设计一个算法,使得数组奇数下标位置放置的都是奇数,偶数下标位置放置的都是偶数..

func main() {
	a := []int{2, 1, 3, 4, 7, 6, 9, 8}

	even, odd := 0, 1
	n := len(a)
	for even < n && odd < n {
		for even < n && a[even]%2 == 0 {
			even += 2
		}
		for odd < n && a[odd]%2 != 0 {
			odd += 2
		}
		if even < n && odd < n {
			a[even], a[odd] = a[odd], a[even]
			even += 2
			odd += 2
		}
	}
	fmt.Println(a)
}
