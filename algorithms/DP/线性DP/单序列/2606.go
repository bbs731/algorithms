package dp

func maximumCostSubstring(s string, chars string, vals []int) int {
	v := [26]int{}
	for i := 0; i < 26; i++ {
		v[i] = i + 1
	}
	for i := range chars {
		c := chars[i]
		v[c-'a'] = vals[i]
	}

	l := make([]int, len(s))
	for i := range s {
		c := s[i]
		l[i] = v[c-'a']
	}
	ans := 0
	res := 0

	//fmt.Println(l)
	for _, n := range l {
		ans = max(max(ans+n, 0), n)
		res = max(res, ans)
	}
	return res
}
