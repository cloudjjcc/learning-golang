package leetcode

import (
	"strings"
	"unsafe"
)

// 给你一个字符串 s ，仅反转字符串中的所有元音字母，并返回结果字符串。
//
// 元音字母包括 'a'、'e'、'i'、'o'、'u'，且可能以大小写两种形式出现不止一次。
//
// 示例 1：
//
// 输入：s = "hello"
// 输出："holle"
// 示例 2：
//
// 输入：s = "leetcode"
// 输出："leotcede"
//
// 提示：
//
// 1 <= s.length <= 3 * 105
// s 由 可打印的 ASCII 字符组成
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode.cn/problems/reverse-vowels-of-a-string
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func reverseVowels(s string) string {
	ss := []byte(s)
	i, j := 0, len(s)-1
	for i < j {
		if !strings.Contains("aeiouAEIOU", s[i:i+1]) {
			i++
			continue
		}
		if !strings.Contains("aeiouAEIOU", s[j:j+1]) {
			j--
			continue
		}
		ss[i], ss[j] = ss[j], ss[i]
		i++
		j--
	}
	return *(*string)(unsafe.Pointer(&ss))
}
