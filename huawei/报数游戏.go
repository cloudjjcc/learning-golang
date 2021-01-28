package main

import (
	"fmt"
	"github.com/cloudjjcc/go-exercises/datastructures"
)

//题目描述
//n个人围成一圈，每个人有一个编码，编号从1开始到n.
//他们从1开始依次报数，报到为M的人自动退出圈圈，然后下一个人接着从1开始报数，直到剩余的人数小于M。
//请问最后剩余的人在原先的编号为多少？
//例如输入M=3时，输出为：“58，91”，输入M=4时，输出为： “34，45， 97”。
//如果m小于等于1， 则输出“ERROR!”;
//如果m大于等于100，则输出“ERROR!”；

func main() {
	for {
		n := 0
		s, b := fmt.Scanln(&n)
		if s == 0 || b != nil {
			return
		}
		fmt.Println(game(n))
	}
}
func game(n int) int {
	// init a circular list
	list := &datastructures.ListNode{
		Value: 1,
	}
	cur := list
	for i := 2; i <= n; i++ {
		cur.Next = &datastructures.ListNode{Value: i}
		cur = cur.Next
	}
	// 模拟报数
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
