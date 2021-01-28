package main

import (
	"fmt"
	"github.com/cloudjjcc/go-exercises/datastructures"
)

//题目描述
//每年六一儿童节,牛客都会准备一些小礼物去看望孤儿院的小朋友,今年亦是如此。
//HF作为牛客的资深元老,自然也准备了一些小游戏。
//其中,有个游戏是这样的:首先,让小朋友们围成一个大圈。
//然后,他随机指定一个数m,让编号为0的小朋友开始报数。
//每次喊到m-1的那个小朋友要出列唱首歌,
//然后可以在礼品箱中任意的挑选礼物,并且不再回到圈中,
//从他的下一个小朋友开始,继续0...m-1报数....这样下去....
//直到剩下最后一个小朋友,可以不用表演,
//并且拿到牛客名贵的“名侦探柯南”典藏版(名额有限哦!!^_^)。
//请你试着想下,哪个小朋友会得到这份礼品呢？(注：小朋友的编号是从0到n-1)
//如果没有小朋友，请返回-1

func main() {
	fmt.Println(lastRemaining(100, 3))
	fmt.Println(lastRemaining2(100, 3))
	fmt.Println(lastRemaining3(100, 3))
}

// 数学归纳法
//        { 0,n=1
// f(n,m)= (f(n-1,m)+m)%n,n>1
//        }
func lastRemaining(n int, m int) int {
	if n == 1 {
		return 0
	}
	pre := 0
	for i := 2; i <= n; i++ {
		pre = (pre + m) % i
	}
	return pre
}

// 数组实现
func lastRemaining2(n, m int) int {
	if n == 1 {
		return 0
	}
	tmp := make([]int, n)
	for i := 0; i < n; i++ {
		tmp[i] = i
	}
	for len(tmp) > 1 {
		cur := -1
		for i := 0; i < m; i++ {
			cur++
			if cur == len(tmp) {
				cur = 0
			}
		}
		tmp = append(tmp[cur+1:], tmp[:cur]...)
	}
	return tmp[0]
}

// 约瑟夫环（Josephuse）
func lastRemaining3(n, m int) int {
	if n == 1 {
		return 0
	}
	// init a circular list
	list := &datastructures.ListNode{
		Value: 0,
	}
	cur := list
	for i := 1; i < n; i++ {
		cur.Next = &datastructures.ListNode{Value: i}
		cur = cur.Next
	}
	cur.Next = list
	count := 0
	for cur.Next != cur {
		if count == 2 {
			cur.Next = cur.Next.Next
			count = 0
		}
		cur = cur.Next
		count++
	}
	return cur.Value.(int)
}
