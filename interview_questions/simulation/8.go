package simulation

// 永远的伤心地！
func myAtoi(s string) int {
	biggest := 1 << 31
	num := 0
	plus := true
	start := 0
	for ; start < len(s) && s[start] == ' '; start++ { // panic 没检查 start 越界
	}

	if start >= len(s) {
		return num
	}
	if s[start] == '-' {
		plus = false
		start++
	} else if s[start] == '+' {
		start++
	}
	for ; start < len(s) && s[start] >= '0' && s[start] <= '9' && num <= biggest; start++ { // panic 没检查 start 越界
		num = 10*num + int(s[start]-'0')
	}
	// end of parsing
	if plus {
		if num >= biggest {
			num = biggest - 1
		}
	} else {
		if num > biggest {
			num = biggest
		}
		num = -num
	}
	return num
}

/***
这样写，也是可以的， 思路好像也挺清晰的， 但是一次写对，真是太难了！
 */
func myAtoi(s string) int {
	biggest := 1 << 31
	ans := 0
	sign := 1
	started := false

	for i := 0; i < len(s); {
		if !started {
			if s[i] == ' ' || s[i] == '\t' {
				i++
				continue
			}
			if s[i] == '-' && !started {
				sign = -1
				i++
				started = true
				continue
			}
			if s[i] == '+' && !started {
				i++
				started = true
				continue
			}
		}

		if s[i] >= '0' && s[i] <= '9' && ans <= biggest {
			ans = 10*ans + int(s[i]-'0')
			started = true
			i++
		} else {
			break
		}

	}

	if ans >= biggest {
		if sign == 1 {
			ans = biggest - 1
		} else {
			ans = biggest
		}
	}
	return ans * sign
}
