package leetcode

import "math"

//给定两个字符串 s 和 t ，如果它们的编辑距离为 1 ，则返回 true ，否则返回 false 。
//
//字符串 s 和字符串 t 之间满足编辑距离等于 1 有三种可能的情形：
//
//往 s 中插入 恰好一个 字符得到 t
//从 s 中删除 恰好一个 字符得到 t
//在 s 中用 一个不同的字符 替换 恰好一个 字符得到 t
//
//
//示例 1：
//
//输入: s = "ab", t = "acb"
//输出: true
//解释: 可以将 'c' 插入字符串 s 来得到 t。
//示例 2:
//
//输入: s = "cab", t = "ad"
//输出: false
//解释: 无法通过 1 步操作使 s 变为 t。
//
//
//提示:
//
//0 <= s.length, t.length <= 104
//s 和 t 由小写字母，大写字母和数字组成

func isOneEditDistance(s string, t string) bool {
	if s == t || math.Abs(float64(len(s)-len(t))) > 1 || (len(s) == 0 && len(t) == 0) {
		return false
	}
	if len(s) == 0 || len(t) == 0 {
		return true
	}
	tt := len(s)
	if len(t) < len(s) {
		tt = len(t)
		s, t = t, s
	}
	for i := 0; i < tt; i++ {
		if s[i] != t[i] {
			if len(s) == len(t) {
				return s[i+1:] == t[i+1:]
			}
			return s[i:] == t[i+1:]
		}
	}
	return len(t)-len(s) == 1
}
