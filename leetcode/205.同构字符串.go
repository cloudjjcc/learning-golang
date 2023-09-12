package leetcode

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	a := make(map[byte]byte)
	b := make(map[byte]byte)
	for i := 0; i < len(s); i++ {
		v, ok := a[s[i]]
		if !ok {
			if vv, ok := b[t[i]]; ok && vv != s[i] {
				return false
			}
			a[s[i]] = t[i]
			b[t[i]] = s[i]
			continue
		}
		if v != t[i] {
			return false
		}
	}
	return true
}
