package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"strconv"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 还是会 TLE, 如何优化？
func countSpecialNumbers(n int) (ans int) {
	s := strconv.Itoa(n)
	m := len(s)

	if n < 0 {
		return -1
	}

	type pair struct {
		i int
		v [10]int
	}
	memo := make(map[pair]int)

	var f func(int, [10]int, bool, bool) int
	f = func(i int, cnts [10]int, isLimit, isNum bool) (res int) {
		if i == m {
			if isNum {
				ans := 0
				for k := 0; k < 10; k++ {
					ans = max(ans, cnts[k])
				}
				return ans
			}
			return
		}

		if !isLimit && isNum {
			p, ok := memo[pair{i, cnts}]
			if ok {
				return p
			}
		}
		if !isNum { // 可以跳过当前数位
			res += f(i+1, cnts, false, false)
		}
		d := 0
		if !isNum {
			d = 1 // 如果前面没有填数字，必须从 1 开始（因为不能有前导零）
		}
		up := 9
		if isLimit {
			up = int(s[i] - '0') // 如果前面填的数字都和 n 的一样，那么这一位至多填数字 s[i]（否则就超过 n 啦）
		}
		for ; d <= up; d++ { // 枚举要填入的数字 d
			cm := cnts
			cm[d]++
			res += f(i+1, cm, isLimit && d == up, true)
		}

		if !isLimit && isNum {
			memo[pair{i, cnts}] = res
		}
		return
	}
	return f(0, [10]int{}, true, false)
}

func cfbmode(_r io.Reader, _w io.Writer) {

	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n1, n2 int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n1, &n2)
		ans := countSpecialNumbers(n2) - countSpecialNumbers(n1-1)
		Fprintln(out, ans)
	}
}

func main() { cfbmode(os.Stdin, os.Stdout) }
