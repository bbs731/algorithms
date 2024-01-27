package loop

/***
模版题目， 模版题目
套路， 套路。
 */

func countHomogenous(s string) int {
	n := len(s)
	ans := 0

	for i:=0; i<n; {
		start := i
		for ;i<n && s[i] == s[start]; i++ {
		}
		l := i - start
		ans += l*(l+1)/2
	}
	return ans%(int(1e9)+7)
}
