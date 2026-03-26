package nowcoder

//题目描述
//一只青蛙一次可以跳上1级台阶，也可以跳上2级……它也可以跳上n级。
//求该青蛙跳上一个n级的台阶总共有多少种跳法。

// f(x)=2f(x-1)
func jumpFloor(i int) int {
	if i < 3 {
		return i
	}
	pre1 := 2
	for j := 3; j <= i; j++ {
		pre1 <<= 1
	}
	return pre1
}
