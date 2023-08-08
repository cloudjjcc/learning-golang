package leetcode

import (
	"strings"
)

//给定一个关键词集合 words 和一个字符串 S，将所有 S 中出现的关键词加粗。
//所有在标签 <b> 和 </b> 中的字母都会加粗。
//返回的字符串需要使用尽可能少的标签，当然标签应形成有效的组合。
//例如，给定 words = ["ab", "bc"] 和 S = "aabcd"，需要返回 "a<b>abc</b>d"。
//注意返回 "a<b>a<b>b</b>c</b>d" 会使用更多的标签，因此是错误的。
//注：
//words 长度的范围为 [0, 50]。
//words[i] 长度的范围为 [1, 10]。
//S 长度的范围为 [0, 500]。
//所有 words[i] 和 S 中的字符都为小写字母。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/bold-words-in-string
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// 暴力解法
func boldWords(words []string, s string) string {
	mask := make([]bool, len(s))
	for _, w := range words {
		tmp := s
		base := 0
		for {
			i := strings.Index(tmp, w)
			if i == -1 {
				break
			}
			for j := 0; j < len(w); j++ {
				mask[base+i+j] = true
			}
			base = base + i + 1
			tmp = s[base:]
		}
	}
	buf := make([]byte, 0)
	const (
		LeftTag  = "<b>"
		RightTag = "</b>"
	)
	for i := 0; i < len(s); i++ {
		if mask[i] && (i == 0 || !mask[i-1]) {
			buf = append(buf, LeftTag...)
		} else if i != 0 && !mask[i] && mask[i-1] {
			buf = append(buf, RightTag...)
		}
		buf = append(buf, s[i])
	}
	if mask[len(mask)-1] {
		buf = append(buf, RightTag...)
	}
	return string(buf)
}

// 使用trie
type TrieNode struct {
	children [26]*TrieNode
}

func (t *TrieNode) add(key string) {
	root := t
	for _, v := range key {
		if root.children[v-'a'] == nil {
			root.children[v-'a'] = &TrieNode{}
		}
		root = root.children[v-'a']
	}
}
func (t *TrieNode) find(key string) {
}
