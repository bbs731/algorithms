package interview_questions

/***
和 84 是一种题目， 编程能力吗？ 完全不用动脑子
 */
func compress(chars []byte) int {
	n := len(chars)
	start, end := 0, 0
	pos := 0
	for end<n {
		for end <n && chars[end] == chars[start]{
			end++
		}
		end--
		chars[pos] = chars[start]
		pos++

		cnts := end- start + 1
		if cnts > 1 {
			digits := make([]byte, 0)
			for cnts > 0 {
				digits = append(digits, byte(cnts%10) + '0')
				cnts /=10
			}
			for i := len(digits)-1; i>=0; i-- {
				chars[pos] = digits[i]
				pos++
			}
		}

		end++
		start = end
	}
	return pos
}
