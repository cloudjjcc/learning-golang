package main

import "fmt"

//给定一个单词列表，我们将这个列表编码成一个索引字符串 S 与一个索引列表 A。
//例如，如果这个列表是 ["time", "me", "bell"]，我们就可以将其表示为 S = "time#bell#" 和 indexes = [0, 2, 5]。
//对于每一个索引，我们可以通过从字符串 S 中索引的位置开始读取字符串，直到 "#" 结束，来恢复我们之前的单词列表。
//那么成功对给定单词列表进行编码的最小字符串长度是多少呢？
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/short-encoding-of-words
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	fmt.Println(minimumLengthEncoding([]string{"time", "me", "bell"}))
}

func minimumLengthEncoding(words []string) int {
	trie := &TrieNode{}
	nodes := make(map[*TrieNode]int)

	for i, s := range words {
		cur := trie
		for j := len(s) - 1; j >= 0; j-- {
			cur = cur.Get(s[j])
		}
		nodes[cur] = i
	}
	ans := 0
	for node, i := range nodes {
		if node.freq == 0 {
			ans += len(words[i]) + 1
		}
	}
	return ans
}

type TrieNode struct {
	childNodes [26]*TrieNode
	freq       int
}

func (t *TrieNode) Get(ch byte) *TrieNode {
	if t.childNodes[ch-'a'] == nil {
		t.childNodes[ch-'a'] = new(TrieNode)
		t.freq++
	}
	return t.childNodes[ch-'a']
}
