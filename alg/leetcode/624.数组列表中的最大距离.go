package leetcode

func maxDistance(arrays [][]int) int {
	var (
		max, min, maxDis = arrays[0][len(arrays[0])-1], arrays[0][0], 0
	)
	maxFn := func(a int, b ...int) int {
		tmax := a
		for _, v := range b {
			if v > tmax {
				tmax = v
			}
		}
		return tmax
	}
	for _, v := range arrays[1:] {
		maxDis = maxFn(maxDis, max-v[0], v[len(v)-1]-min)
		if v[0] < min {
			min = v[0]
		}
		if v[len(v)-1] > max {
			max = v[len(v)-1]
		}
	}
	return maxDis
}
