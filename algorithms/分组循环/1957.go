package loop

//按照灵神的分组循环的模版写的，爽！
func makeFancyString(s string) string {
	n := len(s)
	ns := make([]byte, n)
	pos :=0

	for i:=0; i<n; {
		start := i
		for ; i<n && s[i]== s[start]; i++ {
			if i - start + 1 <= 2 {
				ns[pos]	= s[i]
				pos++
			}
		}
	}
	return string(ns[:pos])
}
