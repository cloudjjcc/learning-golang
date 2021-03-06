package nowcoder

//题目描述
//一只青蛙一次可以跳上1级台阶，也可以跳上2级。
//求该青蛙跳上一个n级的台阶总共有多少种跳法（先后次序不同算不同的结果）。

// f(x)=f(x-1)+f(x-2)
func jump(i int) int {
	if i < 3 {
		return i
	}
	pre1 := 2
	pre2 := 1
	max := 0
	for j := 3; j <= i; j++ {
		max = pre1 + pre2
		pre2 = pre1
		pre1 = max
	}
	return max
}
