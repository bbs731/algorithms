package simulation

/***
这道模拟题目好难啊！

下面自己的解法， 为啥是错误的?
 */
func calculate(s string) int {
	n := len(s)
	st := []int{}

	for i := 0; i < n; {
		if s[i] == ' ' {
			i++
			continue
		}
		if s[i] == '+' || s[i] == '-' || s[i] == '(' {
			st = append(st, int(s[i]))
			i++
			continue
		}
		start := 0
		if s[i] == ')' {
			start = st[len(st)-1]
			// consume the '('
			st = st[:len(st)-2]
			i++
		} else {
			// must be digit
			start = int(s[i] - '0')
			j := i + 1
			for ; j < n && s[j] >= '0' && s[j] <= '9'; j++ {
				start = start*10 + int(s[j]-'0')
			}
			i = j
		}
		// reduce to '(' or empty st
		for len(st) > 0 && st[len(st)-1] != '(' {
			op := st[len(st)-1]
			if op == '+' {
				start = st[len(st)-2] + start
				st = st[:len(st)-2]
			} else if op == '-' {
				if len(st) == 1 || st[len(st)-2] == '(' {
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
		st = append(st, start)
	}
	return st[0]
}

/***
官网的答案太简洁了, WTF
 */
func calculate(s string) int {
	n := len(s)
	st := []int{1} // operator stack, +1 and -1

	ans := 0
	sign := 1

	for i := 0; i < n; {
		switch s[i] {
		case ' ':
			i++
		case '+':
			sign = st[len(st)-1]
			i++
		case '-':
			sign = -st[len(st)-1]
			i++
		case '(':
			st = append(st, sign)
			i++
		case ')':
			st = st[:len(st)-1]
			i++
		default:
			num := 0
			j := i
			for ; j < n && s[j] >= '0' && s[j] <= '9'; j++ {
				num = num*10 + int(s[j]-'0')
			}
			num = sign * num
			ans += num
			i = j
		}
	}
	return ans
}
