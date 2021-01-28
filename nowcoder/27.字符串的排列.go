package main

import "fmt"

//题目描述
//输入一个字符串,按字典序打印出该字符串中字符的所有排列。
//例如输入字符串abc,
//则打印出由字符a,b,c所能排列出来的所有字符串abc,acb,bac,bca,cab和cba。
//输入描述:
//输入一个字符串,长度不超过9(可能有字符重复),字符只包括大小写字母。
func main() {
	fmt.Println(len(permutation("abcdefghi")))
}

// 将字符串转为字符数组arr
// 转化为子问题：arr[0]+arr[1:]
func permutation(str string) []string {
	bytes := []byte(str)
	return permutationCore(bytes[:0], bytes)
}

func permutationCore(pre, next []byte) (res []string) {
	if len(next) < 1 {
		res = append(res, string(pre))
		return
	}
	if len(next) == 1 {
		tmp := append(pre, next...)
		res = append(res, string(tmp))
		return
	}
	for i := 0; i < len(next); i++ {
		next[0], next[i] = next[i], next[0]
		npre := append(pre, next[0])
		res = append(res, permutationCore(npre, next[1:])...)
	}
	return res
}
