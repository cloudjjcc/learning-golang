package leetcode

func candy(ratings []int) int {
	tmp := make([]int, len(ratings))
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			tmp[i] = tmp[i-1] + 1
		}
	}
	for i := len(ratings) - 1; i >= 1; i-- {
		if ratings[i] > ratings[i-1] {
			tmp[i] = tmp[i-1] + 1
		}
	}
	ans := 0
	for _, v := range tmp {
		ans += v
	}
	return ans
}
