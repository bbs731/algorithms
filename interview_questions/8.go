package interview_questions

/***
这玩仍, 啥时候能写一次没bug 的？

bug 了 4次
 */
func myAtoi(s string) int {
	biggest := 1 << 31
	num := 0
	plus := true
	start := 0

	//1. panic 没检查 start 越界
	for ; start < len(s) && s[start] == ' '; start++ {
	}

	//2. 检查边界， 可能会引起 bug   test case "" empty string
	if start >= len(s) {
		return num
	}

	if s[start] == '-' {
		plus = false
		start++
	} else if s[start] == '+' {
		start++
	}

	// 3. panic 没检查 start 越界
	// 4. num > biggest 的时候，需要终止
	for ; start < len(s) && s[start] >= '0' && s[start] <= '9' && num <= biggest; start++ {
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
