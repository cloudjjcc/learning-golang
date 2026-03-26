package datastructures

// ITrie 字典树
type ITrie interface {
	// Insert 插入元素 (仅有小写英文字母构成)
	Insert(string)
	// Search 搜索字符串
	Search(string) bool
	// StartWith 是否包含前缀
	StartWith(prefix string) bool
}

var _ ITrie = (*Trie)(nil)

type Trie struct {
	children [26]*Trie
	isEnd    bool
}

func (t *Trie) StartWith(prefix string) bool {
	cur := t
	for i := 0; i < len(prefix); i++ {
		ch := prefix[i] - 'a'
		if cur.children[ch] == nil {
			return false
		}
		cur = cur.children[ch]
	}
	return cur != nil
}

func (t *Trie) Search(s string) bool {
	cur := t
	for i := 0; i < len(s); i++ {
		ch := s[i] - 'a'
		if cur.children[ch] == nil {
			return false
		}
		cur = cur.children[ch]
	}
	return cur != nil && cur.isEnd
}

func (t *Trie) Insert(s string) {
	cur := t
	for i := 0; i < len(s); i++ {
		ch := s[i] - 'a'
		if cur.children[ch] == nil {
			cur.children[ch] = &Trie{}
		}
		cur = cur.children[ch]
	}
	cur.isEnd = true
}
