package small_representation

/***
给你一个字符串 s ，找出它的所有子串并按字典序排列，返回排在最后的那个子串。

示例 1：

输入：s = "abab"
输出："bab"
解释：我们可以找出 7 个子串 ["a", "ab", "aba", "abab", "b", "ba", "bab"]。按字典序排在最后的子串是 "bab"。
示例 2：

输入：s = "leetcode"
输出："tcode"


提示：

1 <= s.length <= 4 * 10^5
s 仅含有小写英文字符。

 */

/***
结合了 oi-wiki 和 灵神的模板写的。
 */
func lastSubstring(s string) string {

	smallestRepresentation := func(s string) string {
		n := len(s)
		i, k, j := 0, 0, 1
		// 循环不变量 i <= j
		for ; k < n && j+k < n; {
			if s[i+k] == s[j+k] { // 注意这里是 if, 不是灵神的板子用了 loop
				k++
			} else {
				if s[i+k] > s[j+k] { // 改成 > 则返回字典序最大的
					// j 到 j+k 都不会是最小串的开头位置
					j += k + 1
				} else {
					// i 到 i+k 都不会是最小串的开头位置
					i, j = j, max(j, i+k)+1
				}
				k = 0 // 别忘了重置 0 因为 s[i+k] 和 s[j+k] 不相等了
			}
		}
		return s[i:n]
	}
	return smallestRepresentation(s)
}

func lastSubstring(s string) string {
	smallestRepresentation := func(s string) string {
		n := len(s)
		s = s + s
		i, k := 0, 0
		for j := 1; i < n; {
			k = 0
			for k < n && j+k < n && s[i+k] == s[j+k] { // 按照灵神的模板，这里不加 j+k<n 的条件，这道题过不去。
				k++
			}
			/**
				按照灵神的模板， 这里需要，多加判断
			 */
			if k == n || j+k >= n { // 按照灵神的模板， 这里不加 j+k 的条件，这道题过不去。
				break
			}

			if s[i+k] > s[j+k] { // 改成 > 则返回字典序最大的
				// j 到 j+k 都不会是最小串的开头位置
				j += k + 1
			} else {
				// i 到 i+k 都不会是最小串的开头位置
				i, j = j, max(j, i+k)+1
			}
		}
		return s[i:n]
	}
	return smallestRepresentation(s)
}

/***
这个是按照 oi-wiki 上的模板写的
https://oi-wiki.org/string/minimal-string/
 */
func lastSubstring(s string) string {
	biggestRepresentation := func(s string) string {
		n := len(s)
		i, j, k := 0, 1, 0
		for ; k < n && i+k < n && j+k < n; {
			if s[i+k] == s[j+k] {
				k++
			} else {
				if s[i+k] > s[j+k] {
					// j 到 j+k+1 都不会是最小串的开头位置
					j = j + k + 1
				} else {
					i = i + k + 1
				}

				if i == j {
					i++
				}
				k = 0
			}
		}
		return s[min(i, j):n]
	}
	return biggestRepresentation(s)
}
