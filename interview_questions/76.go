package interview_questions

/***

给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。


注意：

对于 t 中重复字符，我们寻找的子字符串中该字符数量必须不少于 t 中该字符数量。
如果 s 中存在这样的子串，我们保证它是唯一的答案。


示例 1：

输入：s = "ADOBECODEBANC", t = "ABC"
输出："BANC"
解释：最小覆盖子串 "BANC" 包含来自字符串 t 的 'A'、'B' 和 'C'。
示例 2：

输入：s = "a", t = "a"
输出："a"
解释：整个字符串 s 是最小覆盖子串。
示例 3:

输入: s = "a", t = "aa"
输出: ""
解释: t 中两个字符 'a' 均应包含在 s 的子串中，
因此没有符合条件的子字符串，返回空字符串。

 */

/***
sliding window + hashtable ?

wolala, 这道题应该能过关把
 */

func minWindow(s string, t string) string {
	//n := len(s)
	m := len(t)
	if m == 0 {
		return ""
	}
	ct := make(map[int32]int, 26)
	for _, tc := range t {
		ct[tc]++
	}

	ans := ""
	satisfy := func(sm, tm map[int32]int) bool {
		for k, v := range tm {
			if sm[k] < v {
				return false
			}
		}
		return true
	}
	l := 0
	st := make(map[int32]int, 26)
	for i, sc := range s { // 枚举右端点
		st[sc]++

		for l <= i && satisfy(st, ct) {
			if ans == "" || len(ans) > i-l+1 {
				ans = s[l : i+1]
			}
			// pop l
			st[int32(s[l])]--
			l++
		}
	}
	return ans
}
