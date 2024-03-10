package weekly

import "strings"

/***

["gfnt","xn","mdz","yfmr","fi","wwncn","hkdy"]

["abc","bcd","abcd"]
 */


func shortestSubstrings(arr []string) []string {
	n := len(arr)
	ans := make([]string, n)

	for i := range arr {
		text := strings.Builder{}
		for j := range arr {
			if j != i {
				text.WriteString("$")
				text.WriteString(arr[j])
				text.WriteString("$")
			}
		}
		pt := text.String()
		for k:=1; k<=len(arr[i]); k++ {
			for j := 0; j+k-1 < len(arr[i]); j++ {
				m := j + k - 1
				if !strings.Contains(pt, arr[i][j:m+1]) {
					if ans[i]== "" {
						ans[i] = arr[i][j : m+1]
					} else {
						if arr[i][j:m+1] < ans[i]{
							ans[i] = arr[i][j : m+1]
						}
					}
				}
			}
			if ans[i] != "" {
				break
			}
		}
	}
	return ans
}