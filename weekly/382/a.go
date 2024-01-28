package weekly

import "strings"

func countKeyChanges(s string) int {
	ss := strings.ToLower(s)
	cnt :=0
	n := len(ss)
	for i:=1; i<n; i++{
		if ss[i] != ss[i-1] {
			cnt++
		}
	}
	return cnt
}
