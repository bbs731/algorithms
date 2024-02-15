package interview_questions


func findSubstring(s string, words []string) []int {
	n := len(words[0])
	m :=len(words)

	wordm := make(map[string]int)
	for _, w := range words {
		wordm[w]++
	}

	ans := make([]int, 0)

	// 枚举 right edge of window
	for i:=0; i<len(s)-n*m+1; i++{
		tmp := s[i:i+m*n]
		counter := make(map[string]int)
		for j:=0; j<len(tmp); j+=n{
			w := tmp[j:j+n]
			counter[w]++
		}
		// check answer
		if len(counter)== len(wordm) {
			valid := true
			for k, v := range counter {
				if wordm[k]	 != v {
					valid = false
					break
				}
			}
			if valid {
				ans = append(ans, i)
			}
		}
	}
	return ans
}
