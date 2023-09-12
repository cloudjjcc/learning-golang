package leetcode

func canConstruct(ransomNote string, magazine string) bool {
	a := make(map[byte]int)
	for _, v := range magazine {
		a[byte(v)]++
	}
	for _, v := range ransomNote {
		a[byte(v)]--
		if a[byte(v)] < 0 {
			return false
		}
	}
	return true
}
func canConstruct2(ransomNote string, magazine string) bool {
	a := [26]int{}
	for _, v := range magazine {
		a[v-'a']++
	}
	for _, v := range ransomNote {
		a[v-'a']--
		if a[v-'a'] < 0 {
			return false
		}
	}
	return true
}
