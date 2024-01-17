package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

func cf1909C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	//var l, r, c []int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		l := make([]int, n)
		for i := range l {
			Fscan(in, &l[i])
		}
		r := make([]int, n)
		for i := range r {
			Fscan(in, &r[i])
		}

		c := make([]int, n)
		for i := range c {
			Fscan(in, &c[i])
		}

		ans := 0
		sort.Ints(c)
		sort.Ints(l)
		sort.Ints(r)
		type pair struct {
			s, e int // 区间的两个端点
		}
		pairs := make([]pair, 0)

		// a sliding window
		// 这里  match intervals 的想法是错误的。
		// 如何 match interval 的左右端点？ 让 e - s 尽量的大， 还能满足，最优能 n paired.

		// 如何去 match interval 答案里说， 这和 bracket matching  '('  和 ')' 配对是一个问题。
		// 灵神用了一个 stack 来实现。
		st := make([]int, 0)
		j := 0
		for _, v := range r {
			for j < n && l[j] < v {
				st = append(st, l[j])
				j++
			}
			// v and top of stack form a pair
			pairs = append(pairs, pair{st[len(st)-1], v})
			// pop stack
			st = st[:len(st)-1]
		}
		//right := n - 1
		//for right >= 0 {
		//	left := right - 1
		//	for left >= 0 {
		//		if r[left] > l[right] {
		//			left--
		//		} else {
		//			break
		//		}
		//	}
		//	j := right
		//	for i := left + 1; i <= right; i++ {
		//		pairs = append(pairs, pair{l[i], r[j]})
		//		j--
		//	}
		//	right = j
		//}

		// 按照 Interval length 倒序排列
		sort.Slice(pairs, func(i, j int) bool { return (pairs[i].e - pairs[i].s) > (pairs[j].e - pairs[j].s) })

		for i := 0; i < n; i++ {
			ans += c[i] * (pairs[i].e - pairs[i].s)
		}
		Fprintln(out, ans)
	}
}

func main() { cf1909C(os.Stdin, os.Stdout) }
