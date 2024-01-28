package montonic_stack


/****
这题，出的太好了。 作为面试题的话， 面一次， 错一次。

需要考虑的边界条件太多了。
 */
func removeKdigits(num string, k int) string {
	st := []byte{}
	for _, d := range num {
		if len(st) == 0 {
			st = append(st, byte(d))
			continue
		}
		for len(st) > 0 && byte(d) < st[len(st)-1] && k>0 {  // 这里是坑1： 需要条件 k > 0
			st = st[:len(st)-1]
			k--
		}
		st = append(st, byte(d))
	}
	// deal with leading zero
	for len(st) > 1 && st[0] == '0'  {  // 这里是坑2： 需要处理前导零
		st = st[1:]
	}

	if k >= len(st) {   // 这里是坑3：需要处理 k 还有剩余的情况。 
		return "0"
	} else {
		st = st[:len(st)-k]
	}
	return string(st)
}
