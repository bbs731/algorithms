package interview_questions



/***
chunlei 你和牛，不要被眼前的挫折所累， 没关系， 坚持下去，证明你的想好好多就是对的，但是表达的不够高明。
向高手学习，进步终有一天你会和他们一样，站在山顶
 */

func findSubstring(s string, words []string) []int {
	n := len(words[0])
	wordm := make(map[string]int)
	for _, w := range words {
		wordm[w]++
	}
	m := len(wordm)
	ans := []int{}

	// loop 里面， 套一层 sliding window
	for j:=0;  j<n; j++ {
		left := j
		counter := make(map[string]int)
		// 枚举 right edge of window
		for i :=left; i <= len(s)-n; i+=n{
			next := s[i : i+n]
			if _, ok := wordm[next]; !ok {
				left = i+n
				// clear counter
				counter = make(map[string]int)
				continue
			}
			counter[next]++
			for counter[next] > wordm[next] {
				// move left edge of window
				leftword := s[left : left+n]
				left += n
				counter[leftword]--
				if counter[leftword] == 0 {
					delete(counter, leftword)
				}
			}
			// check answer
			if len(counter) == m {
				valid := true
				for k, v := range counter {
					if wordm[k] != v {
						valid = false
						break
					}
				}
				if valid {
					ans = append(ans, left)
				}
			}
		}
	}
	return ans
}
