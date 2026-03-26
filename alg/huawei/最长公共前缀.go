package huawei

func mainlcp() {

}
func lcp(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	lcp1 := func(a, b string) string {
		m, n := len(a), len(b)
		if m == 0 || n == 0 {
			return ""
		}
		prefixLen := m
		if n < m {
			prefixLen = n
		}
		for i := 0; i < prefixLen; i++ {
			if a[i] != b[i] {
				return a[:i]
			}
		}
		return a[:prefixLen]
	}
	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		prefix = lcp1(prefix, strs[i])
	}
	return prefix
}
