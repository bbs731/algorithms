package interview_questions

func findSubstring(s string, words []string) []int {
	ls, m, n := len(s), len(words), len(words[0])
	ans := []int{}
	for i := 0; i < n && i + n * m <= ls; i++ {
		count := map[string]int{}
		differ := 0
		for j := 0; j < m; j++ {
			count[words[j]]--
			count[s[i+j*n:i+(j+1)*n]]++
		}
		for _, v := range count {
			if v != 0 {
				differ++
			}
		}
		if differ == 0 {
			ans = append(ans, i)
		}
		for start := i; start < ls - (m+1)*n + 1 ; start += n {
			addWord := s[start+m*n : start + (m+1)*n]
			if count[addWord] == -1 {
				differ--
			} else if count[addWord] == 0  {
				differ++
			}
			count[addWord]++
			delWord := s[start:start+n]
			if count[delWord] == 1 {
				differ--
			} else if count[delWord] == 0 {
				differ++
			}
			count[delWord]--
			if differ == 0 {
				ans = append(ans, start+n)
			}
		}
	}
	return ans
}
