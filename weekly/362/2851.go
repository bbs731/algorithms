package weekly

import "fmt"

/*

状态转移方程：

C = kmp(text+text -1, pattern)

可以选择切分的位置是 n-1， 其中有 C 个可以得到 t,   n-1-C 个得不到 t

f[i][0] = f[i-1][0] * (C-1) +  f[i-1][1] * C
f[i][1] = f[i-1][0] * (n-1 - (C-1))  + f[i-1][1] * (n-1 -C)

[				[								[
f[n][0]    =  		C-1,    C			n * 		f[0][0]
f[n][1]				n-C,    n-1-C
]									]				f[0][1]	]
 */

const mod = 1_000_000_007
func numberOfWays(s string, t string, k int64) int {
	n := len(s)
	pos := kmp(s+s, t, n) // 这里， pass s+s 的话，需要 pass n。 或者像下面那样也可以 pass s+s -1, 就不需要 pass n 了。
	//pos := kmp((s + s)[:2*n-1], t, n)
	c := len(pos)
	fmt.Println(c)

	m := matrix{
		{c - 1, c},
		{n - c, n - 1 - c},
	}.Power(k)

	if s == t {
		return m[0][0]
	}
	return m[0][1]
}

type matrix [][]int

func newIndentity(n int) matrix {
	m := make([][]int, n)
	for i := 0; i < n; i++ {
		m[i] = make([]int, n)
		m[i][i] = 1
	}
	return m
}

func (a matrix) Mul(b matrix) matrix {
	// a=[m, n]    b = [n , p]   c= [m, p]
	m := len(a)
	n := len(a[0])
	p := len(b[0])

	c := make([][]int, m)
	for i := 0; i < m; i++ {
		c[i] = make([]int, p)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < p; j++ {
			for k := 0; k < n; k++ {
				c[i][j] = (c[i][j] + a[i][k]*b[k][j]) % mod
			}
		}
	}
	return c
}

func (m matrix) Power(n int64) matrix {
	ans := newIndentity(len(m))
	a := m
	for n > 0 {
		if n&1 > 0 {
			ans = ans.Mul(a) // 这里掉进坑里面了。
		}
		a = a.Mul(a)
		n /= 2
	}
	return ans
}

func calcMathes(pattern string) []int {
	n := len(pattern)
	matches := make([]int, n)

	c := 0
	for i := 1; i < n; i++ { // 这里的 index 从 1 开始 , 这就是有坑的地方
		for c > 0 && pattern[c] != pattern[i] {
			c = matches[c-1]
		}
		if pattern[c] == pattern[i] {
			c++
		}
		matches[i] = c
	}
	return matches
}

func kmp(s, t string, original int) []int { // 这个 s 是原来的 s + s , 这里的 original 是原来的 len(s)
	n := len(s) // 这个 s 已经是  s+s 的结果了。
	match := calcMathes(t)
	pos := make([]int, 0)

	c := 0
	for i := 0; i < n; i++ { // 这里要从 0 开始， 不像 pattern 的处理需要从 1 开始
		for c > 0 && t[c] != s[i] {
			c = match[c-1]
		}
		if t[c] == s[i] {
			c++
		}

		if c == len(t) {
			if i-len(t)+1 < original {
				//cnt++
				pos = append(pos, i-len(t)+1)
			}

			//pos = append(pos, i-len(t)+1)
			c = match[c-1]
		}
	}
	return pos
}
