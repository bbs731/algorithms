package 定长滑动窗口

/***
给你字符串 s 和整数 k 。

请返回字符串 s 中长度为 k 的单个子字符串中可能包含的最大元音字母数。

英文中的 元音字母 为（a, e, i, o, u）。



示例 1：

输入：s = "abciiidef", k = 3
输出：3
解释：子字符串 "iii" 包含 3 个元音字母。
示例 2：

输入：s = "aeiou", k = 2
输出：2
解释：任意长度为 2 的子字符串都包含 2 个元音字母。

示例 3：

输入：s = "leetcode", k = 3
输出：2
解释："lee"、"eet" 和 "ode" 都包含 2 个元音字母。

示例 4：

输入：s = "rhythms", k = 4
输出：0
解释：字符串 s 中不含任何元音字母。

示例 5：

输入：s = "tryhard", k = 4
输出：1

 */

func isOwl(d byte) bool {
	if d == 'a' || d == 'e' || d == 'i' || d == 'o' || d == 'u' {
		return true
	}
	return false
}

// 定长 k 的滑动窗口
func maxVowels(s string, k int) int {
	left := 0
	right := min(len(s)-1, k-1)
	ans := 0

	ds := 0 // this can be reused between windows
	// 初始化 ds
	for i := left; i < right; i++ {
		d := s[i]
		if isOwl(d) {
			ds += 1
		}

	}

	for right <= len(s)-1 {
		if isOwl(s[right]) {
			ds++
		}

		ans = max(ans, ds)

		if isOwl(s[left]) {
			ds--
		}
		left++
		right++
	}
	return ans
}
