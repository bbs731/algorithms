package monotonic_stack

/*
[6,9,9]
[6,6,9]
[9,6,9]
[1,1,1,9,1,9, 1, 1 ,9,9]
[8,12,7,6,10,10,9,11,12,6]
[12,6,7,12,7,11,12,6,8,6,7,8,7]
看看测试用例， 是不是想简单了。
*/
/*
太难了。但是，最接近真实的面试题目的好题。 多做几遍！
 */
func longestWPI(hours []int) int {
	n := len(hours)
	ans := 0
	presum := make([]int, n+1)
	//st := make([]int, 0)
	st := []int{0}
	for i, x := range hours {
		presum[i+1] = presum[i]
		if x > 8 {
			presum[i+1]++
		} else {
			presum[i+1]--
		}
		if presum[i+1] < presum[st[len(st)-1]] {
			st = append(st, i+1)
		}
	}

	for i:=n-1; i>=0; i-- {
		for len(st) >0 && presum[i+1]> presum[st[len(st)-1]] {
			ans = max(ans, i-st[len(st)-1]+1)
			st = st[:len(st)-1]
		}
	}
	return ans
}

// you are handsome baby!
func longestWPI(hours []int) int {
	n := len(hours)
	ans := 0
	presum := make([]int, n+1)
	st := []int{-1}
	for i, x := range hours {
		presum[i+1] = presum[i]
		if x > 8 {
			presum[i+1]++
		} else {
			presum[i+1]--
		}
		if presum[i+1] < presum[st[len(st)-1] + 1] {
			st = append(st, i)
		}
	}

	for i:=n-1; i>=0; i-- {
		for len(st) >0 && presum[i+1]> presum[st[len(st)-1] + 1 ] {
			ans = max(ans, i-st[len(st)-1])
			st = st[:len(st)-1]
		}
	}
	return ans
}



