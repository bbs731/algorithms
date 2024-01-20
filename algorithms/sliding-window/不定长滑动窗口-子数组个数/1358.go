package sliding_window

//795， 992， 1358， 1248, 930模板题

// 这道题和 2799 用到的技巧是一样的。 这是谁发现的规律？
func numberOfSubstrings(s string) int {
	cnt := make(map[int32]int)
	left := 0
	ans := 0

	// 枚举窗口的右端点。
	for _, c := range s {
		cnt[c]++
		for len(cnt) == 3 {
			// we have collected a, b and c
			l := int32(s[left])
			cnt[l]--
			if cnt[l] == 0 {
				delete(cnt, l)
			}
			left++
		}
		ans += left
	}
	return ans
}
