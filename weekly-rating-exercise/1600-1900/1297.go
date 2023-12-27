package _600_1900

/*


给你一个字符串 s ，请你返回满足以下条件且出现次数最大的 任意 子串的出现次数：

子串中不同字母的数目必须小于等于 maxLetters 。
子串的长度必须大于等于 minSize 且小于等于 maxSize 。


示例 1：

输入：s = "aababcaab", maxLetters = 2, minSize = 3, maxSize = 4
输出：2
解释：子串 "aab" 在原字符串中出现了 2 次。
它满足所有的要求：2 个不同的字母，长度为 3 （在 minSize 和 maxSize 范围内）。
示例 2：

输入：s = "aaaa", maxLetters = 1, minSize = 3, maxSize = 3
输出：2
解释：子串 "aaa" 在原字符串中出现了 2 次，且它们有重叠部分。
示例 3：

输入：s = "aabcabcab", maxLetters = 2, minSize = 2, maxSize = 3
输出：3
示例 4：

输入：s = "abcde", maxLetters = 2, minSize = 3, maxSize = 3
输出：0


提示：

1 <= s.length <= 10^5
1 <= maxLetters <= 26
1 <= minSize <= maxSize <= min(26, s.length)
s 只包含小写英文字母。

 */

/*

难度分数： 1748

和 916题差不多的思路， 但是为啥难度分高了100分？
枚举的力量！ 26 个chars
一遍过！ 蒙的吗？

maxSize 是没有用的，想想为什么？ 考虑一下后缀，那样的substring 不可能比 minSize 的 substring 出现的次数多。

 */

func maxFreq(s string, maxLetters int, minSize int, maxSize int) int {
	ans := 0
	candidates := make(map[string]int)
	for i := minSize - 1; i < len(s); i++ {
		cnt := [26]int{}
		pass := true
		newc := 0
		for j := i - minSize + 1; j <= i; j++ {
			d := s[j] - 'a'
			cnt[d]++
			if cnt[d] == 1 {
				newc++
			}
			if newc > maxLetters {
				pass = false
				break
			}
		}
		if pass {
			pattern := s[i-minSize+1 : i+1]
			candidates[pattern]++
			ans = max(ans, candidates[pattern])
		}
	}
	return ans
}
