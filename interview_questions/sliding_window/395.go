package sliding_window

//wow !
func longestSubstring(s string, k int) int {
	n := len(s)

	ans := 0
	for t := 1; t <= 26; t++ {
		window := make(map[byte]int)
		satisfy := make(map[byte]struct{})
		left, right := 0, 0
		for right < n {
			r := s[right]
			window[r]++
			if window[r] >= k {
				satisfy[r] = struct{}{}
			}
			for len(window) > t {
				l := s[left]
				window[l]--
				if window[l] < k {
					delete(satisfy, l)
				}
				if window[l] == 0 {
					delete(window, l)
				}
				left++
			}

			if len(satisfy) == len(window) {
				ans = max(ans, right-left+1)
			}
			right++
		}
	}
	return ans
}
