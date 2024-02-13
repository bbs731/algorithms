package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)




/***
灵神的答案
 */
func CF1738B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n, m int

o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		g := make([]int, m)
		for j := 0; j < m; j++ {
			Fscan(in, &g[j])
		}

		if m> 1 && g[0] > (n-m+1)*(g[1]-g[0]) {
			Fprintln(out, "No")
			continue
		}

		for i := 2; i < m; i++ {
			if g[i-1]*2 > g[i]+g[i-2] {
				Fprintln(out, "No")
				continue o
			}
		}
		Fprintln(out, "Yes")
	}
}

func main() { CF1738B(os.Stdin, os.Stdout) }
