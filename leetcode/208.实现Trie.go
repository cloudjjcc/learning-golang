package leetcode

type Trie struct {
	nodes [26]*Trie
	isEnd bool
}

func NewTrie() Trie {
	return Trie{}
}

func (t *Trie) Insert(word string) {
	cur := t
	for i := 0; i < len(word); i++ {
		ch := word[i] - 'a'
		if cur.nodes[ch] == nil {
			cur.nodes[ch] = &Trie{}
		}
		cur = cur.nodes[ch]
	}
	cur.isEnd = true
}
func (t *Trie) SearchPrefix(prefix string) *Trie {
	cur := t
	for i := 0; i < len(prefix); i++ {
		ch := prefix[i] - 'a'
		if cur.nodes[ch] == nil {
			return nil
		}
		cur = cur.nodes[ch]
	}
	return cur
}
func (t *Trie) Search(word string) bool {
	n := t.SearchPrefix(word)
	return n != nil && n.isEnd
}

func (t *Trie) StartsWith(prefix string) bool {
	return t.SearchPrefix(prefix) != nil
}
