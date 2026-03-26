package nowcoder

// see https://www.nowcoder.com/share/jump/1534770721773718770349

func compare(version1, version2 string) int {
	n1, n2 := len(version1), len(version2)
	for i, j := 0, 0; i < n1 && j < n2; {
		v1 := 0
		for i < n1 && version1[i] != '.' {
			v1 = v1*10 + int(version1[i]-'0')
			i++
		}
		v2 := 0
		for j < n2 && version2[j] != '.' {
			v2 = v2*10 + int(version2[j]-'0')
			j++
		}
		if v1 < v2 {
			return -1
		} else if v1 > v2 {
			return 1
		}
	}
	return 0
}
