package one_day_exercise

import "strings"

/*

给你一个下标从 0 开始的字符串 s 和一个单词字典 dictionary 。你需要将 s 分割成若干个 互不重叠 的子字符串，每个子字符串都在 dictionary 中出现过。s 中可能会有一些 额外的字符 不在任何子字符串中。

请你采取最优策略分割 s ，使剩下的字符 最少 。



示例 1：

输入：s = "leetscode", dictionary = ["leet","code","leetcode"]
输出：1
解释：将 s 分成两个子字符串：下标从 0 到 3 的 "leet" 和下标从 5 到 8 的 "code" 。只有 1 个字符没有使用（下标为 4），所以我们返回 1 。
示例 2：

输入：s = "sayhelloworld", dictionary = ["hello","world"]
输出：3
解释：将 s 分成两个子字符串：下标从 3 到 7 的 "hello" 和下标从 8 到 12 的 "world" 。下标为 0 ，1 和 2 的字符没有使用，所以我们返回 3 。


提示：

1 <= s.length <= 50
1 <= dictionary.length <= 50
1 <= dictionary[i].length <= 50
dictionary[i] 和 s 只包含小写英文字母。
dictionary 中的单词互不相同。

 */

 // 好题啊， 标准的面试题目， 一遍就过！
// dfs[i] = min(1+dfs[i+1], dfs[i+len(p)])

func minExtraChar(s string, dictionary []string) int {
	n := len(s)
	var dfs func(i int) int
	cache := make([]int, n)
	for i := range cache {
		cache[i] = -1
	}

	dfs = func(i int) int {
		if i == n {
			return 0
		}

		if cache[i] != -1 {
			return cache[i]
		}
		res := n + 1
		// 不选 s[i]
		res = min(res, 1+dfs(i+1))

		// 选 i, 然后枚举看看选哪个。
		for _, p := range dictionary {
			pos := strings.Index(s[i:], p)
			if pos == 0 {
				res = min(res, dfs(i+len(p)))
			}
		}

		cache[i] = res
		return res
	}

	return dfs(0)
}

// dfs[i] = min(1+dfs[i+1], dfs[i+len(p)])
// dp[i] = min(1+dp[i+1], dp[i+len(p)]
// strings.Index 能优化吗？ 应该可以， 把 dictionary 做成 trie 去 match s[i:]
func minExtraChar(s string, dictionary []string) int {
	n := len(s)
	inf := n + 1
	dp := make([]int, n+1)
	for i := 0; i < n; i++ {
		dp[i] = inf
	}
	dp[n] = 0

	for i := n - 1; i >= 0; i-- {
		dp[i] = min(dp[i], 1+dp[i+1])
		for _, p := range dictionary {
			if strings.Index(s[i:], p) == 0 {
				dp[i] = min(dp[i], dp[i+len(p)])
			}
		}
	}
	return dp[0]
}
