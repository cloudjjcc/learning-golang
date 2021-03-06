package nowcoder

import (
	"github.com/cloudjjcc/go-exercises/datastructures"
)

//题目描述
//编写一个程序，将输入字符串中的字符按如下规则排序。
//规则 1 ：英文字母从 A 到 Z 排列，不区分大小写。
//如，输入： Type 输出： epTy
//规则 2 ：同一个英文字母的大小写同时存在时，按照输入顺序排列。
//如，输入： BabA 输出： aABb
//规则 3 ：非英文字母的其它字符保持原来的位置。
//如，输入： By?e 输出： Be?y
//注意有多组测试数据，即输入有多行，每一行单独处理（换行符隔开的表示不同行）

func getResult(s string) string {
	queue := &datastructures.Queue{}
	for c := byte(0); c < 26; c++ {
		for i := 0; i < len(s); i++ {
			if s[i]-'a' == c || s[i]-'A' == c {
				queue.Enqueue(s[i])
			}
		}
	}
	buf := make([]byte, len(s))
	for i := range s {
		if (s[i] <= 'z' && s[i] >= 'a') || (s[i] <= 'Z' && s[i] >= 'A') {
			buf[i] = queue.Dequeue().(byte)
		} else {
			buf[i] = s[i]
		}
	}
	return string(buf)
}
