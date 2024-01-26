package loop


func checkZeroOnes(s string) bool {
	n := len(s)
	i := 0
	l1 := 0
	l0 := 0

	for i < n {
		start := i
		for i <n && s[i] == s[start] {
			i++
		}
		if s[i] == '1' {
			l1 = max(l1, i-start + 1)
		}else {
			l0 = max(l0,i-start +1)
		}
		//i++
	}
	return l1 > l0
}