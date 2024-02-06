package subsequenceAutomation

import "fmt"

/***
给定字符串 s 和字符串数组 words, 返回  words[i] 中是s的子序列的单词个数 。

字符串的 子序列 是从原始字符串中生成的新字符串，可以从中删去一些字符(可以是none)，而不改变其余字符的相对顺序。

例如， “ace” 是 “abcde” 的子序列。


示例 1:

输入: s = "abcde", words = ["a","bb","acd","ace"]
输出: 3
解释: 有三个是 s 的子序列的单词: "a", "acd", "ace"。
Example 2:

输入: s = "dsahjpjauf", words = ["ahjpjau","ja","ahbwzgqnuk","tnmlanowax"]
输出: 2

提示:
1 <= s.length <= 5 * 10^4
1 <= words.length <= 5000
1 <= words[i].length <= 50
words[i]和 s 都只由小写字母组成。

 */

func numMatchingSubseq(s string, words []string) int {
	subsequenceAutomation := func(s string) [][26]int {
		pos := [26]int{}
		for i := range pos {
			pos[i] = len(s)
		}
		nxt := make([][26]int, len(s))

		for i := len(s) - 1; i >= 0; i-- {
			nxt[i] = pos
			pos[s[i]-'a'] = i
		}

		return nxt
	}
	nxt := subsequenceAutomation(s)
	match := func(word, s string, nxt [][26]int) bool {
		i, j := 0, 0
		if word[0] == s[0] {
			j = 1
		}

		for ; j < len(word); j++ {
			i = nxt[i][word[j]-'a']
			if i == len(s) {
				break
			}
		}
		// matched
		if j == len(word) {
			return true
		}
		return false
	}

	cnts := 0
	for _, word := range words {
		if match(word, s, nxt) {
			cnts++
		}
	}
	return cnts
}

/****
这样写也是可以的， 但是为了记忆，还是按照灵神的板子来吧。
 */

func numMatchingSubseq(s string, words []string) int {
	subsequenceAutomation := func(s string) [][26]int {
		pos := [26]int{}
		for i := range pos {
			pos[i] = len(s) + 2
		}
		nxt := make([][26]int, len(s)+1)
		for i := len(s) - 1; i >= 0; i-- {
			nxt[i+1] = pos
			pos[s[i]-'a'] = i + 1
		}
		nxt[0] = pos // 灵神的版本 是没有保存最后一个 pos的, 处理的方法，留在了 match function 去比较一下 t[0] == s[0]
		return nxt
	}

	nxt := subsequenceAutomation(s)
	match := func(word, s string, nxt [][26]int) bool {
		i, j := 0, 0
		for ; j < len(word); j++ {
			i = nxt[i][word[j]-'a']
			if i == len(s)+2 {
				break
			}
		}
		// matched
		if j == len(word) {
			return true
		}
		return false
	}

	cnts := 0
	for _, word := range words {
		if match(word, s, nxt) {
			fmt.Println(word)
			cnts++
		}
	}
	return cnts
}
