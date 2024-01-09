package _600_1900

/*

你将会得到一份单词表 words，一个字母表 letters （可能会有重复字母），以及每个字母对应的得分情况表 score。

请你帮忙计算玩家在单词拼写游戏中所能获得的「最高得分」：能够由 letters 里的字母拼写出的 任意 属于 words 单词子集中，分数最高的单词集合的得分。

单词拼写游戏的规则概述如下：

玩家需要用字母表 letters 里的字母来拼写单词表 words 中的单词。
可以只使用字母表 letters 中的部分字母，但是每个字母最多被使用一次。
单词表 words 中每个单词只能计分（使用）一次。
根据字母得分情况表score，字母 'a', 'b', 'c', ... , 'z' 对应的得分分别为 score[0], score[1], ..., score[25]。
本场游戏的「得分」是指：玩家所拼写出的单词集合里包含的所有字母的得分之和。


示例 1：

输入：words = ["dog","cat","dad","good"], letters = ["a","a","c","d","d","d","g","o","o"], score = [1,0,9,5,0,0,3,0,0,0,0,0,0,0,2,0,0,0,0,0,0,0,0,0,0,0]
输出：23
解释：
字母得分为  a=1, c=9, d=5, g=3, o=2
使用给定的字母表 letters，我们可以拼写单词 "dad" (5+1+5)和 "good" (3+2+2+5)，得分为 23 。
而单词 "dad" 和 "dog" 只能得到 21 分。
示例 2：

输入：words = ["xxxz","ax","bx","cx"], letters = ["z","a","b","c","x","x","x"], score = [4,4,4,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,5,0,10]
输出：27
解释：
字母得分为  a=4, b=4, c=4, x=5, z=10
使用给定的字母表 letters，我们可以组成单词 "ax" (4+5)， "bx" (4+5) 和 "cx" (4+5) ，总得分为 27 。
单词 "xxxz" 的得分仅为 25 。
示例 3：

输入：words = ["leetcode"], letters = ["l","e","t","c","o","d"], score = [0,0,1,1,1,0,0,0,0,0,0,1,0,0,1,0,0,0,0,1,0,0,0,0,0,0]
输出：0
解释：
字母 "e" 在字母表 letters 中只出现了一次，所以无法组成单词表 words 中的单词。


提示：

1 <= words.length <= 14
1 <= words[i].length <= 15
1 <= letters.length <= 100
letters[i].length == 1
score.length == 26
0 <= score[i] <= 10
words[i] 和 letters[i] 只包含小写的英文字母。

 */

// 难度 1882, 代码一次过。 信心，信心！
// 一道 DP 的题目， 可以用递归 + cache 完美的解决。
func maxScoreWords(words []string, letters []byte, score []int) int {
	lm := make(map[int]int, 26)
	n := len(words)

	for _, e := range letters {
		lm[int(e-'a')]++
	}

	// 这个加 cache 的想法， 是灵神，在一道题里面给的。对 golang 很实用，只要是能比较的 struct 就可以作为 map key, 并且可以combine 多个field 作为cache key
	type key struct {
		i      int
		prefix [26]int
	}

	cache := make(map[key]int)

	var dfs func(int, [26]int) int
	dfs = func(i int, prefix [26]int) int {
		if i == n {
			// caculate the score
			total := 0
			for i := 0; i < 26; i++ {
				total += prefix[i] * score[i]
			}
			return total
		}

		if v, ok := cache[key{i, prefix}]; ok {
			return v
		}
		// 不选 ith word
		res := dfs(i+1, prefix)

		// 选 ith word
		can := true
		newPrefix := prefix
		for _, c := range words[i] {
			ic := int(c - 'a')
			newPrefix[ic]++
			if newPrefix[ic] > lm[ic] {
				can = false
				break
			}
		}
		if can {
			res = max(res, dfs(i+1, newPrefix))
		}
		cache[key{i, prefix}] = res
		return res

	}
	return dfs(0, [26]int{})
}
