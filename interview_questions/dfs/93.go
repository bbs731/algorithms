package dfs

import (
	"strconv"
	"strings"
)

func restoreIpAddresses(s string) []string {
	ans := []string{}
	n := len(s)
	var dfs func(int, int, []string)
	dfs= func(i int, p int, l []string) {
		if i == 4 {
			if p == n {
				// found ans
				ans = append(ans, strings.Join(l, "."))
			}
			return
		}

		if p >= n {
			return
		}
		if s[p] == '0' {
			l = append(l, "0")
			dfs(i+1, p+1, l)
			return
		}
		for j:=0; j<=2 && p+j <n ; j++ {
			d, e := strconv.Atoi(s[p:p+j+1])
			if e != nil {
				return
			}
			if d > 255 {
				return
			}
			l = append(l, s[p:p+j+1])
			dfs(i+1, p+j+1, l)
			l = l[:len(l)-1]
		}

	}
	dfs(0, 0, []string{})
	return ans
}
