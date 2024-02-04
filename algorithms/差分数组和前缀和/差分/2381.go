package difference

func shiftingLetters(s string, shifts [][]int) string {
	n := len(s)
	diffs := make([]int, n)

	for _, shift := range shifts {
		start, end, direction := shift[0], shift[1], shift[2]
		if direction == 0 {
			direction = -1
		}
		diffs[start] += direction
		if end+1 < n {
			diffs[end+1] -= direction
		}
	}

	for i := 1; i < n; i++ {
		diffs[i] += diffs[i-1]
	}

	ns := make([]byte, n)
	for i := range s {
		ns[i] = 'a' + byte(((int(byte(s[i]-'a'))+diffs[i])%26+26)%26) // 这里是个技巧， 处理负数的情况， 通过 mod 变成正的。
	}
	return string(ns)

}
