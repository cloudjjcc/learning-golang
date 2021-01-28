package main

import "fmt"

//题目描述
//•连续输入字符串，请按长度为8拆分每个字符串后输出到新的字符串数组；
//•长度不是8整数倍的字符串请在后面补数字0，空字符串不处理。

func main() {
	for i := 0; i < 2; i++ {
		str := ""
		b, err := fmt.Scanln(&str)
		if b == 0 || err != nil {
			return
		}
		m := len(str) % 8
		bytes := []byte(str)
		// fill zero
		if m != 0 {
			for i := 0; i < 8-m; i++ {
				bytes = append(bytes, '0')
			}
		}
		//
		for i := 0; i < len(bytes); i += 8 {
			fmt.Println(string(bytes[i : i+8]))
		}
	}
}
