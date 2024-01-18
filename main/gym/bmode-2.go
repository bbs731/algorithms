package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"runtime/debug"
	"strconv"
	"strings"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 还是会 TLE, 如何优化？
func countSpecialNumbers(low, high int) (ans int) {

	lowS := strconv.Itoa(int(low))
	highS := strconv.Itoa(int(high))
	n := len(highS)
	lowS = strings.Repeat("0", n-len(lowS)) + lowS // 对齐

	type pair struct {
		i int
		v [19]int
	}

	var memo = make(map[pair]int)

	var f func(int, []int, bool, bool, bool) int
	f = func(i int, cnts []int, limitLow, limitHigh, isNum bool) (res int) {

		if i == n {
			if isNum {
				ans := 0
				for k := 0; k < 10; k++ {
					ans = max(ans, cnts[k])
				}
				return ans
			}
			if low == 0 {
				return 1
			}
			return
		}

		//cachekey := [10]int{}
		cachekey := [19]int{} // cachekey[i] 代表出现 i 次 的 digit 的种类数。
		if !limitHigh && !limitLow && isNum {
			// sort cnts  作为 cachekey 的优化还是不够快。
			//d := make([]int, 10)
			//copy(d, cnts)
			//sort.Ints(d)
			//for k := 0; k < len(d); k++ {
			//	cachekey[k] = d[k]
			//}
			for k := 0; k < len(cnts); k++ {
				cachekey[cnts[k]]++
			}
			p, ok := memo[pair{i, cachekey}]
			if ok {
				return p
			}
		}
		if !isNum && lowS[i] == '0' { // 可以跳过当前数位
			res += f(i+1, cnts, true, false, false)
		}

		lo := 0
		if limitLow {
			lo = int(lowS[i] - '0') // 如果前面没有填数字，必须从 1 开始（因为不能有前导零）
		}
		up := 9
		if limitHigh {
			up = int(highS[i] - '0') // 如果前面填的数字都和 n 的一样，那么这一位至多填数字 s[i]（否则就超过 n 啦）
		}
		d := lo

		if !isNum {
			d = max(lo, 1)
		}

		for ; d <= up; d++ { // 枚举要填入的数字 d
			cnts[d]++
			res += f(i+1, cnts, limitLow && d == lo, limitHigh && d == up, true)
			cnts[d]--
		}

		if !limitLow && !limitHigh && isNum {
			memo[pair{i, cachekey}] = res
		}
		return
	}

	return f(0, make([]int, 10), true, true, false)
}

func cfbmode2(_r io.Reader, _w io.Writer) {
	debug.SetGCPercent(-1)
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n1, n2 int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n1, &n2)
		ans := countSpecialNumbers(n1, n2)
		Fprintln(out, ans)
	}
}

func main() { cfbmode2(os.Stdin, os.Stdout) }
