package leetcode

/*
code template from  OI,  digit DP
https://oi-wiki.org/dp/number/

 */
func countDigitOne(n int) (ans int) {

	var p [20]int
	var now [20]int
	var ksm [20]int
	var f [20]int

	ksm[0] = 1
	for i := 1; i < 20; i++ {
		ksm[i] = 10 * ksm[i-1]
	}
	for i := 0; i < 20; i++ {
		f[i] = -1
	}

	d := n
	l := 0

	for d != 0 {
		l++
		p[l] = d % 10
		d /= 10
		now[l] = now[l-1] + p[l]*ksm[l-1]
	}

	var dfs func(int, int, bool, bool) int

	dfs = func(u, x int, f0, lim bool) int {

		if u == 0 {
			if f0 {
				f0 = false
			}
			return 0
		}

		if !lim && !f0 && f[u] != -1 {
			return f[u]
		}

		cnt := 0
		lst := 9
		if lim {
			lst = p[u]
		}

		for i := 0; i <= lst; i++ {
			if f0 && i == 0 {
				cnt += dfs(u-1, x, true, lim && i == lst)
			} else if i == x && lim && i == lst {
				cnt += now[u-1] + 1 + dfs(u-1, x, false, lim && i == lst)
			} else if i == x {
				cnt += ksm[u-1] + dfs(u-1, x, false, lim && i == lst)
			} else {
				cnt += dfs(u-1, x, false, lim && i == lst)
			}
		}

		if !lim && !f0 {
			f[u] = cnt
		}
		return cnt

	}
	return dfs(l, 1, true, true)
}
