package interview_questions


/***
受到了 224 计算器的启发

做过类似的题目确实是不一样啊, 一次过
 */
func calculate(s string) int {
	n := len(s)
	sign := 1
	ans := 0
	st := []int{}
	op := []byte{}

	for i:=0; i<n; {
		if s[i]	 == ' ' {
			i++
			continue
		}

		if s[i] == '+' {
			prev := st[len(st)-1]
			ans += sign*prev
			st = st[:len(st)-1]
			sign= 1
			i++
			continue
		}

		if s[i] == '-' {
			prev := st[len(st)-1]
			ans += sign*prev
			st = st[:len(st)-1]
			sign = -1
			i++
			continue
		}

		if s[i] == '*' {
			op = append(op, s[i])
			i++
			continue
		}

		if s[i]== '/'{
			op = append(op, s[i])
			i++
			continue
		}

		// now reach digit
		num := 0
		for i<n && s[i]>='0' && s[i]<='9' {
			num = num*10 + int(s[i]-'0')
			i++
		}
		if len(op) != 0 {
			if op[len(op)-1] == '*' {
				st[len(st)-1] = st[len(st)-1]*num
			}else {
				st[len(st)-1] = st[len(st)-1]/num
			}
			// pop op
			op = op[:len(op)-1]
		} else {
			st = append(st, num)
		}
	}
	// now op is empty and, check st
	if len(st) != 0 {
		ans += sign*st[0]
	}
	return ans
}