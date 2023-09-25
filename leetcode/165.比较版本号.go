package leetcode

func compareVersion(version1 string, version2 string) int {
	i, j := 0, 0
	for i < len(version1) || j < len(version2) {
		a, b := 0, 0
		for i < len(version1) && version1[i] != '.' {
			a = a*10 + int(version1[i]-'0')
			i++
		}
		i++
		for j < len(version2) && version2[j] != '.' {
			b = b*10 + int(version2[j]-'0')
			j++
		}
		j++
		if a > b {
			return 1
		}
		if a < b {
			return -1
		}
	}
	return 0
}
