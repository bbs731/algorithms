package weekly

func lengthOfLongestSubstring(s string) int {
	l := 0
	r := 0
	n := len(s)
	m := make(map[byte]int)
	ans := 0

	for r < n {
		if pos, ok := m[s[r]]; ok {
			for i := l; i < pos; i++ {
				delete(m, s[i])
			}
			l = pos + 1
			m[s[r]] = r
		} else {
			m[s[r]] = r
			ans = max(ans, r-l+1)
		}
		r++
	}

	return ans
}
