package nowcoder

//题目描述
//汇编语言中有一种移位指令叫做循环左移（ROL），现在有个简单的任务，
//就是用字符串模拟这个指令的运算结果。对于一个给定的字符序列S，
//请你把其循环左移K位后的序列输出。
//例如，字符序列S=”abcXYZdef”,要求输出循环左移3位后的结果，即“XYZdefabc”。
//是不是很简单？OK，搞定它！

// 利用切片
// 额外空间 O(n)
func leftRotateString(str string, n int) string {
	if len(str) == 0 {
		return str
	}
	n %= len(str)
	return str[n:] + str[:n]
}

// 翻转字符串
func leftRotateString2(str string, n int) string {
	if len(str) == 0 {
		return str
	}
	n %= len(str)
	buf := []byte(str)
	// reverse str
	reverse := func(arr []byte) {
		if len(arr) < 2 {
			return
		}
		var (
			left  = 0
			right = len(arr) - 1
		)
		for left < right {
			arr[left], arr[right] = arr[right], arr[left]
			left++
			right--
		}
	}
	// reverse 0:n
	reverse(buf[:n])
	reverse(buf[n:])
	reverse(buf)
	return string(buf)
}
