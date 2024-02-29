package dp

/***

子集型回溯的样板题目！
 */

func maxScoreWords(words []string, letters []byte, score []int) int {
	n := len(words)

	ll := [26]int{}
	for _, e := range letters {
		ll[ e-'a']++
	}

	cost := func(bitmask int) int {
		sum := 0
		total := [26]int{}
		for i := 0; i < n; i++ {
			if bitmask&(1<<i) != 0 { // bitmask &(1<<i) == 1  你咋能这么自信，这是对的？ 你只是检查 bit i, 哎， 蠢到无语了啊
				w := words[i]
				for j := 0; j < len(w); j++ {
					total[w[j]-'a']++
				}
			}
		}
		for i := 0; i < 26; i++ {
			sum += total[i] * score[i]
		}
		return sum
	}

	check := func(bitmask int) bool {
		// check whether letters can satisfy the selected words
		total := [26]int{}
		for i := 0; i < n; i++ {
			if bitmask&(1<<i) != 0 {
				w := words[i]
				for j := 0; j < len(w); j++ {
					ic := w[j] - 'a'
					total[ic]++
					if total[ic] > ll[ic] {
						return false
					}
				}
			}
		}
		return true
	}

	ans := 0
	var dfs func(int, int)
	dfs = func(i int, bitmask int) {
		if i == n {
			ans = max(ans, cost(bitmask))
			return
		}

		dfs(i+1, bitmask)

		if check((1 << i) | bitmask) {
			dfs(i+1, bitmask|(1<<i))
		}
	}

	dfs(0, 0)
	return ans
}
