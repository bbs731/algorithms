package interview_questions

func findSubstring(s string, words []string) []int {
	k := len(words)
	n := len(words[0])
	if len(s) < n {
		return []int{}
	}

	wordCounts := make(map[string]int, n)
	for _, w := range words {
		wordCounts[w]++
	}

	seenTimes := make(map[string]int, n)
	var head, qlen int
	q := make([]string, 0, k)
	reset := func() {
		head, qlen = 0, 0
		q = q[:0]
		for k := range seenTimes {
			delete(seenTimes, k)
		}
	}

	res := []int{}
	for i := 0; i < n; i++ {
		reset()
		for j := i; j+n <= len(s); j += n {
			w := s[j : j+n]
			wordCount, exist := wordCounts[w]
			if !exist {
				reset()
				continue
			}

			for seenTimes[w] == wordCount {
				seenTimes[q[head]]--
				head++
				qlen--
			}

			seenTimes[w]++
			q = append(q, w)
			qlen++
			if qlen == k {
				res = append(res, j-(k-1)*n)
			}
		}
	}
	return res
}
