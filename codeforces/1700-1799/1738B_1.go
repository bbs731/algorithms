package main

import (
	"bufio"
	"io"
	"os"
	. "fmt"
)

func CF1738B(_r io.Reader, _w io.Writer){
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n, k int
o:
	for Fscan(in, &T); T > 0; T--{
		Fscan(in, &n, &k)
		g := make([]int, k)
		for j:=0; j<k; j++ {
			Fscan(in, &g[j])
		}
		if k < 2 {
			Fprintln(out, "Yes")
			continue
		}
		// otherwise out of index
		lastd := g[k-1]-g[k-2]
		remaining := n

		for j:=k-1; j>=1; j-- {
			nextd := g[j]- g[j-1]
			if nextd > lastd {
				Fprintln(out, "No")
				continue o
			}
			lastd = nextd
			remaining--
		}
		if remaining == 0 && g[0] > lastd {
			Fprintln(out, "No")
			continue
		}
		//if remaining !=0 && (g[0]+remaining -1)/remaining > lastd {  // 这个逻辑是错的，因为 g[0] 有可能是个负数。
		if remaining != 0 &&g[0] > lastd*remaining {
			Fprintln(out, "No")
			continue
		}
		Fprintln(out, "Yes")
	}
}

func main() { CF1738B(os.Stdin, os.Stdout) }
