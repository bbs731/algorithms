package dp

/***

给定一个字符串数组 words，找到以 words 中每个字符串作为子字符串的最短字符串。如果有多个有效最短字符串满足题目条件，返回其中 任意一个 即可。

我们可以假设 words 中没有字符串是 words 中另一个字符串的子字符串。



示例 1：

输入：words = ["alex","loves","leetcode"]
输出："alexlovesleetcode"
解释："alex"，"loves"，"leetcode" 的所有排列都会被接受。
示例 2：

输入：words = ["catg","ctaagt","gcta","ttca","atgcatc"]
输出："gctaagttcatgcatc"

ctaagt gcta ttca tgcatc

gctaagttcatgcatc


提示：

1 <= words.length <= 12
1 <= words[i].length <= 20
words[i] 由小写英文字母组成
words 中的所有字符串 互不相同

 */

func shortestSuperstring(words []string) string {
	n := len(words)
	m := 1 << uint(n)
	tot := 0
	ans := ""

	for _, w := range words {
		tot += len(w)
		ans += w
	}

	type pair struct {
		p string
		m int
	}
	//visited := make([][]bool, n)
	visited := make(map[pair]bool)
	//for i := range visited {
	//	visited[i] = make([]bool, m)
	//}

	prefix := preCalculation(words)

	var dfs func(int, int, string)
	dfs = func(i int, bitmask int, l string) {
		if bitmask == m-1 {
			if len(l) < len(ans) {
				ans = l
			}
			//return len(l)
			return
		}

		if _, ok := visited[pair{l, bitmask}]; ok {
			return
		}

		//res := tot
		for k, w := range words {
			if (1<<uint(k))&bitmask != 0 {
				continue
			}
			//pl := calculate_prefix(words[i], w)
			pl := prefix[i][k]

			if len(l)+len(w[pl:]) >= len(ans) {
				continue
			}
			//res = min(res, dfs(k, bitmask|1<<k, l+w[pl:]))
			dfs(k, bitmask|1<<k, l+w[pl:])
		}

		//visited[i][bitmask] = true
		visited[pair{l, bitmask}] = true
		return
	}

	for i, w := range words {
		dfs(i, 1<<i, w)
	}
	return ans
}

func preCalculation(words []string) [][]int {
	n := len(words)
	prefix := make([][]int, n)
	for i := range prefix {
		prefix[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			prefix[i][j] = calculate_prefix(words[i], words[j])
		}
	}
	return prefix
}

func calculate_prefix(s1, s2 string) int {
	//l1 := len(s1)
	l2 := len(s2)
	s := s2 + s1
	n := len(s)

	z := make([]int, n)
	for i, l, r := 1, 0, 0; i < n; i++ {
		if i <= r {
			z[i] = min(z[i-l], r-i+1)
		}
		for i+z[i] < n && s[z[i]] == s[i+z[i]] {
			l, r = i, i+z[i]
			z[i]++
		}
	}

	l := 0
	for i := n - l2; i < n; i++ {
		if z[i] == n-i {
			l = n - i
			break
		}
	}
	return l
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
