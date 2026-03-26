package leetcode

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	words := [26]int{}
	for _, v := range s {
		words[v-'a']++
	}
	for _, v := range t {
		words[v-'a']--
	}
	for _, v := range words {
		if v != 0 {
			return false
		}
	}
	return true
}
