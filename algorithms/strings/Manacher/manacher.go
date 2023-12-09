package Manacher

/*
https://oi-wiki.org/string/manacher/

给定一个长度为 n 的字符串 s，请找到所有对 (i, j)  使得子串 s[i...j] 为一个回文串。
根据回文的长度， 计算， d1 和 d2
 */
func d1(d1 []int, s string) []int {
	n := len(s)

	l, r := 0, -1

	for i := 0; i < n; i++ {
		k := 1
		if i < r {
			k = min(d1[l+r-i], r-i+1)
		}
		for i-k >= 0 && i+k < n && s[i-k] == s[i+k] {
			k++
		}
		d1[i] = k
		k--
		if i+k > r {
			l = i - k
			r = i + k
		}
	}
	return d1
}

func d2(d2 []int, s string) []int {
	n := len(s)
	l, r := 0, -1

	for i := 0; i < n; i++ {
		k := 1
		if i < r {
			k = min(d2[l+r-i+1], r-i+1)
		}
		for i-k-1 >= 0 && i+k < n && s[i-k-1] == s[i+k] {
			k++
		}
		d2[i] = k
		k--

		if i+k > r {
			r = i + k
			l = i - k - 1
		}
	}
	return d2
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
