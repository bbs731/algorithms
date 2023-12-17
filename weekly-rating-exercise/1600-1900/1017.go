package _600_1900

/*
烂题，搞不定！

下面是错误的答案
 */
func lowest_neg_number (num int) (int, int){
	base := -2
	cnt := 1
	for num < base {
		base = base *4
		cnt += 2
	}
	return cnt, base
}
func lowest_base_number(num int) (int, int) {
	base := 1
	cnt := 0
	for num > base {
		base = base *2
		cnt++
	}
	return cnt, base
}
func baseNeg2(n int) string {
	var ans string
	if n == 0 {
		return "0"
	}
	content := make([]byte, 34)
	for ;;{
		ep, eb :=lowest_base_number(n)
		if n == eb {
			content[ep]= 1
			break
		}

		if ep %2 == 1 {
			content[ep-1]	= 1
			n = n- eb/2
		} else {
			content[ep]	= 1
			n = n-eb
			op, ob := lowest_neg_number(n)
			content[op]= 1
			if ob == n {
				break
			}
			n = n - ob
		}
	}

	started := false
	for i := len(content)-1; i>=0; i-- {
		if content[i] == 1 {
			started = true
		}
		if started {
			if content[i] == 1 {
				ans += "1"
			} else {
				ans += "0"
			}
		}
	}
	return ans
}
