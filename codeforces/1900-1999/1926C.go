package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

/***
 https://oi-wiki.org/dp/number/
用了 OI上   digit DP 的模板。
 */

func calSum(n int) int {
	const N int = 15
	var dp [N]int
	var mi [N]int
	var digits [N]int
	var ans [N]int

	tmp := n // 这里是天坑啊！
	mi[0] = 1
	for i := 1; i <= 13; i++ {
		dp[i] = dp[i-1]*10 + mi[i-1]
		mi[i] = 10 * mi[i-1]
	}

	var len int
	for n > 0 {
		len++
		digits[len] = n % 10
		n = n / 10
	}

	for i := len; i >= 1; i-- {
		//d := digits[i]
		for j := 1; j < 10; j++ {
			ans[j] += dp[i-1] * digits[i]
		}
		for j := 1; j < digits[i]; j++ {
			ans[j] += mi[i-1]
		}
		tmp -= mi[i-1] * digits[i]
		ans[digits[i]] += tmp + 1
		//ans[0] -= mi[i-1]
	}

	sum := 0
	for i := 1; i < 10; i++ {
		sum += ans[i] * i
	}
	return sum

}

func CF1926C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n int

	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		ans := calSum(n)
		Fprintln(out, ans)
	}
}

func main() { CF1926C(os.Stdin, os.Stdout) }
