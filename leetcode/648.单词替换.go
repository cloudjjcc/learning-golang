package leetcode

import "unsafe"

type trie1 struct {
	nodes [26]*trie1
	count int
}

func (t *trie1) prefixWithShortPath(word string) (prefix string) {
	cur := t
	for i, v := range word {
		tt := cur.nodes[v-'a']
		if tt == nil {
			return ""
		}
		prefix = word[:i+1]
		cur = tt
		if cur.count > 0 {
			return prefix
		}
	}
	return prefix
}
func (t *trie1) add(word string) {
	if word == "" {
		return
	}
	cur := t
	cur.count++
	for _, v := range word {
		tt := cur.nodes[v-'a']
		if tt == nil {
			tt = &trie1{
				nodes: [26]*trie1{},
				count: 0,
			}
			cur.nodes[v-'a'] = tt
		}
		cur = tt
	}
	cur.count++
}
func replaceWords(dictionary []string, sentence string) string {
	t := new(trie1)
	for _, v := range dictionary {
		t.add(v)
	}
	target := make([]byte, 0, len(sentence))
	wordStart := 0
	for i := 0; i < len(sentence); i++ {
		if sentence[i] == ' ' || i == len(sentence)-1 {
			if i == len(sentence)-1 {
				i++
			}
			if wordStart != 0 {
				target = append(target, ' ')
			}
			prefix := t.prefixWithShortPath(sentence[wordStart:i])
			if prefix != "" {
				target = append(target, []byte(prefix)...)
			} else {
				target = append(target, sentence[wordStart:i]...)
			}
			wordStart = i + 1
		}
	}
	return *(*string)(unsafe.Pointer(&target))
}
