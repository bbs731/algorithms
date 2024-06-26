package weekly

func prefix_function(pattern string) []int {
	n := len(pattern)
	pi := make([]int, n)

	j := 0 // j 记录的是 pi[i-1], 初始化为 pi[0]  即为 0
	for i := 1; i < n; i++ {
		for j > 0 && pattern[i] != pattern[j] {
			j = pi[j-1]
		}

		if pattern[i] == pattern[j] {
			j++
		}
		pi[i] = j
	}
	return pi
}

func kmp(text, pattern string) []int {

	pi := prefix_function(pattern)
	pos := make([]int, 0)

	j := 0
	for i := 0; i < len(text); i++ {
		for j > 0 && pattern[j] != text[i] {
			j = pi[j-1]
		}
		if pattern[j] == text[i] {
			j++
		}
		if j == len(pattern) {
			pos = append(pos, i-len(pattern)+1)
			j = pi[j-1]
		}
	}
	return pos
}

func abs (a, b int)int {
	if a > b {
		return a - b
	}
	return b - a
}

func beautifulIndices(s string, a string, b string, k int) []int {
	apos := kmp(s, a)
	bpos := kmp(s, b)

	ans := []int{}
	for _, ap := range apos {
		for _, bp := range bpos {
			if abs (ap, bp)	 <=k {
				ans = append(ans, ap)
				break
			}
		}
	}
	return ans
}