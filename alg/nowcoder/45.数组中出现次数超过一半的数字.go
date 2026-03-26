package nowcoder

//题目描述
//数组中有一个数字出现的次数超过数组长度的一半，请找出这个数字。
//例如输入一个长度为9的数组{1,2,3,2,2,2,5,4,2}。
//由于数字2在数组中出现了5次，超过数组长度的一半，因此输出2。如果不存在则输出0。

func moreThanHalfNum(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	if len(arr) == 1 {
		return arr[0]
	}
	// if exist more than half num,it will be res
	res := arr[0]
	times := 1
	for i := 1; i < len(arr); i++ {
		if times == 0 {
			res = arr[i]
			times = 1
		} else if arr[i] == res {
			times++
		} else {
			times--
		}
	}
	// check res
	times = 0
	for i := 0; i < len(arr); i++ {
		if res == arr[i] {
			times++
		}
	}
	if times<<1 > len(arr) {
		return res
	}
	return 0
}
