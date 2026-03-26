package nowcoder

import (
	"math"
)

//题目描述
//把只包含质因子2、3和5的数称作丑数（Ugly Number）。
//例如6、8都是丑数，但14不是，因为它包含质因子7。
//习惯上我们把1当做是第一个丑数。求按从小到大的顺序的第N个丑数。

func getUglyNum(n int) int {
	// less 7,all num is ugly
	if n < 7 {
		return n
	}
	// min for nums
	minNum := func(nums ...int) int {
		min := math.MaxInt64
		for i := 0; i < len(nums); i++ {
			if nums[i] < min {
				min = nums[i]
			}
		}
		return min
	}
	//
	var (
		t2, t3, t5 = 0, 0, 0
	)
	tmp := make([]int, 0)
	tmp = append(tmp, 1)
	for i := 1; i < n; i++ {
		tmp = append(tmp, minNum(tmp[t2]*2, tmp[t3]*3, tmp[t5]*5))
		if tmp[i] == tmp[t2]*2 {
			t2++
		}
		if tmp[i] == tmp[t3]*3 {
			t3++
		}
		if tmp[i] == tmp[t5]*5 {
			t5++
		}
	}
	return tmp[n-1]
}
