package main

import "fmt"

//给你一份『词汇表』（字符串数组） words 和一张『字母表』（字符串） chars。
//假如你可以用 chars 中的『字母』（字符）拼写出 words 中的某个『单词』（字符串），
//那么我们就认为你掌握了这个单词。
//注意：每次拼写时，chars 中的每个字母都只能用一次。
//返回词汇表 words 中你掌握的所有单词的 长度之和。
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/find-words-that-can-be-formed-by-characters
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	fmt.Println(countCharacters([]string{"cat", "bt", "hat", "tree"}, "atach"))
}

func countCharacters(words []string, chars string) int {
	if len(chars) == 0 || len(words) == 0 {
		return 0
	}
	var tmp [26]int
	for i := 0; i < len(chars); i++ {
		tmp[chars[i]-'a']++
	}
	res := 0
loop:
	for i := 0; i < len(words); i++ {
		ttmep := tmp
		for j := 0; j < len(words[i]); j++ {
			if ttmep[words[i][j]-'a'] != 0 {
				ttmep[words[i][j]-'a']--
			} else {
				continue loop
			}
		}
		res += len(words[i])
	}
	return res
}
