package leetcode

func longestCommonPrefix(strs []string) string {
	t := new(trie)
	for _, v := range strs {
		t.add(v)
	}
	return t.prefixWithCount(strs[0], len(strs))
}

type trie struct {
	nodes [26]*trie
	count int
}

func (t *trie) prefixWithCount(word string, count int) (prefix string) {
	cur := t
	for i, v := range word {
		tt := cur.nodes[v-'a']
		if tt == nil || tt.count < count {
			break
		}
		prefix = word[:i+1]
		cur = tt
	}
	return prefix
}
func (t *trie) add(word string) {
	if word == "" {
		return
	}
	cur := t
	cur.count++
	for _, v := range word {
		tt := cur.nodes[v-'a']
		if tt == nil {
			tt = &trie{
				nodes: [26]*trie{},
				count: 0,
			}
			cur.nodes[v-'a'] = tt
		}
		tt.count++
		cur = tt
	}
}

//
func longestCommonPrefix2(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	lcp := func(a string, b string) string {
		minLen := len(a)
		if len(b) < minLen {
			minLen = len(b)
		}
		for i := 0; i < minLen; i++ {
			if a[i] != b[i] {
				return a[:i]
			}
		}
		return a[:minLen]
	}
	ans := strs[0]
	for i := 1; i < len(strs); i++ {
		ans = lcp(ans, strs[i])
	}
	return ans
}
