package stack


/***
太难了！
 */

func longestValidParentheses(s string) int {
	st := []int{-1}  // 这里太难想了！
	ans := 0
	for i :=range s {
		if s[i]	== ')'{
			//if len(st) == 0 {
			//	st = append(st, pair{')', i})
			//	continue
			//}
			st = st[:len(st)-1]
			if len(st) == 0 {
				st = append(st, i)	// put a dummy node  // 这里也是太难想了。
				continue
			}
			top := st[len(st)-1]
			ans = max(ans, i-top)
			// pop
		} else {
			//if len(st)>0 {
			//	ans = max(ans, i-st[len(st)-1].i-1)
			//}
			st = append(st,i)
		}
	}
	//if len(st) == 0 {
	//	return len(s)
	//}
	//ans = max(ans, len(s)-st[len(st)-1].i-1)
	return ans
}



func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}