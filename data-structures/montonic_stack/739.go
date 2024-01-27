package montonic_stack


func dailyTemperatures(temp []int) []int {
	n := len(temp)
	st := make([]int, 0, n) // 0 初始化长度， 犯错 1
	ans := make([]int, n)

	// 从左向右，还是从右向左遍历都可以。
	for i:=n-1; i>=0; i-- {
		for len(st)!=0 && temp[i] > temp[st[len(st)-1]] {   //  这是 犯得最2 个错误。
			st = st[:len(st)-1]
		}
		if len(st) == 0 {
			ans[i] = 0  // 初始化的时候，已经设置为 0 了，就不需要 if clause 了, 可以把逻辑变的更短
		}else {
			ans[i] = st[len(st)-1] - i
		}
		st = append(st, i)  // 忘了放index 犯错3
	}
	return ans
}

