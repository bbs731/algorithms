package z_algorithm

/***
code from endlesscheng
https://github.com/EndlessCheng/codeforces-go/blob/master/misc/atcoder/abc284/f/f.go#L11
*/

func z_function(s string) []int {
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
	return z
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
