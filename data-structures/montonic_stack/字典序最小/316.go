package montonic_stack

/****
不会做这个题， 结果做完 402之后，就有感觉了。
 */

func removeDuplicateLetters(s string) string {
	counter := make(map[byte]int)
	for _, c := range s {
		counter[byte(c)]++
	}

	st := []byte{}
	instack := make(map[byte]bool)
	for _, c := range s {
		if len(st) == 0 {
			st = append(st, byte(c))
			instack[byte(c)] = true
		} else {
			for len(st) > 0 && instack[byte(c)] == false && byte(c) <= st[len(st)-1] && counter[st[len(st)-1]] > 1{
				// pop stack
				counter[st[len(st)-1]]--
				instack[st[len(st)-1]] = false
				st = st[:len(st)-1]
			}
			if instack[byte(c)] == false{
				st = append(st, byte(c))
				instack[byte(c)] = true
			} else {
				// drop char c
				counter[byte(c)]--
			}
		}
	}
	return string(st)
}