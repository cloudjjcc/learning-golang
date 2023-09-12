package leetcode

func combine(n int, k int) [][]int {
	var dfsFn func(kk int)
	ans := make([][]int, 0)
	var tmp []int
	dfsFn = func(cur int) {
		if len(tmp)+(n-cur+1) < k {
			return
		}
		if len(tmp) == k {
			dst := make([]int, k)
			copy(dst, tmp)
			ans = append(ans, dst)
			return
		}
		tmp = append(tmp, cur)
		dfsFn(cur + 1)
		tmp = tmp[:len(tmp)-1]
		dfsFn(cur + 1)
	}
	dfsFn(1)
	return ans
}

func combine1(n int, k int) [][]int {
	var dfsFn func(kk int)
	ans := make([][]int, 0)
	var path []int
	dfsFn = func(start int) {
		if len(path) == k {
			t := make([]int, k)
			copy(t, path)
			ans = append(ans, t)
			return
		}
		for i := start; i <= n-(k-len(path))+1; i++ {
			path = append(path, i)
			dfsFn(i + 1)
			path = path[:len(path)-1]
		}
	}
	dfsFn(1)
	return ans
}
