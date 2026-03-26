package leetcode

// 给你一个字符数组 s ，反转其中 单词 的顺序。
//
// 单词 的定义为：单词是一个由非空格字符组成的序列。s 中的单词将会由单个空格分隔。
//
// 必须设计并实现 原地 解法来解决此问题，即不分配额外的空间。
//
// 示例 1：
//
// 输入：s = ["t","h","e"," ","s","k","y"," ","i","s"," ","b","l","u","e"]
// 输出：["b","l","u","e"," ","i","s"," ","s","k","y"," ","t","h","e"]
// 示例 2：
//
// 输入：s = ["a"]
// 输出：["a"]
//
// 提示：
//
// 1 <= s.length <= 105
// s[i] 可以是一个英文字母（大写或小写）、数字、或是空格 ' ' 。
// s 中至少存在一个单词
// s 不含前导或尾随空格
// 题目数据保证：s 中的每个单词都由单个空格分隔
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode.cn/problems/reverse-words-in-a-string-ii
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func reverseWords(s []byte) {
	reverseFn := func(a, b int) { //反转序列
		for a < b {
			s[a], s[b] = s[b], s[a]
			a++
			b--
		}
	}
	reverseFn(0, len(s)-1)
	j := 0
	for i := 0; i < len(s); i++ { //反转单词
		if s[i] == ' ' {
			reverseFn(j, i-1)
			j = i + 1
		}
	}
	reverseFn(j, len(s)-1)
}
