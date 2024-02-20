package interview_questions


func calculate(s string) int {
	n := len(s)
	st := []int{}

	for i:=0; i<n; {
		if s[i]== ' ' {
			i++
			continue
		}
		if s[i] == '+'|| s[i]== '-'|| s[i]== '(' {
			st = append(st, int(s[i]))
			i++
			continue
		}
		start := 0
		if s[i]==')' {
			start = st[len(st)-1]
			// consume the '('
			st = st[:len(st)-2]
			i++
		} else {
			// must be digit
			start = int(s[i] - '0')
			j := i+1
			for ; j<n && s[j]>='0' && s[j]<='9'; j++ {
				start = start*10 + int(s[j]-'0')
			}
			i = j
		}
		// reduce to '(' or empty st
		for len(st)>0 && st[len(st)-1] != '(' {
			op :=  st[len(st)-1]
			if op == '+' {
				start = st[len(st)-2] +start
				st = st[:len(st)-2]
			} else if op == '-' {
				if len(st) == 1  ||  st[len(st)-2] == '(' {
					start = - start
					st = st[:len(st)-1]
				} else if st[len(st)-2] == '(' {
					st = st[:len(st)-2]
				} else {
					start = st[len(st)-2] - start
					st = st[:len(st)-2]
				}
			}
		}
		st = append(st,start)
	}
	return st[0]
}
